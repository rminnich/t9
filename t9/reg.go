// https://github.com/usbarmory/tamago
//
// Copyright (c) WithSecure Corporation
// https://foundry.withsecure.com
//
// Use of this source code is governed by the license
// that can be found in the LICENSE file.

// Package reg provides primitives for retrieving and modifying hardware
// registers.
//
// This package is only meant to be used with `GOOS=tamago` as supported by the
// TamaGo framework for bare metal Go on ARM/RISC-V SoCs, see
// https://github.com/usbarmory/tamago.
package t9

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"math"
	"runtime"
	"time"
)

// HACK because you can have Read/ReadAt, but not Write/WriteAt
// annoying that this is still not a thing.
type w struct {
	off int64
	f   IO
}

func (w *w) Write(b []byte) (int, error) {
	return w.f.WriteAt(b, w.off)
}

// end HACK, just ignore it.

type reg struct {
	name string
	fd   IO
}

// NewReg returns an IO usable for registers.
func NewReg(t9 FS) (*reg, error) {
	fd, err := t9.Open("/dev/reg32")
	if err != nil {
		log.Printf("/dev/reg32: %v", err)
		return nil, fmt.Errorf("/dev/reg32:%v", err)
	}
	return &reg{name: "/dev/reg32", fd: fd}, nil
}

func (r *reg) Read(addr uint32) uint32 {
	l := uint32(math.MaxUint32)
	// reading from "air" always gives all 1s on most systems.
	if err := binary.Read(io.NewSectionReader(r.fd, 4, int64(addr)), binary.LittleEndian, &l); err != nil {
		log.Printf("%s.Read(%#x): %v", r.name, addr)
	}

	return l
}

func (r *reg) Write(addr uint32, val uint32) {
	ws := &w{off: int64(addr), f: r.fd}
	if err := binary.Write(ws, binary.LittleEndian, &val); err != nil {
		log.Printf("%s.Write(%#x, %#x): %v", r.name, addr, val)
	}
}

func (r *reg) Get(addr uint32, pos int, mask int) uint32 {
	l := r.Read(addr)
	return uint32((l >> uint32(pos)) & uint32(mask))
}

func (r *reg) Set(addr uint32, pos int) {
	l := r.Read(addr)
	l |= (1 << pos)

	r.Write(addr, l)
}

func (r *reg) Clear(addr uint32, pos int) {
	l := r.Read(addr)
	l &= ^(1 << pos)

	r.Write(addr, l)
}

func (r *reg) SetTo(addr uint32, pos int, val bool) {
	if val {
		r.Set(addr, pos)
	} else {
		r.Clear(addr, pos)
	}
}

func (r *reg) SetN(addr uint32, pos int, mask int, val uint32) {
	l := r.Read(addr)
	l = (l & (^(uint32(mask) << pos))) | (val << pos)
	r.Write(addr, l)
}

func (r *reg) ClearN(addr uint32, pos int, mask int) {
	l := r.Read(addr)
	l &= ^(uint32(mask) << pos)
	r.Write(addr, l)
}

func (r *reg) WriteBack(addr uint32) {
	r.Write(addr, r.Read(addr))
}

func (r *reg) Or(addr uint32, val uint32) {
	r.Write(addr, r.Read(addr)|val)
}

// Wait waits for a specific register bit to match a value. This function
// cannot be used before runtime initialization with `GOOS=tamago`.
func (r *reg) Wait(addr uint32, pos int, mask int, val uint32) {
	for r.Get(addr, pos, mask) != val {
		// tamago is single-threaded, give other goroutines a chance
		runtime.Gosched()
	}
}

// WaitFor waits, until a timeout expires, for a specific register bit to match
// a value. The return boolean indicates whether the wait condition was checked
// (true) or if it timed out (false). This function cannot be used before
// runtime initialization.
func (r *reg) WaitFor(timeout time.Duration, addr uint32, pos int, mask int, val uint32) bool {
	start := time.Now()

	for r.Get(addr, pos, mask) != val {
		// tamago is single-threaded, give other goroutines a chance
		runtime.Gosched()

		if time.Since(start) >= timeout {
			return false
		}
	}

	return true
}

// WaitSignal waits, until a channel is closed, for a specific register bit to
// match a value. The return boolean indicates whether the wait condition was
// checked (true) or cancelled (false). This function cannot be used before
// runtime initialization.
func (r *reg) WaitSignal(done chan bool, addr uint32, pos int, mask int, val uint32) bool {
	for r.Get(addr, pos, mask) != val {
		// tamago is single-threaded, give other goroutines a chance
		runtime.Gosched()

		select {
		case <-done:
			return false
		default:
		}
	}

	return true
}
