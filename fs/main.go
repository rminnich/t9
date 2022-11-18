// Copyright 2012-2017 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"os"

	"github.com/usbarmory/tamago-example/t9"
	"harvey-os.org/ninep/protocol"
)

var v = log.Printf

func main() {
	var (
		v       = func(string, ...interface{}) {}
		logger  = flag.Bool("l", false, "Enable client logging")
		debug   = flag.Bool("d", false, "enable debug prints")
		n       = flag.String("net", "tcp", "net type")
		aname   = flag.String("aname", "/", "attach name (i.e. root)")
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
	/*
		for _, arg := range a[1:] {
			op, f := arg[0], arg[1:]
			switch op {
			default:
				log.Fatalf("%v unknown; only r or w", op)
			case 'r':
				b, err := read(v, root, f)
				if err != nil {
					log.Printf("Reading %v: got (%v, %v)", f, b, err)
					continue
				}
				os.Stdout.Write(b)
				// l is like r save we assume dirents.
			case 'w':
				n, err := write(v, root, f, os.Stdin)
				if err != nil {
					log.Printf("Writing %v: got (%v, %v)", f, n, err)
					continue
				}
				// l is like r save we assume dirents.
			case 'l':
				ents, err := readdir(v, root, f)
				if err != nil {
					log.Printf("%v:%v", f, err)
					continue
				}
				for _, d := range ents {
					fmt.Printf("%v\n", d)
				}
			}
		}
	*/
}
