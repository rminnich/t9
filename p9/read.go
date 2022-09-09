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

func read(v func(string, ...interface{}), root *protocol.Client, f string) ([]byte, error) {
	rfid, fid := protocol.FID(0), root.GetFID()
	v("Walk fid %d to %d name %q", rfid, fid, f)
	w, err := root.CallTwalk(rfid, fid, filepath.SplitList(f))
	if err != nil {
		return nil, err
	}
	v("Walk is %v", w)

	defer func(fid protocol.FID) {
		if err := root.CallTclunk(fid); err != nil {
			log.Printf("CallTclunk(%d) failed: %v", fid, err)
		}
	}(fid)

	q, iounit, err := root.CallTopen(fid, 0)
	if err != nil {
		return nil, err
	}
	v("Open is %v %v", q, iounit)

	var off int64
	var out bytes.Buffer
	for {
		d, err := root.CallTread(fid, protocol.Offset(off), protocol.Count(iounit))
		v("Reading got %d bytes @ %d", len(d), off)
		if err != nil {
			return out.Bytes(), err
		}
		if len(d) == 0 {
			return out.Bytes(), nil
		}
		off += int64(len(d))
		out.Write(d)
	}
	return out.Bytes(), nil
}

func readdir(v func(string, ...interface{}), root *protocol.Client, f string) ([]protocol.Dir, error) {
	isdir, err := isDir(v, root, f)
	if err != nil {
		return nil, err
	}
	if !isdir {
		// wtf OSX
		return nil, fmt.Errorf("not a directory")
	}
	dat, err := read(v, root, f)
	if err != nil {
		return nil, err
	}
	b := bytes.NewBuffer(dat)
	var ents []protocol.Dir
	for b.Len() > 0 {
		d, err := protocol.Unmarshaldir(b)
		if err != nil {
			return ents, err
		}
		ents = append(ents, d)
	}
	return ents, nil
}
