package main

import (
	"fmt"
	"log"
	"time"
)

var MDI = log.Printf

// MXS is for the MXS registers
type MXS struct {
	addr uint32
	name string
	val  uint32 // value from the most recent read, convenience
	err  error  // This is an accumulated error.
	rw   rw
}

// NewMXS returns an MXS for a given base.
func NewMXS(rw rw, addr uint32, name string) *MXS {
	return &MXS{rw: rw, addr: addr, name: name}
}

func (m *MXS) Clr(bits uint32) *MXS {
	if m.err != nil {
		return m
	}
	m.err = writel(m.rw, bits, m.addr+8)
	return m

}

func (m *MXS) Set(bits uint32) *MXS {
	if m.err != nil {
		return m
	}
	m.err = writel(m.rw, bits, m.addr+4)
	return m
}

func (m *MXS) Read() *MXS {
	if m.err != nil {
		return m
	}
	m.val, m.err = readl(m.rw, m.addr)
	return m
}

func (m *MXS) Wait(bits uint32, timeout time.Duration) *MXS {
	if m.err != nil {
		return m
	}
	now := time.Now()
	for {
		m.Read()
		if m.err != nil {
			break
		}
		if m.val&bits == bits {
			break
		}
		if time.Since(now) > timeout {
			m.err = fmt.Errorf("Timed out on %s after %v", m.name, timeout)
			break
		}
	}
	return m
}
