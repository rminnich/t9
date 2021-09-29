// https://en.wikipedia.org/wiki/Media-independent_interface
// The standard MII features a small set of registers:[1]

// Basic Mode Configuration (#0)
// Status Word (#1)
// PHY Identification (#2, #3)
// Ability Advertisement (#4)
// Link Partner Ability (#5)
// Auto Negotiation Expansion (#6)
// The MII Status Word is the most useful datum, since it may be used to detect whether an Ethernet NIC is connected to a network. It contains a bit field with the following information:[2]

// Bit value	Meaning
// 0x8000	Capable of 100BASE-T4
// 0x7800	Capable of 10/100 HD/FD (most common)
// 0x0040	Preamble suppression permitted
// 0x0020	Autonegotiation complete
// 0x0010	Remote fault
// 0x0008	Capable of Autonegotiation
// 0x0004	Link established
// 0x0002	Jabber detected
// 0x0001	Extended MII registers exist
package main

const (
	T4                           = 0x8000
	HDFD                         = 0x7800
	PreambleSuppressionPermitted = 0x40
	AutoNegotiationComplete      = 0x20
	RemoteFault                  = 0x10
	AutoNegotiationCapable       = 0x8
	Link                         = 4
	Jabber                       = 2
	ExtendedMII                  = 1
)

type (
	Status uint16
	CSR    uint16
)

type MII interface {
	Name() string
	Read(addr, devad, reg uint32) (uint32, error)
	Write(addr, devad, reg uint32, val uint16) error
	Reset() error
	// later: add phy
}

//         int (*reset)(struct mii_dev *bus);
//         struct phy_device *phymap[PHY_MAX_ADDR];
// 	u32 phy_mask;
// }

type ULLMII struct {
	Base uint32
}

func (s Status) String() string {
	var out string
	if s&T4 == T4 {
		out += "T4|"
	} else {
		out += "~T4|"
	}
	if s&HDFD == HDFD {
		out += "HDFD|"
	} else {
		out += "~HDFD|"
	}
	if s&PreambleSuppressionPermitted == PreambleSuppressionPermitted {
		out += "PreambleSuppressionPermitted|"
	} else {
		out += "~PreambleSuppressionPermitted|"
	}
	if s&AutoNegotiationComplete == AutoNegotiationComplete {
		out += "AutoNegotiationComplete|"
	} else {
		out += "~AutoNegotiationComplete|"
	}
	if s&RemoteFault == RemoteFault {
		out += "RemoteFault|"
	} else {
		out += "~RemoteFault|"
	}
	if s&AutoNegotiationCapable == AutoNegotiationCapable {
		out += "AutoNegotiationCapable|"
	} else {
		out += "~AutoNegotiationCapable|"
	}
	if s&Link == Link {
		out += "Link|"
	} else {
		out += "~Link|"
	}
	if s&Jabber == Jabber {
		out += "Jabber|"
	} else {
		out += "~Jabber|"
	}
	if s&ExtendedMII == ExtendedMII {
		out += "ExtendedMII"
	} else {
		out += "~ExtendedMII"
	}
	return out
}

const ()

func (u *ULLMII) Name() string {
	return "IMX6"
}

func NewMII() (MII, error) {
	return &ULLMII{}, nil
}
