package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"syscall"
	"unsafe"
)

type longMemory struct {
	// user offsets are adjusted by 'adjust'
	adjust int64

	// adjusted addresses must be in mem[base:length]
	base   int64
	length int64
}

func init() {
	if err := syscall.MkDev("/dev/long", 0666, func() (syscall.DevFile, error) {
		return &longMemory{
			adjust: 0x80000000,
			base:   0x80000000,
			length: 0x20000000,
		}, nil
	}); err != nil {
		log.Printf("Can't set up long: %v", err)
	}
}

func okaddr(addr, size int64, length, align int, write bool) error {
	log.Printf("okaddr: %#x %#x %#x write %v", addr, length, align, write)
	if (length & (align - 1)) != 0 {
		return fmt.Errorf("Length %d is not %d-aligned", length, align)
	}
	if (addr & int64(align-1)) != 0 {
		return fmt.Errorf("address %#x is not %d-aligned", addr, align)
	}
	end := addr + int64(length) - 1
	if end > size {
		return fmt.Errorf("Bad address: %#x > %#x", end, size)
	}
	log.Printf("okaddr: is ok")
	return nil
}

// Pread reads longwords from memory.
func (u *longMemory) Pread(out []byte, addr int64) (int, error) {
	addr += u.adjust
	if err := okaddr(addr, u.length, len(out), 4, false); err != nil {
		return -1, err
	}
	for i := 0; i < len(out); i += 4 {
		log.Printf("Reading %#x: ", uintptr(addr)+uintptr(i))
		w := *(*uint32)(unsafe.Pointer(uintptr(addr) + uintptr(i)))
		log.Printf("Got %#x: ", w)
		binary.LittleEndian.PutUint32(out[i:], w)
	}
	return len(out), nil
}

// Pwrite writes longwords to memory.
func (u *longMemory) Pwrite(in []byte, addr int64) (int, error) {
	addr += u.adjust
	if err := okaddr(addr, u.length, len(in), 4, true); err != nil {
		return -1, err
	}

	for i := 0; i < len(in); i += 4 {
		w := binary.LittleEndian.Uint32(in[i:])
		*(*uint32)(unsafe.Pointer(uintptr(addr) + uintptr(i))) = w
	}
	return len(in), nil
}

var _ syscall.DevFile = &longMemory{}
