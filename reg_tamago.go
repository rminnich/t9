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
package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"
	"syscall"

	"github.com/usbarmory/tamago/board/usbarmory/mk2"
)

type reg32File struct{}

func init() {
	log.Printf("=======> t9 device reg32")
	err := syscall.MkDev("/dev/reg32", 0666, openReg32)
	log.Printf("err %v", err)
}

func openReg32() (syscall.DevFile, error) {
	return reg32File{}, nil
}

func (f reg32File) Close() error {
	return nil
}

var ErrNotAligned = errors.New("Not aligned")

func check(size int, offset int64) error {
	// The offset must be aligned x4
	if (offset & 3) != 0 {
		return fmt.Errorf("offset %#x: %w", ErrNotAligned)
	}
	// The size must be aligned x4
	if (size & 3) != 0 {
		return fmt.Errorf("size %#x: %w", ErrNotAligned)
	}
	return nil
}

func (f reg32File) Pread(b []byte, offset int64) (int, error) {
	if err := check(len(b), offset); err != nil {
		return -1, err
	}
	var longs = make([]uint32, len(b)/4)
	// read them from the registers ...
	binary.Write(bytes.NewBuffer(b), binary.LittleEndian, longs)

	return len(b), nil
}

func (f reg32File) Pwrite(b []byte, offset int64) (int, error) {
	if err := check(len(b), offset); err != nil {
		return -1, err
	}
	var longs = make([]uint32, len(b)/4)
	binary.Read(bytes.NewBuffer(b), binary.LittleEndian, longs)
	return len(b), nil
}

type ledFile struct {
	blue, white string
}

var leds = &ledFile{
	blue:  "on",
	white: "on",
}

func init() {
	syscall.MkDev("/dev/led", 0666, openled)
}

func openled() (syscall.DevFile, error) {
	return leds, nil
}

func (f ledFile) Close() error {
	return nil
}

func (f ledFile) Pread(b []byte, offset int64) (int, error) {
	n, err := bytes.NewReader([]byte(fmt.Sprintf(`{"blue": %q,"white": %q}`, f.blue, f.white))).ReadAt(b, offset)
	if err == io.EOF && n > 0 {
		err = nil
	}
	return n, err
}

func (f ledFile) Pwrite(b []byte, offset int64) (int, error) {
	cmd := strings.Fields(string(b))
	if len(cmd) != 2 {
		return -1, fmt.Errorf("usage: blue|white on|off")
	}
	var onoff bool
	switch cmd[1] {
	case "on":
		onoff = true
	case "off":
	default:
		return -1, fmt.Errorf("usage: blue|white on|off")
	}
	err := mk2.LED(cmd[0], onoff)
	return len(b), err

}
