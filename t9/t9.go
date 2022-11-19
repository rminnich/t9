// Copyright 2022 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package t9

import "io/fs"

// FS is the file system interface. Its primary function
// is to gain access to files.
type FS interface {
	Open(name string) (IO, error)
	Stat(name string) (fs.FileInfo, error)
	Close() error
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
