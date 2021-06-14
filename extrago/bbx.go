package bbmain

// For circular dependency reasons, we need to duplicate this type.
type BBC struct {
	Name string
	F    func()
}

func GetBBCCmds() []BBC {
	var b []BBC
	for n, c := range bbCmds {
		if n == "bb" {
			continue
		}
		b = append(b, BBC{Name: n, F: c.main})
	}
	return b
}
