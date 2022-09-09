// Copyright 2012 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// attach
package main

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net"
	"path/filepath"
	"testing"

	"harvey-os.org/ninep/protocol"
	"harvey-os.org/ninep/ufs"
)

func TestRead(t *testing.T) {
	tmp := t.TempDir()
	if err := ioutil.WriteFile(filepath.Join(tmp, "1"), []byte("123"), 0644); err != nil {
		t.Fatalf(`"ioutil.WriteFile(filepath.Join(tmp, "1"), []byte("123"), 0644): %v != nil`, err)
	}
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

	root, err := attach(t.Logf, p, tmp)
	if err != nil {
		t.Fatal(err)
	}

	b, err := read(t.Logf, root, "1")
	if err != nil && !errors.Is(err, io.EOF) {
		t.Fatal(err)
	}
	if string(b) != "123" {
		t.Fatalf("%s: %q != %q", filepath.Join(tmp, "1"), string(b), "123")
	}
}

func TestWrite(t *testing.T) {
	tmp := t.TempDir()
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

	root, err := attach(t.Logf, p, tmp)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := write(t.Logf, root, "1", bytes.NewBufferString("123")); err != nil {
		t.Fatal(err)
	}
	b, err := ioutil.ReadFile(filepath.Join(tmp, "1"))
	if err != nil {
		t.Fatalf(`"ioutil.ReadFile(filepath.Join(tmp, "1"):%v != nil`, err)
	}
	if string(b) != "123" {
		t.Fatalf("%s: %q != %q", filepath.Join(tmp, "1"), string(b), "123")
	}
}
