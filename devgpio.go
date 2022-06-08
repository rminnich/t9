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
	"os"
	"strconv"
	"strings"
	"sync"
	"syscall"

	"github.com/usbarmory/tamago/soc/imx6"
)

type gpioclone struct {
	m   sync.Mutex
	ctl []*gpioctl
}

var (
	_    syscall.DevFile = &gpioclone{}
	gpio                 = &gpioclone{}
)

func (m *gpioclone) Pread(out []byte, addr int64) (int, error) {
	return -1, fmt.Errorf("%v: gpioclone read", m)
	var b = bytes.NewReader([]byte(fmt.Sprintf("gpioclone: %d devices", len(m.ctl))))
	return b.ReadAt(out, addr)
}

func (m *gpioclone) Pwrite(in []byte, addr int64) (int, error) {
	return -1, fmt.Errorf("fuck off")
}

// gpioctl is returned from a clone.
type gpioctl struct {
	m     sync.Mutex
	name  string
	ready bool
	num   uint
	dev   *gpiodata
}

func init() {
	// gpioclone creates a new gpio instance, unitialized.
	if err := syscall.MkDev("/dev/gpioclone", 0666, func() (syscall.DevFile, error) {
		gpio.m.Lock()
		defer gpio.m.Unlock()
		if false {
			return nil, fmt.Errorf("locked")
		}
		n := fmt.Sprintf("/dev/gpio%dctl", len(gpio.ctl))
		if false {
			return nil, fmt.Errorf(fmt.Sprintf("New dev %s", n))
		}
		m := &gpioctl{num: uint(len(gpio.ctl))}
		gpio.ctl = append(gpio.ctl, m)
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
		log.Printf("Can't set up devgpio: %v", err)
	}
}

func (m *gpioctl) String() string {
	if !m.ready {
		return "not ready"
	}
	return fmt.Sprintf("gpio %s num %d data %d dir %d", m.dev.pad.Name(), m.num, m.dev.pad.Name(), m.dev.data, m.dev.dir)
}

// Pread reads gpio information.
func (m *gpioctl) Pread(out []byte, addr int64) (int, error) {
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
	//	return -1, fmt.Errorf("gpioctl read %q %v %d into %d len buf @ %d n %v err %v", s, b, b.Len(), len(out), addr, n, err)
}

// Pwrite implements syscall.Pwrite
func (m *gpioctl) Pwrite(in []byte, addr int64) (int, error) {
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
		return -1, fmt.Errorf("Usage: gpio pad daisy")
	}
	if false {
		return -1, fmt.Errorf("gpioctl write %v", f)
	}

	v, err := strconv.ParseUint(f[1], 0, 5)
	if err != nil {
		return -1, err
	}
	num := uint(v)

	if v, err = strconv.ParseUint(f[1], 0, 2); err != nil {
		return -1, err
	}
	base := []int{imx6.GPIO1_BASE, imx6.GPIO2_BASE, imx6.GPIO3_BASE, imx6.GPIO4_BASE}[v]
	m.ready = true

	n := fmt.Sprintf("/dev/gpio%ddata", m.num)
	if err := syscall.MkDev(n, 0666, func() (syscall.DevFile, error) {
		gpio.m.Lock()
		defer gpio.m.Unlock()
		f, err := os.OpenFile(n, os.O_RDWR, 0)
		if err != nil {
			return nil, err
		}
		m.dev = &gpiodata{
			pad:  f,
			num:  num,
			data: uint32(base + imx6.GPIO_DR),
			dir:  uint32(base + imx6.GPIO_GDIR),
		}
		return m.dev, nil
	}); err != nil {
		return -1, fmt.Errorf("Can't set up %s: %w", m.num, err)
	}
	return len(in), nil
}

var _ syscall.DevFile = &gpioctl{}

type gpiodata struct {
	m    sync.Mutex
	pad  *os.File
	num  uint
	data uint32
	dir  uint32
}

// Pread reads gpiodata
func (m *gpiodata) Pread(out []byte, addr int64) (int, error) {
	return -1, fmt.Errorf("gpiodata read")
	m.m.Lock()
	defer m.m.Unlock()
	return -1, fmt.Errorf("not yet")
}

// Pwrite writes longwords to gpiodata
func (m *gpiodata) Pwrite(in []byte, addr int64) (int, error) {
	m.m.Lock()
	defer m.m.Unlock()
	if len(in) != 1 {
		return -1, fmt.Errorf(`usage: one-byte value, 0, 1, "0", "1", not %#x`, in)
	}
	switch in[0] {
	case 0, byte('0'):
	case 1, byte('1'):
	default:
		return -1, fmt.Errorf(`usage: 0, 1, "0", "1", not %#x`, in[0])
	}
	return len(in), nil

}

var _ syscall.DevFile = &gpiodata{}
