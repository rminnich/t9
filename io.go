package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"syscall"
	"unsafe"
)

type byteMemory struct {
}

type shortMemory struct {
}

type longMemory struct {
}

func init() {
	if err := syscall.MkDev("/dev/byte", 0666, func() (syscall.DevFile, error) {
		return &byteMemory{}, nil
	}); err != nil {
		log.Printf("Can't set up byte: %v", err)
	}
	if err := syscall.MkDev("/dev/short", 0666, func() (syscall.DevFile, error) {
		return &shortMemory{}, nil
	}); err != nil {
		log.Printf("Can't set up short: %v", err)
	}
	if err := syscall.MkDev("/dev/long", 0666, func() (syscall.DevFile, error) {
		return &longMemory{}, nil
	}); err != nil {
		log.Printf("Can't set up long: %v", err)
	}
}

func okaddr(addr int64, length int, write bool) error {
	if false {
		return fmt.Errorf("Bad address: %#x:%#x", addr, addr+int64(length))
	}
	return nil
}

// Pread bytes.
func (u *byteMemory) Pread(out []byte, addr int64) (int, error) {
	log.Printf("Byte Pread %d bytes @ %#x", len(out), addr)
	if err := okaddr(addr, len(out), false); err != nil {
		log.Printf("Return error %v", err)
		return -1, err
	}

	for i := range out {
		out[i] = *(*byte)(unsafe.Pointer(uintptr(addr) + uintptr(i)))
	}
	return len(out), nil
}

// Pwrite memory.
func (u *byteMemory) Pwrite(in []byte, addr int64) (int, error) {
	log.Printf("Byte Pwrite %d bytes @ %#x", len(in), addr)
	if err := okaddr(addr, len(in), true); err != nil {
		log.Printf("Return error %v", err)
		return -1, err
	}

	for i := range in {
		*(*byte)(unsafe.Pointer(uintptr(addr) + uintptr(i))) = in[i]
	}
	return len(in), nil
}

var _ syscall.DevFile = &byteMemory{}

// Pread shorts
func (u *shortMemory) Pread(out []byte, addr int64) (int, error) {
	if err := okaddr(addr, len(out), false); err != nil {
		return -1, err
	}
	for i := 0; i < len(out); i += 2 {
		w := *(*uint16)(unsafe.Pointer(uintptr(addr) + uintptr(i)))
		binary.LittleEndian.PutUint16(out[i:], w)
	}
	return len(out), nil
}

// Pwrite memory.
func (u *shortMemory) Pwrite(in []byte, addr int64) (int, error) {
	if err := okaddr(addr, len(in), true); err != nil {
		return -1, err
	}

	for i := 0; i < len(in); i += 2 {
		w := binary.LittleEndian.Uint16(in[i:])
		*(*uint16)(unsafe.Pointer(uintptr(addr) + uintptr(i))) = w
	}
	return len(in), nil
}

var _ syscall.DevFile = &shortMemory{}

// Pread long
func (u *longMemory) Pread(out []byte, addr int64) (int, error) {
	if err := okaddr(addr, len(out), false); err != nil {
		return -1, err
	}
	for i := 0; i < len(out); i += 4 {
		w := *(*uint32)(unsafe.Pointer(uintptr(addr) + uintptr(i)))
		binary.LittleEndian.PutUint32(out[i:], w)
	}
	return len(out), nil
}

// Pwrite memory.
func (u *longMemory) Pwrite(in []byte, addr int64) (int, error) {
	if err := okaddr(addr, len(in), true); err != nil {
		return -1, err
	}

	for i := 0; i < len(in); i += 2 {
		w := binary.LittleEndian.Uint32(in[i:])
		*(*uint32)(unsafe.Pointer(uintptr(addr) + uintptr(i))) = w
	}
	return len(in), nil
}

var _ syscall.DevFile = &longMemory{}
