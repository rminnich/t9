package main

import (
	"fmt"
)

// the u-boot cpio code uses Linux code for some reason.
// It adds more heat than light, by far. We don't need to bitstuff!

const (
	Data    = 0
	Dir     = 4
	PadSR   = 8
	ICR1    = 12
	ICR2    = 16
	IMR     = 20
	ISR     = 24
	EdgeSel = 28
)

const (
	g1 = 0x209_c000
	g2 = 0x20a_0000
	g3 = 0x20a_4000
	g4 = 0x20a_8000
	g5 = 0x20a_c000
)

var GDI = fmt.Printf

var gports = map[int]string{
	Data:    "Data",
	Dir:     "Dir",
	PadSR:   "PadSR",
	ICR1:    "ICR1",
	ICR2:    "ICR2",
	IMR:     "IMR",
	ISR:     "ISR",
	EdgeSel: "EdgeSel",
}

// These are int64 so they work as offsets.
var gbase = []struct {
	base int64
	name string
}{
	{g1, "GPIO1"},
	{g2, "GPIO2"},
	{g3, "GPIO3"},
	{g4, "GPIO4"},
	{g5, "GPIO5"},
}

type GPIO struct {
	port     uint
	mask     uint32
	base     int64
	pin      uint
	name     string
	portName string
	err      error // THis is an accumulated error.
	rw       rw
}

// NewGPIO returns a GPIO for a given port and pin.
// Further use of the GPIO should only ever result in IO errors, but not
// specification (bad port, bad pin) errors.
func NewGPIO(rw rw, port, pin uint, name string) *GPIO {
	g := &GPIO{rw: rw, port: port, pin: pin, name: name}
	if port > uint(len(gbase)-1) {
		g.err = fmt.Errorf("port %d is invalid, has to be in range 0..%d", pin, len(gbase)-1)
		return g
	}
	b := gbase[port]
	g.base = b.base
	g.name = name
	g.portName = b.name
	if pin > 31 {
		g.err = fmt.Errorf("Bad pin %d has to be in range 0..31", pin)
		return g
	}
	g.pin = pin
	g.mask = 1 << pin
	return g
}

const (
	in = 0
	out
	trifloat
)

func (g *GPIO) regRead(reg uint) uint32 {
	if g.err != nil {
		return 0
	}
	v, err := readl(g.rw, uint32(g.base)+uint32(reg))
	GDI("regRead %#x %#x reg %v val %#x err %v", g.portName, g.pin, gports[int(reg)], v, g.err)
	g.err = err
	return v
}

func (g *GPIO) regWrite(reg uint, val uint32) {
	if g.err != nil {
		return
	}
	g.err = writel(g.rw, val, uint32(g.base)+uint32(reg))
	GDI("regWrite %#x %#x reg %v val %#x err %v", g.portName, g.pin, gports[int(reg)], val, g.err)
}

func (g *GPIO) direction(out bool) *GPIO {
	if g.err != nil {
		return g
	}
	dr := g.regRead(Dir)
	if g.err != nil {
		return g
	}
	if out {
		dr |= g.mask
	} else {
		dr &= ^g.mask
	}
	g.regWrite(Dir, dr)
	return g
}

// Set sets the value.
func (g *GPIO) Set(val int) *GPIO {
	if g.err != nil {
		return g
	}
	g.regWrite(Data, uint32(val))
	return g
}

// Output sets the output direction.
func (g *GPIO) Output() *GPIO {
	if g.err != nil {
		return g
	}
	g.regWrite(Dir, 1)
	return g
}

// Input sets the GPIO to be an input
func (g *GPIO) Input() *GPIO {
	if g.err != nil {
		return g
	}
	g.regWrite(Dir, 0)
	return g
}

// Delay delays using the functions passed in.
// They might wait for a fixed time, wait for a bit with a timeout,
// and so on.
func (g *GPIO) Delay(f ...func() error) *GPIO {
	if g.err != nil {
		return g
	}
	for _, df := range f {
		g.err = df()
		if g.err != nil {
			break
		}
	}
	return g
}
