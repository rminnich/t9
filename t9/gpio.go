// NXP GPIO support
// https://github.com/usbarmory/tamago
//
// Copyright (c) WithSecure Corporation
// https://foundry.withsecure.com
//
// Use of this source code is governed by the license
// that can be found in the LICENSE file.

// Package gpio implements helpers for GPIO configuration on NXP SoCs.
//
// This package is only meant to be used with `GOOS=tamago GOARCH=arm` as
// supported by the TamaGo framework for bare metal Go on ARM SoCs, see
// https://github.com/usbarmory/tamago.
package t9

import (
	"errors"
	"fmt"
)

// GPIO registers
const (
	GPIO_DR   = 0x00
	GPIO_GDIR = 0x04
)

// GPIO controller instance
type GPIO struct {
	// Controller index
	Index int
	// Base register
	Base uint32
	// Clock gate register
	CCGR uint32
	// Clock gate
	CG int

	clk bool

	reg *reg
}

// Pin instance
type Pin struct {
	num  int
	data uint32
	dir  uint32
	gpio *GPIO
}

// Init initializes a GPIO.
func (hw *GPIO) Init(num int) (*Pin, error) {
	if hw.Base == 0 || hw.CCGR == 0 {
		return nil, errors.New("invalid GPIO controller instance")
	}

	if num > 31 {
		return nil, fmt.Errorf("invalid GPIO number %d", num)
	}

	pin := &Pin{
		num:  num,
		data: hw.Base + GPIO_DR,
		dir:  hw.Base + GPIO_GDIR,
		gpio: hw,
	}

	if !hw.clk {
		// enable clock
		pin.gpio.reg.SetN(hw.CCGR, hw.CG, 0b11, 0b11)
		hw.clk = true
	}

	return pin, nil
}

// Out configures a PIN as output.
func (pin *Pin) Out() {
	pin.gpio.reg.Set(pin.dir, pin.num)
}

// In configures a PIN as input.
func (pin *Pin) In() {
	pin.gpio.reg.Clear(pin.dir, pin.num)
}

// High configures a PIN signal as high.
func (pin *Pin) High() {
	pin.gpio.reg.Set(pin.data, pin.num)
}

// Low configures a PIN signal as low.
func (pin *Pin) Low() {
	pin.gpio.reg.Clear(pin.data, pin.num)
}

// Value returns the PIN signal level.
func (pin *Pin) Value() (high bool) {
	return pin.gpio.reg.Get(pin.data, pin.num, 1) == 1
}
