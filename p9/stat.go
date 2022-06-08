// Copyright 2012 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"fmt"
	"log"
	"path/filepath"

	"harvey-os.org/ninep/protocol"
)

func stat(v func(string, ...interface{}), root *protocol.Client, f string) ([]byte, error) {
	v("Walk fid %d to %d name %q", 0, 1, f)
	w, err := root.CallTwalk(0, 1, filepath.SplitList(f))
	if err != nil {
		return nil, err
	}
	defer func(fid protocol.FID) {
		if err := root.CallTclunk(fid); err != nil {
			log.Printf("CallTclunk(%d) failed: %v", fid, err)
		}
	}(1)
	v("Walk is %v", w)
	b, err := root.CallTstat(1)
	if err != nil {
		return nil, fmt.Errorf("CallTstat(1) failed: %v", err)
	}
	return b, nil
}

func isDir(v func(string, ...interface{}), root *protocol.Client, f string) (bool, error) {
	b, err := stat(v, root, f)
	if err != nil {
		return false, err
	}
	d, err := protocol.Unmarshaldir(bytes.NewBuffer(b))
	if err != nil {
		return false, err
	}
	return (d.QID.Type & protocol.QTDIR) != 0, nil
}
