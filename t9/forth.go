// Copyright 2012-2017 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package t9

import (
	"fmt"
	"io/fs"
	"log"
	"strconv"

	"github.com/u-root/u-root/pkg/forth"
)

// The great panic discussion.
// Rob has given talks on using panic for parsers.
// I have talked to Russ about using panic for parsers.
// Short form: it's ok. In general, don't panic.
// But parsers are special: using panic
// in a parser makes the code tons cleaner.

func pushReg(f forth.Forth) {
	forth.Debug("pushReg")
	r := f.Pop().(FS)
	c, err := NewReg(r)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}
	f.Push(c)
}

// Note that if any type asserts fail the forth interpret loop catches
// it. It also catches stack underflow, all that stuff.
func open(f forth.Forth) {
	forth.Debug("open")
	g := f.Pop().(string)
	forth.Debug("%v", g)
	r := f.Pop().(FS)
	c, err := r.Open(g)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}
	f.Push(c)
}

func stat(f forth.Forth) {
	forth.Debug("stat")
	g := f.Pop().(string)
	forth.Debug("%v", g)
	r := f.Pop().(FS)
	c, err := r.Stat(g)
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
	g := f.Pop().(IO)
	f.Push(g)
	forth.Debug("%v", g)
	var b [8192]byte
	n, err := g.ReadAt(b[:], 0)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}
	f.Push(b[:n])
}

// write the string at TOS
func write(f forth.Forth) {
	forth.Debug("write")
	toString(f)
	b := f.Pop().(string)
	g := f.Pop().(IO)
	f.Push(g)
	log.Printf("=========> Write %q", b)
	if _, err := g.WriteAt([]byte(b), 0); err != nil {
		panic(fmt.Sprintf("%v", err))
	}
}

func pin(f forth.Forth) {
	forth.Debug("pin")
	i := f.Pop().(string)
	pno, err := strconv.Atoi(i)
	if err != nil {
		panic(fmt.Errorf("convering %s to pin:%w", i, err))
	}
	g := f.Pop().(*GPIO)
	f.Push(g)
	p, err := g.Init(pno)
	if err != nil {
		panic(err)
	}
	f.Push(p)
}

func value(f forth.Forth) {
	forth.Debug("value")
	p := f.Pop().(*Pin)
	f.Push(p)
	f.Push(p.Value())
}

func toString(f forth.Forth) {
	forth.Debug("string")
	g := f.Pop()
	switch v := g.(type) {
	case []byte:
		f.Push(string(v))
	case string:
		f.Push(v)
	default:
		if v, ok := g.(fs.FileInfo); ok {
			f.Push(fmt.Sprintf("%s: %v, %d bytes", v.Name(), v.Mode(), v.Size()))
			return
		}
		f.Push(fmt.Sprintf("Don't know how to string %T: best try:%v", v, v))
	}
}

type Connect func() (FS, error)

func New(c Connect) (forth.Forth, error) {
	// Get a connection to the FS for the name space.
	root, err := c()
	if err != nil {
		return nil, err
	}
	v("Attached %v", root)

	// Now attach all devices ...
	r, err := NewReg(root)
	if err != nil {
		log.Printf("non-fatal: %v", err)
		return nil, nil
	}

	// Clean this up for later.
	// The trick is the hierarchy of needs ...
	// wtf did I mean by that. Anyway, this sets up GPIO4 to be
	// read and written.
	GPIO4.reg = r

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
		{name: "create", op: func(f forth.Forth) {
			forth.Debug("create")
			for _, c := range creators {
				if d, err := c(root, r); err != nil {
					log.Printf("%v:%v", c, err)
				} else {
					log.Printf("%v:%v", c, d)
				}
			}
		},
		},
		{name: "gpio4", op: func(f forth.Forth) {
			forth.Debug("gpio4")
			f.Push(GPIO4)
		},
		},
		{name: "open", op: open},
		{name: "read", op: read},
		{name: "write", op: write},
		{name: "reg", op: pushReg},
		{name: "pin", op: pin},
		{name: "stat", op: stat},
		{name: "string", op: toString},
		{name: "value", op: value},
	} {
		forth.Putop(o.name, o.op)
	}

	return f, nil
}
