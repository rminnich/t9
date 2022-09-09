// Copyright 2012 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"errors"
	"io"
	"log"
	"path/filepath"

	"harvey-os.org/ninep/protocol"
)

func write(v func(string, ...interface{}), root *protocol.Client, f string, r io.Reader) (int, error) {
	rfid, fid := protocol.FID(0), root.GetFID()
	v("Walk fid %d to %d name %q", rfid, fid, f)
	dir, base := filepath.Split(f)
	w, err := root.CallTwalk(rfid, fid, filepath.SplitList(dir))
	if err != nil {
		return -1, err
	}
	v("Walk is %v", w)

	defer func(clunkfid protocol.FID) {
		if err := root.CallTclunk(clunkfid); err != nil {
			log.Printf("CallTclunk(%d) failed: %v", clunkfid, err)
		}
	}(fid)

	q, iounit, err := root.CallTcreate(fid, base, 0777, 2)
	if err != nil {
		return -1, err
	}
	v("Open is %v %v", q, iounit)

	var off int64
	var tot int
	b := make([]byte, iounit, iounit)
	for {
		n, err := r.Read(b)
		if err != nil && !errors.Is(err, io.EOF) {
			return tot, err
		}
		if n == 0 {
			return tot, nil
		}
		tn, err := root.CallTwrite(protocol.FID(fid), protocol.Offset(off), b[:n])
		v("Writeing got %d bytes @ %d", n, off)
		if err != nil {
			return tot, err
		}
		tot += int(tn)
		off += int64(tn)
	}
	return tot, nil
}
