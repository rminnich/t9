package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"syscall"
	"unsafe"
)

type byteMemory struct {
	addr   int64
	length int64
}

type shortMemory struct {
	addr   int64
	length int64
}

type longMemory struct {
	addr   int64
	length int64
}

func init() {
	// Looks like only word-aligned IO is allowed.
	if false {
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
	}
	if err := syscall.MkDev("/dev/long", 0666, func() (syscall.DevFile, error) {
		return &longMemory{
			addr:   0x80000000,
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

// Pread bytes.
func (u *byteMemory) Pread(out []byte, addr int64) (int, error) {
	log.Printf("Byte Pread %d bytes @ %#x", len(out), addr)
	if err := okaddr(addr, u.length, len(out), 1, false); err != nil {
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
	if err := okaddr(addr, u.length, len(in), 1, true); err != nil {
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
	if err := okaddr(addr, u.length, len(out), 2, false); err != nil {
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
	if err := okaddr(addr, u.length, len(in), 2, true); err != nil {
		return -1, err
	}

	for i := 0; i < len(in); i += 2 {
		w := binary.LittleEndian.Uint16(in[i:])
		*(*uint16)(unsafe.Pointer(uintptr(addr) + uintptr(i))) = w
	}
	return len(in), nil
}

var _ syscall.DevFile = &shortMemory{}

// Pread long, the address has to be in the range 0..u.length
func (u *longMemory) Pread(out []byte, addr int64) (int, error) {
	if err := okaddr(addr, u.length, len(out), 4, false); err != nil {
		return -1, err
	}
	for i := 0; i < len(out); i += 4 {
		log.Printf("Reading %#x: ", uintptr(addr+u.addr)+uintptr(i))
		w := *(*uint32)(unsafe.Pointer(uintptr(addr+u.addr) + uintptr(i)))
		log.Printf("Got %#x: ", w)
		binary.LittleEndian.PutUint32(out[i:], w)
	}
	return len(out), nil
}

// Pwrite memory, the address has to be in the range 0..u.length.
func (u *longMemory) Pwrite(in []byte, addr int64) (int, error) {
	if err := okaddr(addr, u.length, len(in), 4, true); err != nil {
		return -1, err
	}

	for i := 0; i < len(in); i += 4 {
		w := binary.LittleEndian.Uint32(in[i:])
		*(*uint32)(unsafe.Pointer(uintptr(u.addr+addr) + uintptr(i))) = w
	}
	return len(in), nil
}

var _ syscall.DevFile = &longMemory{}
