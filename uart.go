// https://github.com/f-secure-foundry/tamago-example
//
// Copyright (c) F-Secure Corporation
// https://foundry.f-secure.com
//
// Use of this source code is governed by the license
// that can be found in the LICENSE file.

// Basic test example for tamago/arm running on supported i.MX6 targets.

package main

import (
	"log"
	"syscall"

	"github.com/f-secure-foundry/tamago/soc/imx6"
)

type uart struct {
	in  chan byte
	out chan byte
	u   *imx6.UART
}

func init() {
	u, err := NewUART(imx6.UART1)
	if err == nil {
		err := syscall.MkDev("/dev/console", 0666, func() (syscall.DevFile, error) {
			return u, nil
		})
		if err != nil {
			log.Printf("Can't set up console: %v", err)
		}
	}
}

func NewUART(r *imx6.UART) (*uart, error) {
	u := &uart{in: make(chan byte, 128), out: make(chan byte, 128), u: r}
	return u, nil
}

// Pread the UART. This will block, and depends on the underlying bit polling
// to yield. It will return with the first []byte that is non-zero. It
// can return more than one byte.
// For now, just have it read one byte.
func (u *uart) Pread(out []byte, _ int64) (int, error) {
	var n int
	var b [1]byte
	for n < len(b) {
		n = u.u.Read(b[:])
		if n > 0 {
			u.u.Write(b[:n])
			copy(out, b[:n])
		}
	}
	return n, nil
}

// Pwrite the UART. This will block, and depends on the underlying bit polling
// to yield.
func (u *uart) Pwrite(b []byte, _ int64) (int, error) {
	u.u.Write(b[:])
	return len(b), nil
}

var _ syscall.DevFile = &uart{}
