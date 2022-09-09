// Copyright 2012 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// attach
package main

import (
	"fmt"
	"net"

	"harvey-os.org/ninep/protocol"
)

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
