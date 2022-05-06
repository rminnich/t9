// https://github.com/usbarmory/tamago-example
//
// Copyright (c) WithSecure Corporation
// https://foundry.withsecure.com
//
// Use of this source code is governed by the license
// that can be found in the LICENSE file.

package main

import (
	_ "net/http/pprof"

	"github.com/arl/statsviz"
)

func init() {
	statsviz.RegisterDefault()
}
