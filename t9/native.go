// Copyright 2012 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package t9

import (
	"io/fs"
	"os"
	"path/filepath"
	"syscall"
	"time"
)

// nfs is the native fs.
type nfs struct {
	root string
}

type nativefile struct {
	*os.File
}

var _ FS = &nfs{}

func NewNativeFS(root string) (*nfs, error) {
	return &nfs{root: root}, nil
}

func (*nfs) Close() error {
	return nil
}

// Pread implements PRead.
func (f *nativefile) Pread(b []byte, off int64) (int, error) {
	return f.File.ReadAt(b, off)
}

// Pwrite implements PWrite.
func (f *nativefile) Pwrite(b []byte, off int64) (int, error) {
	return f.File.WriteAt(b, off)
}

// Open implements os.Open. Always take abs paths,
// the Tamago environment is not complex enough
// to warrant anything else.
func (root *nfs) Open(name string) (syscall.DevFile, error) {
	f, err := os.Open(filepath.Join(root.root, name))
	if err != nil {
		return nil, err
	}
	return &nativefile{File: f}, nil
}

func (root *nfs) Stat(name string) (os.FileInfo, error) {
	return os.Stat(filepath.Join(root.root, name))
}

// implement fs.FileInfo
func (f *nativefile) Name() string {
	return f.Name()
}

var _ syscall.DevFile = &nativefile{}

func (f *nativefile) Size() int64 {
	fi, err := f.Stat()
	if err != nil {
		return -1
	}
	return fi.Size()
}

func (f *nativefile) Mode() fs.FileMode {
	fi, err := f.Stat()
	if err != nil {
		return 0
	}
	return fi.Mode().Perm()
}

func (f *nativefile) ModTime() time.Time {
	fi, err := f.Stat()
	if err != nil {
		return time.Unix(0, 0)
	}
	return fi.ModTime()
}

func (f *nativefile) IsDir() bool {
	fi, err := f.Stat()
	if err != nil {
		return false
	}
	return fi.IsDir()
}

func (f *nativefile) Sys() any {
	fi, err := f.Stat()
	if err != nil {
		return nil
	}
	return fi.Sys()
}

var _ fs.FileInfo = &nativefile{}
