// Copyright 2012 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package t9

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"os"

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
	fid  protocol.FID
	name string
}

var _ FS = &ninep{}
var _ IO = &file{}

func NewNinep(netname, addr, root string, opt ...protocol.ClientOpt) (*ninep, error) {
	conn, err := net.Dial(netname, addr)
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

func (n *ninep) Open(name string) (IO, error) {
	return nil, fmt.Errorf("open: not yet")
}

func (n *ninep) Stat(name string) (*os.FileInfo, error) {
	return nil, fmt.Errorf("stat: not yet")
}

func (f *file) ReadAt([]byte, int64) (int, error) {
	return -1, fmt.Errorf("readat: not yet")
}

func (f *file) WriteAt([]byte, int64) (int, error) {
	return -1, fmt.Errorf("writeat: not yet")
}
