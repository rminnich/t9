// Copyright 2012-2017 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/u-root/u-root/pkg/forth"
	"github.com/usbarmory/tamago-example/t9"
	"harvey-os.org/ninep/protocol"
)

var (
	v = log.Printf
)

// The great panic discussion.
// Rob has given talks on using panic for parsers.
// I have talked to Russ about using panic for parsers.
// Short form: it's ok. In general, don't panic.
// But parsers are special: using panic
// in a parser makes the code tons cleaner.

// Note that if any type asserts fail the forth interpret loop catches
// it. It also catches stack underflow, all that stuff.
func open(f forth.Forth) {
	forth.Debug("open")
	g := f.Pop().(string)
	forth.Debug("%v", g)
	r := f.Pop().(t9.FS)
	c, err := r.Open(g)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}
	f.Push(c)
}

// read up to 8k of data, starting at 0.
// For now, files are so small that we don't bother with a loop.
// That's cheap insurance against someone doing something
// foolish.
func read(f forth.Forth) {
	forth.Debug("read")
	g := f.Pop().(t9.IO)
	f.Push(g)
	forth.Debug("%v", g)
	var b [8192]byte
	n, err := g.ReadAt(b[:], 0)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}
	f.Push(b[:n])
}

func toString(f forth.Forth) {
	forth.Debug("string")
	g := f.Pop().([]byte)
	f.Push(string(g))
}

func main() {
	var (
		v      = func(string, ...interface{}) {}
		logger = flag.Bool("l", false, "Enable client logging")
		debug  = flag.Bool("d", false, "enable debug prints")
		n      = flag.String("net", "tcp", "net type")
		aname  = flag.String("aname", "/", "attach name (i.e. root)")
	)

	flag.Parse()
	var opts []protocol.ClientOpt
	if *logger {

	}
	if *debug {
		v = log.Printf
	}
	a := flag.Args()
	if len(a) < 1 {
		log.Fatalf("usage: %s ipaddr [commands]", os.Args[0])
	}
	addr := a[0]
	v("Attach: (%q,%q):", *n, addr)
	root, err := t9.NewNinep(*n, a[0], *aname, opts...)
	if err != nil {
		log.Fatal(err)
	}
	v("Attached %v", root)
	f := forth.New()
	forth.Debug = log.Printf

	for _, o := range []struct {
		name string
		op   forth.Op
	}{
		{name: "root", op: func(f forth.Forth) {
			forth.Debug("root")
			f.Push(root)
		},
		},
		{name: "open", op: open},
		{name: "read", op: read},
		{name: "string", op: toString},
	} {
		forth.Putop(o.name, o.op)
	}

	for _, arg := range a[1:] {
		fmt.Printf("%sOK\n", f.Stack())
		if err := forth.EvalString(f, arg); err != nil {
			fmt.Printf("%v\n", err)
		}
	}
}
