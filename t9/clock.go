// NXP i.MX6UL clock control
// https://github.com/usbarmory/tamago
//
// Copyright (c) WithSecure Corporation
// https://foundry.withsecure.com
//
// Use of this source code is governed by the license
// that can be found in the LICENSE file.

package t9

// Clock registers
const (
	CCM_CACRR      = 0x020c4010
	CACRR_ARM_PODF = 0

	CCM_CBCDR      = 0x020c4014
	CBCDR_IPG_PODF = 8

	CCM_CSCDR1           = 0x020c4024
	CSCDR1_USDHC2_PODF   = 16
	CSCDR1_USDHC1_PODF   = 11
	CSCDR1_UART_CLK_SEL  = 6
	CSCDR1_UART_CLK_PODF = 0

	CCM_CSCMR1            = 0x020c401c
	CSCMR1_USDHC2_CLK_SEL = 17
	CSCMR1_USDHC1_CLK_SEL = 16
	CSCMR1_PERCLK_SEL     = 6
	CSCMR1_PERCLK_PODF    = 0

	CCM_ANALOG_PLL_ARM = 0x020c8000
	PLL_LOCK           = 31
	PLL_BYPASS         = 16
	PLL_BYPASS_CLK_SRC = 14
	PLL_ENABLE         = 13
	PLL_POWER          = 12
	PLL_DIV_SELECT     = 0

	CCM_ANALOG_PLL_USB1 = CCM_ANALOG_PLL_ARM + 0x10
	CCM_ANALOG_PLL_USB2 = CCM_ANALOG_PLL_ARM + 0x20
	PLL_EN_USB_CLKS     = 6

	CCM_ANALOG_PLL_ENET  = CCM_ANALOG_PLL_ARM + 0xe0
	PLL_ENET2_125M_EN    = 20
	PLL_ENET1_125M_EN    = 13
	PLL_ENET1_DIV_SELECT = 2
	PLL_ENET0_DIV_SELECT = 0

	CCM_ANALOG_PFD_480  = 0x020c80f0
	CCM_ANALOG_PFD_528  = 0x020c8100
	ANALOG_PFD3_CLKGATE = 31
	ANALOG_PFD3_FRAC    = 24
	ANALOG_PFD2_CLKGATE = 23
	ANALOG_PFD2_FRAC    = 16
	ANALOG_PFD1_CLKGATE = 15
	ANALOG_PFD1_FRAC    = 8
	ANALOG_PFD0_CLKGATE = 7
	ANALOG_PFD0_FRAC    = 0

	PMU_REG_CORE   = 0x020c8140
	CORE_REG2_TARG = 18
	CORE_REG0_TARG = 0

	CCM_CCGR0 = 0x020c4068
	CCM_CCGR1 = 0x020c406c
	CCM_CCGR2 = 0x020c4070
	CCM_CCGR3 = 0x020c4074
	CCM_CCGR5 = 0x020c407c
	CCM_CCGR6 = 0x020c4080

	CCGRx_CG15 = 30
	CCGRx_CG14 = 28
	CCGRx_CG13 = 26
	CCGRx_CG12 = 24
	CCGRx_CG11 = 22
	CCGRx_CG10 = 20
	CCGRx_CG9  = 18
	CCGRx_CG8  = 16
	CCGRx_CG7  = 14
	CCGRx_CG6  = 12
	CCGRx_CG5  = 10
	CCGRx_CG4  = 8
	CCGRx_CG3  = 6
	CCGRx_CG2  = 4
	CCGRx_CG1  = 2
	CCGRx_CG0  = 0
)

const (
	IOMUXC_GPR_GPR1  = 0x020e4004
	ENET2_TX_CLK_DIR = 18
	ENET1_TX_CLK_DIR = 17
	ENET2_CLK_SEL    = 14
	ENET1_CLK_SEL    = 13
)

// Oscillator frequencies
const (
	OSC_FREQ  = 24000000
	PLL2_FREQ = 528000000
	PLL3_FREQ = 480000000
)

// Operating ARM core frequencies in MHz
// (p24, Table 10. Operating Ranges, IMX6ULLCEC).
const (
	FreqMax = Freq900
	Freq900 = 900
	Freq792 = 792
	Freq528 = 528
	Freq396 = 396
	Freq198 = 198
	FreqLow = Freq198
)

// Clocks at boot time
// (p261, Table 8-4. Normal frequency clocks configuration, IMX6ULLRM)
const (
	IPG_FREQ = 66000000
	AHB_FREQ = 132000000
)
