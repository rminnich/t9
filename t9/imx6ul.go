// NXP i.MX6UL configuration and support
// https://github.com/usbarmory/tamago
//
// Copyright (c) WithSecure Corporation
// https://foundry.withsecure.com
//
// Use of this source code is governed by the license
// that can be found in the LICENSE file.

// Package imx6ul provides support to Go bare metal unikernels, written using
// the TamaGo framework, on the NXP i.MX6UL family of System-on-Chip (SoC)
// application processors.
//
// The package implements initialization and drivers for NXP
// i.MX6UL/i.MX6ULL/i.MX6ULZ SoCs, adopting the following reference
// specifications:
//   - IMX6ULCEC  - i.MX6UL  Data Sheet                               - Rev 2.2 2015/05
//   - IMX6ULLCEC - i.MX6ULL Data Sheet                               - Rev 1.2 2017/11
//   - IMX6ULZCEC - i.MX6ULZ Data Sheet                               - Rev 0   2018/09
//   - IMX6ULRM   - i.MX 6UL  Applications Processor Reference Manual - Rev 1   2016/04
//   - IMX6ULLRM  - i.MX 6ULL Applications Processor Reference Manual - Rev 1   2017/11
//   - IMX6ULZRM  - i.MX 6ULZ Applications Processor Reference Manual - Rev 0   2018/10
//
// This package is only meant to be used with `GOOS=tamago GOARCH=arm` as
// supported by the TamaGo framework for bare metal Go on ARM SoCs, see
// https://github.com/usbarmory/tamago.
package t9

// Peripheral registers
const (
	// Central Security Unit
	CSU_BASE = 0x021c0000

	// Data Co-Processor (ULL/ULZ only)
	DCP_BASE = 0x02280000

	// General Interrupt Controller
	GIC_BASE = 0x00a00000

	// General Purpose I/O
	GPIO1_BASE = 0x0209c000
	GPIO2_BASE = 0x020a0000
	GPIO3_BASE = 0x020a4000
	GPIO4_BASE = 0x020a8000
	GPIO5_BASE = 0x020ac000

	// Ethernet MAC (UL/ULL only)
	ENET1_BASE = 0x02188000
	ENET2_BASE = 0x020b4000

	// I2C
	I2C1_BASE = 0x021a0000
	I2C2_BASE = 0x021a4000

	// On-Chip OTP Controller
	OCOTP_BASE      = 0x021bc000
	OCOTP_BANK_BASE = 0x021bc400

	// On-Chip Random-Access Memory
	OCRAM_START = 0x00900000
	OCRAM_SIZE  = 0x20000

	// True Random Number Generator (ULL/ULZ only)
	RNGB_BASE = 0x02284000

	// Secure Non-Volatile Storage
	SNVS_BASE = 0x020cc000

	// TrustZone Address Space Controller
	TZASC_BASE            = 0x021d0000
	TZASC_BYPASS          = 0x020e4024
	GPR1_TZASC1_BOOT_LOCK = 23

	// Serial ports
	UART1_BASE = 0x02020000
	UART2_BASE = 0x021e8000
	UART3_BASE = 0x021ec000
	UART4_BASE = 0x021f0000

	// USB 2.0
	USB_ANALOG1_BASE   = 0x020c81a0
	USB_ANALOG2_BASE   = 0x020c8200
	USB_ANALOG_DIGPROG = 0x020c8260
	USBPHY1_BASE       = 0x020c9000
	USBPHY2_BASE       = 0x020ca000
	USB1_BASE          = 0x02184000
	USB2_BASE          = 0x02184200

	// SD/MMC
	USDHC1_BASE = 0x02190000
	USDHC2_BASE = 0x02194000
)

// Peripheral instances
var (
	// GPIO controller 4
	GPIO4 = &GPIO{
		Index: 4,
		Base:  GPIO4_BASE,
		CCGR:  CCM_CCGR3,
		CG:    CCGRx_CG6,
	}
)

// SiliconVersion returns the SoC silicon version information
// (p3945, 57.4.11 Chip Silicon Version (USB_ANALOG_DIGPROG), IMX6ULLRM).
func SiliconVersion(reg reg) (sv, family, revMajor, revMinor uint32) {
	sv = reg.Read(USB_ANALOG_DIGPROG)

	family = (sv >> 16) & 0xff
	revMajor = (sv >> 8) & 0xff
	revMinor = sv & 0xff

	return
}
