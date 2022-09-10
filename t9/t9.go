package main

type T9 interface {
	Open(name string) (T9, error)
	Close() error
	Pread(out []byte, off int64) (int, error)
	Pwrite(dat []byte, off int64) (int, error)
}
