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
	"fmt"
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
	u := &uart{in: make(chan byte), out: make(chan byte), u: r}
	go func(u *uart) {
		for {
			var b [1]byte
			n := u.u.Read(b[:])
			if n > 0 {
				u.in <- b[0]
			}
		}

	}(u)
	go func(u *uart) {
		for {
			var b [1]byte
			b[0] = <-u.out
			u.u.Write(b[:])
		}
	}(u)
	return u, nil
}

func (u *uart) Pread(b []byte, _ int64) (int, error) {
	b[0] = <-u.in
	return 1, nil
}

func (u *uart) Pwrite(b []byte, _ int64) (int, error) {
	return -1, fmt.Errorf("NFW")
}

var _ syscall.DevFile = &uart{}
