// Copyright 2022 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package t9

import "os"

// Meta is the interface for operations on metadata
// These operations work whether the file is open
// or not.
type Meta interface {
}

// FS is the file system interface. Its primary function
// is to gain access to files.
type FS interface {
	Open(name string) (*IO, error)
	Stat() (os.FileInfo, error)
}

// IO is the interface for IO on files.
type IO interface {
	ReadAt([]byte, int64) (int, error)
	WriteAt([]byte, int64) (int, error)
}

// Closer is for files that can/must be closed.
type Closer interface {
	Close() error
}

var v = func(string, ...interface{}) {}
