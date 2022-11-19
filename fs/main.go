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
)

func main() {
	var (
		v      = func(string, ...interface{}) {}
		logger = flag.Bool("l", false, "Enable client logging")
		debug  = flag.Bool("d", false, "enable debug prints")
		n      = flag.String("net", "tcp", "net type")
		aname  = flag.String("aname", "/", "attach name (i.e. root)")
	)

	flag.Parse()
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
	f, err := t9.New(*n, addr, *aname)
	if err != nil {
		log.Fatal(err)
	}
	v("Attach: (%q,%q):", *n, addr)
	for _, arg := range a[1:] {
		fmt.Printf("%sOK\n", f.Stack())
		if err := forth.EvalString(f, arg); err != nil {
			fmt.Printf("%v\n", err)
		}
	}
}
