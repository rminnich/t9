package bbmain

import "log"

// For circular dependency reasons, we need to duplicate this type.
type BBC struct {
	Name string
	F    func() error
}

func GetBBCCmds() []BBC {
	var b []BBC
	for n, c := range bbCmds {
		if n == "bb" {
			continue
		}
		vv := func(n string, bbc bbCmd) func() error {
			return func() error {
				log.Printf("GO FUCK YOURSELF@ %v %v %v", n, bbc.init, bbc.main)
				bbc.init()
				bbc.main()
				return nil
			}
		}(n, c)
		log.Printf("ADD NAME %s init %v %v func %v", n, c.init, c.main, vv)
		b = append(b, BBC{Name: n, F: vv})
	}
	return b
}
