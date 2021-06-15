package main

import (
	"log"
	"os"

	"bb.u-root.com/bb/pkg/bbmain"
	t "github.com/f-secure-foundry/tamago-example"
	r "github.com/u-root/u-root/cmds/exp/rush"
)

func addBBBuiltInsandRun() {
		for _, c := range bbmain.GetBBCCmds() {
			if err := r.ExternalBuiltIn(c.Name, c.F); err != nil {
				log.Printf("Add %s: %v %v", c.Name, c.F, err)
			}
			log.Printf("ADDED %s", c.Name)
		}
	log.Printf("LET's GO RUN %q", os.Args)
	run()
	log.Printf("RUN RETURNED!!! os.Args is now %q. WTF?", os.Args)
}

func init() {
	log.Printf("init for adding shit!!!")
	bbmain.ListCmds()
	t.Runner = addBBBuiltInsandRun
}
