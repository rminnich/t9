package main

import (
	"encoding/binary"
	"io"
	"log"
)

var FDI = log.Printf

//var rstack = debug.Stack
//var wstack = debug.Stack

var rstack = func() string { return "readl" }
var wstack = func() string { return "writel" }

// the ARM host is little endian.
func readl(r io.ReaderAt, o uint32) (uint32, error) {
	if o == 0 {
		panic("writel o is 0")
	}
	var b [4]byte
	var err error
	var l uint32
	if _, err = r.ReadAt(b[:], int64(o)); err != nil {
		FDI("%s(%#x) -> %#x, %v\n", rstack(), o, l, err)
		return 0, err
	}
	l = binary.LittleEndian.Uint32(b[:])
	FDI("%s(%#x) -> %#x, %v\n", rstack(), o, l, err)
	return l, nil
}

func writel(w io.WriterAt, v, o uint32) error {
	if o == 0 {
		panic("writel o is 0")
	}
	var b [4]byte
	var err error
	defer FDI("%s(%#x,%#x) -> %v\n", wstack(), v, o, err)
	binary.LittleEndian.PutUint32(b[:], v)
	if _, err = w.WriteAt(b[:], int64(o)); err != nil {
		return err
	}

	return nil
}

type rw interface {
	io.ReaderAt
	io.WriterAt
}

func bitset(rw rw, v, o uint32) error {
	reg, err := readl(rw, o)
	if err != nil {
		return err
	}
	reg |= v
	return writel(rw, reg, o)
}

func bitclr(rw rw, v, o uint32) error {
	reg, err := readl(rw, o)
	if err != nil {
		return err
	}
	reg &= ^v
	return writel(rw, reg, o)
}

func bitsetclr(rw rw, o, clr, set uint32) error {
	reg, err := readl(rw, o)
	if err != nil {
		return err
	}
	reg &= ^clr
	reg |= set
	return writel(rw, reg, o)

}
