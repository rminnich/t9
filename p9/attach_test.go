// Copyright 2012 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// attach
package main

import (
	"net"
	"testing"

	"harvey-os.org/ninep/protocol"
	"harvey-os.org/ninep/ufs"
)

func TestAttach(t *testing.T) {
	p, p2 := net.Pipe()
	n, err := ufs.NewUFS("", 0, func(l *protocol.NetListener) error {
		l.Trace = t.Logf
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}

	if err := n.Accept(p2); err != nil {
		t.Fatalf("Accept: want nil, got %v", err)
	}

	if _, err := attach(t.Logf, p, "/"); err != nil {
		t.Fatalf("attach(%v, /): got %v, want nil", p, err)
	}
}
