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
	"fmt"
	"log"
	"syscall"
)

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
		m := &gpioctl{num: uint(len(gpio.ctl)), name: n}
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
