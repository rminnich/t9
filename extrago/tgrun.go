package main

import (
	"bb.u-root.com/bb/pkg/bbmain"
	t "github.com/f-secure-foundry/tamago-example"
	r "github.com/u-root/u-root/cmds/exp/rush"
)

func init() {
	bbmain.ListCmds()
	for _, c := range bbmain.GetBBCCmds(){
		r.ExternalBuiltIn(c.Name, c.F)
	}
	t.Runner = run
}
