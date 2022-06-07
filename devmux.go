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
	"errors"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
	"sync"
	"syscall"

	"github.com/usbarmory/tamago/soc/imx6"
)

type muxclone struct {
	m   sync.Mutex
	ctl []*muxctl
}

var (
	_   syscall.DevFile = &muxclone{}
	mux                 = &muxclone{}
)

func (m *muxclone) Pread(out []byte, addr int64) (int, error) {
	return -1, fmt.Errorf("%v: muxclone read", m)
	var b = bytes.NewReader([]byte(fmt.Sprintf("muxclone: %d devices", len(m.ctl))))
	return b.ReadAt(out, addr)
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
		// DO NOT DO log.printf. Ends very badly.
		// likely because, well, we're in the middle of network traffic
		// or something? not sure. Need to clean this up.
		// this works and gets a nice error message.
		//return nil, fmt.Errorf("fuck me")
		// crash. log.Printf("muxclone open")
		mux.m.Lock()
		defer mux.m.Unlock()
		if false {
			return nil, fmt.Errorf("locked")
		}
		n := fmt.Sprintf("/dev/mux%dctl", len(mux.ctl))
		if false {
			return nil, fmt.Errorf(fmt.Sprintf("New dev %s", n))
		}
		m := &muxctl{num: uint(len(mux.ctl))}
		mux.ctl = append(mux.ctl, m)
		if false {
			return nil, fmt.Errorf(fmt.Sprintf("cons'ed up dev %s", n))
		}
		// OK to here.
		if err := syscall.MkDev(n, 0666, func() (syscall.DevFile, error) {
			if false {
				return nil, fmt.Errorf(fmt.Sprintf("open dev %s", n))
			}
			m.m.Lock()
			defer m.m.Unlock()
			if false {
				return nil, fmt.Errorf(fmt.Sprintf("dev %s locked", n))
			}
			return m, nil
		}); err != nil {
			return nil, fmt.Errorf("Can't set up %s: %w", n, err)
		}
		return m, nil
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
	// Something, somewhere, is seeing an io.EOF and dropping data. dammit.
	// Well, for now, we'll eat the extra message.
	n, err := b.ReadAt(out, addr)
	if errors.Is(err, io.EOF) && n > 0 {
		err = nil
	}
	return n, err
	//	return -1, fmt.Errorf("muxctl read %q %v %d into %d len buf @ %d n %v err %v", s, b, b.Len(), len(out), addr, n, err)
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
	return -1, fmt.Errorf("muxdata read")
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
