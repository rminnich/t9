// Copyright 2022 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package t9

import "os"

type FS interface {
	Open(name string) (*File, error)
}

type File interface {
	Stat() (os.FileInfo, error)
	ReadAt([]byte, int64) (int, error)
	WriteAt([]byte, int64) (int, error)
	Close() error
}

var v = func(string, ...interface{}) {}
