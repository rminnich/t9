// Copyright 2012 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !tamago
// +build !tamago

package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"path/filepath"

	"harvey-os.org/ninep/protocol"
)

var niner struct {
	conn   net.Conn
	client *protocol.Client
	root   protocol.FID
}

type NineFile struct {
	FID    protocol.FID
	QID    protocol.QID
	iounit int
}

func attach(v func(string, ...interface{}), conn net.Conn, root string, opt ...protocol.ClientOpt) (*protocol.Client, error) {
	v("attach %v", conn)
	c, err := protocol.NewClient(func(c *protocol.Client) error {
		c.FromNet, c.ToNet = conn, conn
		return nil
	},
		func(c *protocol.Client) error {
			c.Msize = 8192
			c.Trace = v
			return nil
		})
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	msize, vers, err := c.CallTversion(8000, "9P2000")
	if err != nil {
		return nil, fmt.Errorf("CallTversion: want nil, got %v", err)
	}
	v("CallTversion: msize %v version %v", msize, vers)
	if _, err := c.CallTattach(0, protocol.NOFID, "", root); err != nil {
		return nil, err
	}
	return c, nil
}

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

func open(f string) (*NineFile, error) {
	rfid, fid := niner.root, niner.client.GetFID()
	v("Walk fid %d to %d name %q", rfid, fid, f)
	w, err := niner.client.CallTwalk(rfid, fid, filepath.SplitList(f))
	if err != nil {
		return nil, err
	}
	v("Walk is %v", w)

	q, iounit, err := niner.client.CallTopen(fid, 0)
	if err != nil {
		return nil, err
	}
	v("Open is %v %v", q, iounit)
	return &NineFile{FID: fid, iounit: int(iounit), QID: q}, nil
}

func (n *NineFile) Open(f string) (T9, error) {
	fid := niner.GetFID()
	v("Open fid %d name %q", fid, f)
	w, err := niner.client.CallTwalk(n.FID, fid, filepath.SplitList(f))
	if err != nil {
		return nil, err
	}
	v("Walk is %v", w)

	q, iounit, err := niner.client.CallTopen(fid, 0)
	if err != nil {
		return nil, err
	}
	v("Open is %v %v", q, iounit)
	return &NineFile{FID: fid, iounit: int(iounit), QID: q}, nil
}

func close(fid protocol.FID) error {
	return niner.client.CallTclunk(fid)
}

func (n *NineFile) Close() error {
	return close(n.FID)
}

func (n *NineFile) Pread(out []byte, off int64) (int, error) {
	var tot int
	for {
		d, err := niner.client.CallTread(n.FID, protocol.Offset(off), protocol.Count(n.iounit))
		v("Reading got %d bytes @ %d", len(d), off)
		if err != nil {
			return tot, err
		}
		if len(d) == 0 {
			return tot, nil
		}
		off += int64(len(d))
		out = append(out, d...)
	}
	return tot, nil
}

func (n *NineFile) Pwrite(dat []byte, off int64) (int, error) {
	var tot int
	b := make([]byte, n.iounit, n.iounit)
	r := bytes.NewReader(dat)
	for {
		amt, err := r.Read(b)
		if err != nil && !errors.Is(err, io.EOF) {
			return tot, err
		}
		if amt == 0 {
			return tot, nil
		}
		tn, err := niner.client.CallTwrite(protocol.FID(n.FID), protocol.Offset(off), b[:amt])
		v("Writeing got %d bytes @ %d", amt, off)
		if err != nil {
			return tot, err
		}
		tot += int(tn)
		off += int64(tn)
	}
	return tot, nil
}

var _ T9 = &NineFile{}
