// https://github.com/usbarmory/tamago-example
//
// Copyright (c) WithSecure Corporation
// https://foundry.withsecure.com
//
// Use of this source code is governed by the license
// that can be found in the LICENSE file.

//go:build mx6ullevk || usbarmory
// +build mx6ullevk usbarmory

package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
	"syscall"

	"github.com/usbarmory/tamago/soc/imx6"
)

type muxclone struct {
	m    sync.Mutex
	nmux int
	ctl  []muxctl
}

var (
	_   syscall.DevFile = &muxclone{}
	mux                 = &muxclone{}
)

func (m *muxclone) Pread(out []byte, addr int64) (int, error) {
	var b = bytes.NewBufferString(fmt.Sprintf("%d devices", m.nmux))
	return b.Read(out)
}

func (m *muxclone) Pwrite(in []byte, addr int64) (int, error) {
	return -1, fmt.Errorf("fuck off")
}

// muxctl is returned from a clone.
type muxctl struct {
	m     sync.Mutex
	num   uint
	mux   uint32
	pad   uint32
	daisy uint32
	dev   *muxdata
	ready bool
}

func init() {
	// muxclone creates a new mux instance, unitialized.
	if err := syscall.MkDev("/dev/muxclone", 0666, func() (syscall.DevFile, error) {
		mux.m.Lock()
		defer mux.m.Unlock()
		i := mux.nmux
		mux.nmux++
		n := fmt.Sprintf("/dev/%dctl", i)
		if err := syscall.MkDev(n, 0666, func() (syscall.DevFile, error) {
			mux.m.Lock()
			defer mux.m.Unlock()
			m := &mux.ctl[i]
			m.num = uint(i)
			return m, nil
		}); err != nil {
			return nil, fmt.Errorf("Can't set up %s: %w", n, err)
		}
		return mux, nil
	}); err != nil {
		log.Printf("Can't set up devmux: %v", err)
	}
}

func (m *muxctl) String() string {
	if !m.ready {
		return "not ready"
	}
	return fmt.Sprintf("mux %#x pad %#x daisy %#x", m.mux, m.pad, m.daisy)

}

// Pread reads mux information.
func (m *muxctl) Pread(out []byte, addr int64) (int, error) {
	m.m.Lock()
	defer m.m.Unlock()
	s := m.String()
	b := bytes.NewReader([]byte(s))
	return b.ReadAt(out, addr)
}

// Pwrite writes longwords to memory.
func (m *muxctl) Pwrite(in []byte, addr int64) (int, error) {
	m.m.Lock()
	defer m.m.Unlock()
	if m.ready {
		return -1, fmt.Errorf("%v is already set", m.String())
	}
	if len(in) == 0 {
		return 0, nil
	}
	f := strings.Fields(string(in))
	if len(f) != 3 {
		return -1, fmt.Errorf("Usage: mux pad daisy")
	}

	for i, p := range []*uint32{&m.mux, &m.pad, &m.daisy} {
		// the range seems to be 14 bits?
		v, err := strconv.ParseUint(f[i+1], 0, 14)
		if err != nil {
			return -1, err
		}
		*p = uint32(v)
	}
	p, err := imx6.NewPad(m.mux, m.pad, m.daisy)

	if err != nil {
		return -1, err
	}
	m.ready = true
	n := fmt.Sprintf("/dev/%dctl", m.num)
	if err := syscall.MkDev(n, 0666, func() (syscall.DevFile, error) {
		mux.m.Lock()
		defer mux.m.Unlock()
		m.dev = &muxdata{
			p: p,
		}
		return m.dev, nil
	}); err != nil {
		return -1, fmt.Errorf("Can't set up %s: %w", m.num, err)
	}
	return len(in), nil
}

var _ syscall.DevFile = &muxctl{}

type muxdata struct {
	m sync.Mutex
	// Pad has no stringer and exports nothing. FIX.
	p *imx6.Pad
}

// Pread reads muxdata
func (m *muxdata) Pread(out []byte, addr int64) (int, error) {
	m.m.Lock()
	defer m.m.Unlock()
	return -1, fmt.Errorf("not yet")
}

// Pwrite writes longwords to muxdata
func (m *muxdata) Pwrite(in []byte, addr int64) (int, error) {
	m.m.Lock()
	defer m.m.Unlock()
	if len(in) == 0 {
		return 0, nil
	}
	return -1, fmt.Errorf("not yet")

}

var _ syscall.DevFile = &muxdata{}
