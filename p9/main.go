// Copyright 2012-2017 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"

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
		mountIt = flag.Bool("m", false, "mount the file system?")
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
	conn, err := net.Dial(*n, addr)
	if err != nil {
		log.Fatalf("dial server (%q, %q): %v", *n, addr, err)
	}
	// go ahead and mount it too. But we'll leave the option to use commands.
	if *mountIt {
		fs, p, err := NewP9FS(conn, "/", time.Second, time.Second)
		if err != nil {
			log.Fatal(err)
		}
		v("fs: %v, %v, %v", fs, p, err)
	}
	root, err := attach(v, conn, *aname, opts...)
	if err != nil {
		log.Fatal(err)
	}
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
}
