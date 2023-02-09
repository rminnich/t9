// Copyright 2012 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package t9

import (
	"bytes"
	"fmt"
	"io/fs"
	"log"
	"net"
	"os"
	"path/filepath"
	"syscall"
	"time"

	"harvey-os.org/ninep/protocol"
)

type ninep struct {
	cl   *protocol.Client
	net  string
	addr string
	root string
	fid  protocol.FID
	fi   os.FileInfo
	dent protocol.Dir
}

type file struct {
	fid    protocol.FID
	qid    protocol.QID
	iounit protocol.MaxSize
	name   string
	dent   protocol.Dir
	root   *ninep
}

var _ FS = &ninep{}

func NewNinep(netname, addr, root string, opt ...protocol.ClientOpt) (*ninep, error) {
	conn, err := net.DialTimeout(netname, addr, 5*time.Second)
	if err != nil {
		log.Fatalf("dial server (%q, %q): %v", netname, addr, err)
	}
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
	var fid protocol.FID
	// It should be possible to stat the root.
	b, err := c.CallTstat(fid)
	if err != nil {
		return nil, fmt.Errorf("CallTstat(%d) failed: %v", fid, err)
	}

	d, err := protocol.Unmarshaldir(bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	return &ninep{
		cl:   c,
		net:  netname,
		addr: addr,
		root: root,
		dent: d,
	}, nil

}

func (n *ninep) Close() error {
	return nil
}

// Open implements os.Open. Always take abs paths,
// the Tamago environment is not complex enough
// to warrant anything else.
func (root *ninep) Open(name string) (syscall.DevFile, error) {
	rfid, fid := root.fid, root.cl.GetFID()
	v("Walk fid %d to %d name %q", rfid, fid, name)
	w, err := root.cl.CallTwalk(rfid, fid, filepath.SplitList(name))
	if err != nil {
		return nil, fmt.Errorf("%v: walk fid %d to %v: %v", root.cl, fid, name, err)
	}
	v("Walk is %v", w)

	q, iounit, err := root.cl.CallTopen(fid, 2)
	if err != nil {
		return nil, fmt.Errorf("%v: open fid %d for %v: %v", root.cl, fid, name, err)
	}
	v("Open is %v %v", q, iounit)
	return &file{
		fid:    fid,
		qid:    q,
		iounit: iounit,
		name:   name,
		root:   root,
	}, nil
}

func (root *ninep) Stat(name string) (os.FileInfo, error) {
	rfid, fid := root.fid, root.cl.GetFID()
	v("Walk fid %d to %d name %q", rfid, fid, name)
	w, err := root.cl.CallTwalk(rfid, fid, filepath.SplitList(name))
	if err != nil {
		return nil, fmt.Errorf("%v: walk fid %d to %v: %v", root.cl, fid, name, err)
	}
	v("Walk is %v", w)
	defer func(fid protocol.FID) {
		if err := root.cl.CallTclunk(fid); err != nil {
			log.Printf("CallTclunk(%d) failed: %v", fid, err)
		}
	}(fid)
	v("Walk is %v", w)
	b, err := root.cl.CallTstat(fid)
	if err != nil {
		return nil, fmt.Errorf("CallTstat(%d) failed: %v", fid, err)
	}

	d, err := protocol.Unmarshaldir(bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	return &file{
		fid:  fid,
		qid:  w[len(w)-1],
		name: name,
		root: root,
		dent: d,
	}, nil
}

func (f *file) ReadAt(b []byte, off int64) (int, error) {
	d, err := f.root.cl.CallTread(f.fid, protocol.Offset(off), protocol.Count(f.iounit))
	v("Reading got %d bytes @ %d", len(d), off)
	if err != nil {
		return -1, err
	}
	n := copy(b, d)
	return n, nil
}

func (f *file) WriteAt(b []byte, off int64) (int, error) {
	n, err := f.root.cl.CallTwrite(f.fid, protocol.Offset(off), b)
	v("Wrieing got %d bytes @ %d", n, off)
	return int(n), err
}

// Pread implements Pread
func (f *file) Pread(b []byte, off int64) (int, error) {
	return f.ReadAt(b, off)
}

func (f *file) Pwrite(b []byte, off int64) (int, error) {
	return f.WriteAt(b, off)
}

// implement fs.FileInfo
func (f *file) Name() string {
	return f.name
}

func (f *file) Size() int64 {
	return int64(f.dent.Length)
}

func (f *file) Mode() fs.FileMode {
	return fs.FileMode(f.dent.Mode).Perm()
}

func (f *file) ModTime() time.Time {
	return time.Unix(int64(f.dent.Mtime), 0)
}

func (f *file) IsDir() bool {
	return fs.FileMode(f.dent.Mode).IsDir()
}

func (f *file) Sys() any {
	return &f.dent
}

var _ fs.FileInfo = &file{}
