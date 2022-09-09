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
	rfid, fid := protocol.FID(0), root.GetFID()
	v("Walk fid %d to %d name %q", rfid, fid, f)
	w, err := root.CallTwalk(rfid, fid, filepath.SplitList(f))
	if err != nil {
		return nil, err
	}
	defer func(fid protocol.FID) {
		if err := root.CallTclunk(fid); err != nil {
			log.Printf("CallTclunk(%d) failed: %v", fid, err)
		}
	}(fid)
	v("Walk is %v", w)
	b, err := root.CallTstat(fid)
	if err != nil {
		return nil, fmt.Errorf("CallTstat(%d) failed: %v", fid, err)
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
