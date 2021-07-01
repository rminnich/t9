package main

import (
	"encoding/binary"
	"io"
	"log"
	"runtime/debug"
)

var FDI = log.Printf
var stack = debug.Stack

// the ARM host is little endian.
func readl(r io.ReaderAt, o uint32) (uint32, error) {
	var b [4]byte
	var err error
	var l uint32
	defer FDI("%s(%#x) -> %#x, %v", stack(), o, l, err)
	if _, err = r.ReadAt(b[:], int64(o)); err != nil {
		return 0, err
	}
	l = binary.LittleEndian.Uint32(b[:])
	return l, nil
}

func writel(w io.WriterAt, v, o uint32) error {
	var b [4]byte
	var err error
	defer FDI("%s(%#x,%#x) -> %v", stack(), v, o, err)
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
