// https://github.com/f-secure-foundry/tamago-example
//
// Copyright (c) F-Secure Corporation
// https://foundry.f-secure.com
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
