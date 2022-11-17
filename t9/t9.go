package main

import "io/fs"

type T9 interface {
	Open(name string) (T9, error)
	Close() error
	Pread(out []byte, off int64) (int, error)
	Pwrite(dat []byte, off int64) (int, error)
	Readdir(off int64) (fs.FileInfo, error)
}

var v = func(string, ...interface{}) {}
