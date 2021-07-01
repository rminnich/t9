package main

const (
	// this is a readonly field locked to med. define PAD_CTL_SPEED_LOW       (1 << 6)

	// invalid value, sentinal, remove.
	NO_PAD_CTRL = 1 << 17

	PAD_CTL_SPEED_MED  = (2 << 6)
	PAD_CTL_SPEED_HIGH = (3 << 6)

	PAD_CTL_DSE_DISABLE = (0 << 3)
	PAD_CTL_DSE_240ohm  = (1 << 3)
	PAD_CTL_DSE_120ohm  = (2 << 3)
	PAD_CTL_DSE_80ohm   = (3 << 3)
	PAD_CTL_DSE_60ohm   = (4 << 3)
	PAD_CTL_DSE_48ohm   = (5 << 3)
	PAD_CTL_DSE_40ohm   = (6 << 3)
	PAD_CTL_DSE_34ohm   = (7 << 3)

	PAD_CTL_DSE_260ohm = (1 << 3)
	PAD_CTL_DSE_130ohm = (2 << 3)
	PAD_CTL_DSE_88ohm  = (3 << 3)
	PAD_CTL_DSE_65ohm  = (4 << 3)
	PAD_CTL_DSE_52ohm  = (5 << 3)
	PAD_CTL_DSE_43ohm  = (6 << 3)
	PAD_CTL_DSE_37ohm  = (7 << 3)

	IOMUX_CONFIG_SION = 0x10

	PAD_CTL_HYS = (1 << 16)

	PAD_CTL_PUS_100K_DOWN = (0<<14 | PAD_CTL_PUE)
	PAD_CTL_PUS_47K_UP    = (1<<14 | PAD_CTL_PUE)
	PAD_CTL_PUS_100K_UP   = (2<<14 | PAD_CTL_PUE)
	PAD_CTL_PUS_22K_UP    = (3<<14 | PAD_CTL_PUE)
	PAD_CTL_PUE           = (1<<13 | PAD_CTL_PKE)
	PAD_CTL_PKE           = (1 << 12)

	PAD_CTL_ODE = (1 << 11)
)

type pad struct {
	padCtlOFS   uint32
	padCtl      uint32
	muxCtlOFS   uint32
	muxMode     uint32
	selInputOFS uint32
	selInput    uint32
}

var (
	MX6_PAD_BOOT_MODE0__GPIO5_IO10 = &pad{padCtlOFS: 0x02A0, padCtl: 0, muxCtlOFS: 0x0014, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_BOOT_MODE1__GPIO5_IO11 = &pad{padCtlOFS: 0x02A4, padCtl: 0, muxCtlOFS: 0x0018, muxMode: 5, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_SNVS_TAMPER0__GPIO5_IO00 = &pad{padCtlOFS: 0x02A8, padCtl: 0, muxCtlOFS: 0x001C, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_SNVS_TAMPER1__GPIO5_IO01 = &pad{padCtlOFS: 0x02AC, padCtl: 0, muxCtlOFS: 0x0020, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_SNVS_TAMPER2__GPIO5_IO02 = &pad{padCtlOFS: 0x02B0, padCtl: 0, muxCtlOFS: 0x0024, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_SNVS_TAMPER3__GPIO5_IO03 = &pad{padCtlOFS: 0x02B4, padCtl: 0, muxCtlOFS: 0x0028, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_SNVS_TAMPER4__GPIO5_IO04 = &pad{padCtlOFS: 0x02B8, padCtl: 0, muxCtlOFS: 0x002C, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_SNVS_TAMPER5__GPIO5_IO05 = &pad{padCtlOFS: 0x02BC, padCtl: 0, muxCtlOFS: 0x0030, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_SNVS_TAMPER6__GPIO5_IO06 = &pad{padCtlOFS: 0x02C0, padCtl: 0, muxCtlOFS: 0x0034, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_SNVS_TAMPER7__GPIO5_IO07 = &pad{padCtlOFS: 0x02C4, padCtl: 0, muxCtlOFS: 0x0038, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_SNVS_TAMPER8__GPIO5_IO08 = &pad{padCtlOFS: 0x02C8, padCtl: 0, muxCtlOFS: 0x003C, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_SNVS_TAMPER9__GPIO5_IO09 = &pad{padCtlOFS: 0x02CC, padCtl: 0, muxCtlOFS: 0x0040, muxMode: 5, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_JTAG_MOD__SJC_MOD           = &pad{padCtlOFS: 0x02D0, padCtl: 0, muxCtlOFS: 0x0044, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_JTAG_MOD__GPT2_CLK          = &pad{padCtlOFS: 0x02D0, padCtl: 0, muxCtlOFS: 0x0044, muxMode: 1, selInputOFS: 0x05A0, selInput: 0}
	MX6_PAD_JTAG_MOD__SPDIF_OUT         = &pad{padCtlOFS: 0x02D0, padCtl: 0, muxCtlOFS: 0x0044, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_JTAG_MOD__ENET1_REF_CLK_25M = &pad{padCtlOFS: 0x02D0, padCtl: 0, muxCtlOFS: 0x0044, muxMode: 3, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_JTAG_MOD__CCM_PMIC_RDY      = &pad{padCtlOFS: 0x02D0, padCtl: 0, muxCtlOFS: 0x0044, muxMode: 4, selInputOFS: 0x04C0, selInput: 0}
	MX6_PAD_JTAG_MOD__GPIO1_IO10        = &pad{padCtlOFS: 0x02D0, padCtl: 0, muxCtlOFS: 0x0044, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_JTAG_MOD__SDMA_EXT_EVENT00  = &pad{padCtlOFS: 0x02D0, padCtl: 0, muxCtlOFS: 0x0044, muxMode: 6, selInputOFS: 0x0610, selInput: 0}

	MX6_PAD_JTAG_TMS__SJC_TMS          = &pad{padCtlOFS: 0x02D4, padCtl: 0, muxCtlOFS: 0x0048, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_JTAG_TMS__GPT2_CAPTURE1    = &pad{padCtlOFS: 0x02D4, padCtl: 0, muxCtlOFS: 0x0048, muxMode: 1, selInputOFS: 0x0598, selInput: 0}
	MX6_PAD_JTAG_TMS__SAI2_MCLK        = &pad{padCtlOFS: 0x02D4, padCtl: 0, muxCtlOFS: 0x0048, muxMode: 2, selInputOFS: 0x05F0, selInput: 0}
	MX6_PAD_JTAG_TMS__CCM_CLKO1        = &pad{padCtlOFS: 0x02D4, padCtl: 0, muxCtlOFS: 0x0048, muxMode: 3, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_JTAG_TMS__CCM_WAIT         = &pad{padCtlOFS: 0x02D4, padCtl: 0, muxCtlOFS: 0x0048, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_JTAG_TMS__GPIO1_IO11       = &pad{padCtlOFS: 0x02D4, padCtl: 0, muxCtlOFS: 0x0048, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_JTAG_TMS__SDMA_EXT_EVENT01 = &pad{padCtlOFS: 0x02D4, padCtl: 0, muxCtlOFS: 0x0048, muxMode: 6, selInputOFS: 0x0614, selInput: 0}
	MX6_PAD_JTAG_TMS__EPIT1_OUT        = &pad{padCtlOFS: 0x02D4, padCtl: 0, muxCtlOFS: 0x0048, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_JTAG_TDO__SJC_TDO       = &pad{padCtlOFS: 0x02D8, padCtl: 0, muxCtlOFS: 0x004C, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_JTAG_TDO__GPT2_CAPTURE2 = &pad{padCtlOFS: 0x02D8, padCtl: 0, muxCtlOFS: 0x004C, muxMode: 1, selInputOFS: 0x059C, selInput: 0}
	MX6_PAD_JTAG_TDO__SAI2_TX_SYNC  = &pad{padCtlOFS: 0x02D8, padCtl: 0, muxCtlOFS: 0x004C, muxMode: 2, selInputOFS: 0x05FC, selInput: 0}
	MX6_PAD_JTAG_TDO__CCM_CLKO2     = &pad{padCtlOFS: 0x02D8, padCtl: 0, muxCtlOFS: 0x004C, muxMode: 3, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_JTAG_TDO__CCM_STOP      = &pad{padCtlOFS: 0x02D8, padCtl: 0, muxCtlOFS: 0x004C, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_JTAG_TDO__GPIO1_IO12    = &pad{padCtlOFS: 0x02D8, padCtl: 0, muxCtlOFS: 0x004C, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_JTAG_TDO__MQS_RIGHT     = &pad{padCtlOFS: 0x02D8, padCtl: 0, muxCtlOFS: 0x004C, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_JTAG_TDO__EPIT2_OUT     = &pad{padCtlOFS: 0x02D8, padCtl: 0, muxCtlOFS: 0x004C, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_JTAG_TDI__SJC_TDI         = &pad{padCtlOFS: 0x02DC, padCtl: 0, muxCtlOFS: 0x0050, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_JTAG_TDI__GPT2_COMPARE1   = &pad{padCtlOFS: 0x02DC, padCtl: 0, muxCtlOFS: 0x0050, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_JTAG_TDI__SAI2_TX_BCLK    = &pad{padCtlOFS: 0x02DC, padCtl: 0, muxCtlOFS: 0x0050, muxMode: 2, selInputOFS: 0x05F8, selInput: 0}
	MX6_PAD_JTAG_TDI__PWM6_OUT        = &pad{padCtlOFS: 0x02DC, padCtl: 0, muxCtlOFS: 0x0050, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_JTAG_TDI__GPIO1_IO13      = &pad{padCtlOFS: 0x02DC, padCtl: 0, muxCtlOFS: 0x0050, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_JTAG_TDI__MQS_LEFT        = &pad{padCtlOFS: 0x02DC, padCtl: 0, muxCtlOFS: 0x0050, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_JTAG_TDI__SIM1_POWER_FAIL = &pad{padCtlOFS: 0x02DC, padCtl: 0, muxCtlOFS: 0x0050, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_JTAG_TCK__SJC_TCK         = &pad{padCtlOFS: 0x02E0, padCtl: 0, muxCtlOFS: 0x0054, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_JTAG_TCK__GPT2_COMPARE2   = &pad{padCtlOFS: 0x02E0, padCtl: 0, muxCtlOFS: 0x0054, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_JTAG_TCK__SAI2_RX_DATA    = &pad{padCtlOFS: 0x02E0, padCtl: 0, muxCtlOFS: 0x0054, muxMode: 2, selInputOFS: 0x05F4, selInput: 0}
	MX6_PAD_JTAG_TCK__PWM7_OUT        = &pad{padCtlOFS: 0x02E0, padCtl: 0, muxCtlOFS: 0x0054, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_JTAG_TCK__GPIO1_IO14      = &pad{padCtlOFS: 0x02E0, padCtl: 0, muxCtlOFS: 0x0054, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_JTAG_TCK__SIM2_POWER_FAIL = &pad{padCtlOFS: 0x02E0, padCtl: 0, muxCtlOFS: 0x0054, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_JTAG_TRST_B__SJC_TRSTB        = &pad{padCtlOFS: 0x02E4, padCtl: 0, muxCtlOFS: 0x0058, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_JTAG_TRST_B__GPT2_COMPARE3    = &pad{padCtlOFS: 0x02E4, padCtl: 0, muxCtlOFS: 0x0058, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_JTAG_TRST_B__SAI2_TX_DATA     = &pad{padCtlOFS: 0x02E4, padCtl: 0, muxCtlOFS: 0x0058, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_JTAG_TRST_B__PWM8_OUT         = &pad{padCtlOFS: 0x02E4, padCtl: 0, muxCtlOFS: 0x0058, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_JTAG_TRST_B__GPIO1_IO15       = &pad{padCtlOFS: 0x02E4, padCtl: 0, muxCtlOFS: 0x0058, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_JTAG_TRST_B__CAAM_RNG_OSC_OBS = &pad{padCtlOFS: 0x02E4, padCtl: 0, muxCtlOFS: 0x0058, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_GPIO1_IO00__I2C2_SCL             = &pad{padCtlOFS: 0x02E8, padCtl: 0, muxCtlOFS: 0x005C, muxMode: IOMUX_CONFIG_SION | 0, selInputOFS: 0x05AC, selInput: 1}
	MX6_PAD_GPIO1_IO00__GPT1_CAPTURE1        = &pad{padCtlOFS: 0x02E8, padCtl: 0, muxCtlOFS: 0x005C, muxMode: 1, selInputOFS: 0x058C, selInput: 0}
	MX6_PAD_GPIO1_IO00__ANATOP_OTG1_ID       = &pad{padCtlOFS: 0x02E8, padCtl: 0, muxCtlOFS: 0x005C, muxMode: 2, selInputOFS: 0x04B8, selInput: 0}
	MX6_PAD_GPIO1_IO00__ENET1_REF_CLK1       = &pad{padCtlOFS: 0x02E8, padCtl: 0, muxCtlOFS: 0x005C, muxMode: 3, selInputOFS: 0x0574, selInput: 0}
	MX6_PAD_GPIO1_IO00__MQS_RIGHT            = &pad{padCtlOFS: 0x02E8, padCtl: 0, muxCtlOFS: 0x005C, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_GPIO1_IO00__GPIO1_IO00           = &pad{padCtlOFS: 0x02E8, padCtl: 0, muxCtlOFS: 0x005C, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_GPIO1_IO00__ENET1_1588_EVENT0_IN = &pad{padCtlOFS: 0x02E8, padCtl: 0, muxCtlOFS: 0x005C, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_GPIO1_IO00__SRC_SYSTEM_RESET     = &pad{padCtlOFS: 0x02E8, padCtl: 0, muxCtlOFS: 0x005C, muxMode: 7, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_GPIO1_IO00__WDOG3_WDOG_B         = &pad{padCtlOFS: 0x02E8, padCtl: 0, muxCtlOFS: 0x005C, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_GPIO1_IO01__I2C2_SDA              = &pad{padCtlOFS: 0x02EC, padCtl: 0, muxCtlOFS: 0x0060, muxMode: IOMUX_CONFIG_SION | 0, selInputOFS: 0x05B0, selInput: 1}
	MX6_PAD_GPIO1_IO01__GPT1_COMPARE1         = &pad{padCtlOFS: 0x02EC, padCtl: 0, muxCtlOFS: 0x0060, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_GPIO1_IO01__USB_OTG1_OC           = &pad{padCtlOFS: 0x02EC, padCtl: 0, muxCtlOFS: 0x0060, muxMode: 2, selInputOFS: 0x0664, selInput: 0}
	MX6_PAD_GPIO1_IO01__ENET2_REF_CLK2        = &pad{padCtlOFS: 0x02EC, padCtl: 0, muxCtlOFS: 0x0060, muxMode: 3, selInputOFS: 0x057C, selInput: 0}
	MX6_PAD_GPIO1_IO01__MQS_LEFT              = &pad{padCtlOFS: 0x02EC, padCtl: 0, muxCtlOFS: 0x0060, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_GPIO1_IO01__GPIO1_IO01            = &pad{padCtlOFS: 0x02EC, padCtl: 0, muxCtlOFS: 0x0060, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_GPIO1_IO01__ENET1_1588_EVENT0_OUT = &pad{padCtlOFS: 0x02EC, padCtl: 0, muxCtlOFS: 0x0060, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_GPIO1_IO01__SRC_EARLY_RESET       = &pad{padCtlOFS: 0x02EC, padCtl: 0, muxCtlOFS: 0x0060, muxMode: 7, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_GPIO1_IO01__WDOG1_WDOG_B          = &pad{padCtlOFS: 0x02EC, padCtl: 0, muxCtlOFS: 0x0060, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_GPIO1_IO02__I2C1_SCL          = &pad{padCtlOFS: 0x02F0, padCtl: 0, muxCtlOFS: 0x0064, muxMode: IOMUX_CONFIG_SION | 0, selInputOFS: 0x05A4, selInput: 0}
	MX6_PAD_GPIO1_IO02__GPT1_COMPARE2     = &pad{padCtlOFS: 0x02F0, padCtl: 0, muxCtlOFS: 0x0064, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_GPIO1_IO02__USB_OTG2_PWR      = &pad{padCtlOFS: 0x02F0, padCtl: 0, muxCtlOFS: 0x0064, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_GPIO1_IO02__ENET1_REF_CLK_25M = &pad{padCtlOFS: 0x02F0, padCtl: 0, muxCtlOFS: 0x0064, muxMode: 3, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_GPIO1_IO02__USDHC1_WP         = &pad{padCtlOFS: 0x02F0, padCtl: 0, muxCtlOFS: 0x0064, muxMode: 4, selInputOFS: 0x066C, selInput: 0}
	MX6_PAD_GPIO1_IO02__GPIO1_IO02        = &pad{padCtlOFS: 0x02F0, padCtl: 0, muxCtlOFS: 0x0064, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_GPIO1_IO02__SDMA_EXT_EVENT00  = &pad{padCtlOFS: 0x02F0, padCtl: 0, muxCtlOFS: 0x0064, muxMode: 6, selInputOFS: 0x0610, selInput: 1}
	MX6_PAD_GPIO1_IO02__SRC_ANY_PU_RESET  = &pad{padCtlOFS: 0x02F0, padCtl: 0, muxCtlOFS: 0x0064, muxMode: 7, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_GPIO1_IO02__UART1_DCE_TX      = &pad{padCtlOFS: 0x02F0, padCtl: 0, muxCtlOFS: 0x0064, muxMode: 8, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_GPIO1_IO02__UART1_DTE_RX      = &pad{padCtlOFS: 0x02F0, padCtl: 0, muxCtlOFS: 0x0064, muxMode: 8, selInputOFS: 0x0624, selInput: 0}

	MX6_PAD_GPIO1_IO03__I2C1_SDA        = &pad{padCtlOFS: 0x02F4, padCtl: 0, muxCtlOFS: 0x0068, muxMode: IOMUX_CONFIG_SION | 0, selInputOFS: 0x05A8, selInput: 1}
	MX6_PAD_GPIO1_IO03__GPT1_COMPARE3   = &pad{padCtlOFS: 0x02F4, padCtl: 0, muxCtlOFS: 0x0068, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_GPIO1_IO03__USB_OTG2_OC     = &pad{padCtlOFS: 0x02F4, padCtl: 0, muxCtlOFS: 0x0068, muxMode: 2, selInputOFS: 0x0660, selInput: 0}
	MX6_PAD_GPIO1_IO03__USDHC1_CD_B     = &pad{padCtlOFS: 0x02F4, padCtl: 0, muxCtlOFS: 0x0068, muxMode: 4, selInputOFS: 0x0668, selInput: 0}
	MX6_PAD_GPIO1_IO03__GPIO1_IO03      = &pad{padCtlOFS: 0x02F4, padCtl: 0, muxCtlOFS: 0x0068, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_GPIO1_IO03__CCM_DI0_EXT_CLK = &pad{padCtlOFS: 0x02F4, padCtl: 0, muxCtlOFS: 0x0068, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_GPIO1_IO03__SRC_TESTER_ACK  = &pad{padCtlOFS: 0x02F4, padCtl: 0, muxCtlOFS: 0x0068, muxMode: 7, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_GPIO1_IO03__UART1_DCE_RX    = &pad{padCtlOFS: 0x02F4, padCtl: 0, muxCtlOFS: 0x0068, muxMode: 8, selInputOFS: 0x0624, selInput: 1}
	MX6_PAD_GPIO1_IO03__UART1_DTE_TX    = &pad{padCtlOFS: 0x02F4, padCtl: 0, muxCtlOFS: 0x0068, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_GPIO1_IO04__ENET1_REF_CLK1       = &pad{padCtlOFS: 0x02F8, padCtl: 0, muxCtlOFS: 0x006C, muxMode: 0, selInputOFS: 0x0574, selInput: 1}
	MX6_PAD_GPIO1_IO04__PWM3_OUT             = &pad{padCtlOFS: 0x02F8, padCtl: 0, muxCtlOFS: 0x006C, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_GPIO1_IO04__USB_OTG1_PWR         = &pad{padCtlOFS: 0x02F8, padCtl: 0, muxCtlOFS: 0x006C, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_GPIO1_IO04__USDHC1_RESET_B       = &pad{padCtlOFS: 0x02F8, padCtl: 0, muxCtlOFS: 0x006C, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_GPIO1_IO04__GPIO1_IO04           = &pad{padCtlOFS: 0x02F8, padCtl: 0, muxCtlOFS: 0x006C, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_GPIO1_IO04__ENET2_1588_EVENT0_IN = &pad{padCtlOFS: 0x02F8, padCtl: 0, muxCtlOFS: 0x006C, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_GPIO1_IO04__UART5_DCE_TX         = &pad{padCtlOFS: 0x02F8, padCtl: 0, muxCtlOFS: 0x006C, muxMode: 8, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_GPIO1_IO04__UART5_DTE_RX         = &pad{padCtlOFS: 0x02F8, padCtl: 0, muxCtlOFS: 0x006C, muxMode: 8, selInputOFS: 0x0644, selInput: 2}

	MX6_PAD_GPIO1_IO05__ENET2_REF_CLK2        = &pad{padCtlOFS: 0x02FC, padCtl: 0, muxCtlOFS: 0x0070, muxMode: 0, selInputOFS: 0x057C, selInput: 1}
	MX6_PAD_GPIO1_IO05__PWM4_OUT              = &pad{padCtlOFS: 0x02FC, padCtl: 0, muxCtlOFS: 0x0070, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_GPIO1_IO05__ANATOP_OTG2_ID        = &pad{padCtlOFS: 0x02FC, padCtl: 0, muxCtlOFS: 0x0070, muxMode: 2, selInputOFS: 0x04BC, selInput: 0}
	MX6_PAD_GPIO1_IO05__CSI_FIELD             = &pad{padCtlOFS: 0x02FC, padCtl: 0, muxCtlOFS: 0x0070, muxMode: 3, selInputOFS: 0x0530, selInput: 0}
	MX6_PAD_GPIO1_IO05__USDHC1_VSELECT        = &pad{padCtlOFS: 0x02FC, padCtl: 0, muxCtlOFS: 0x0070, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_GPIO1_IO05__GPIO1_IO05            = &pad{padCtlOFS: 0x02FC, padCtl: 0, muxCtlOFS: 0x0070, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_GPIO1_IO05__ENET2_1588_EVENT0_OUT = &pad{padCtlOFS: 0x02FC, padCtl: 0, muxCtlOFS: 0x0070, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_GPIO1_IO05__UART5_DCE_RX          = &pad{padCtlOFS: 0x02FC, padCtl: 0, muxCtlOFS: 0x0070, muxMode: 8, selInputOFS: 0x0644, selInput: 3}
	MX6_PAD_GPIO1_IO05__UART5_DTE_TX          = &pad{padCtlOFS: 0x02FC, padCtl: 0, muxCtlOFS: 0x0070, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_GPIO1_IO06__ENET1_MDIO       = &pad{padCtlOFS: 0x0300, padCtl: 0, muxCtlOFS: 0x0074, muxMode: 0, selInputOFS: 0x0578, selInput: 0}
	MX6_PAD_GPIO1_IO06__ENET2_MDIO       = &pad{padCtlOFS: 0x0300, padCtl: 0, muxCtlOFS: 0x0074, muxMode: 1, selInputOFS: 0x0580, selInput: 0}
	MX6_PAD_GPIO1_IO06__USB_OTG_PWR_WAKE = &pad{padCtlOFS: 0x0300, padCtl: 0, muxCtlOFS: 0x0074, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_GPIO1_IO06__CSI_MCLK         = &pad{padCtlOFS: 0x0300, padCtl: 0, muxCtlOFS: 0x0074, muxMode: 3, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_GPIO1_IO06__USDHC2_WP        = &pad{padCtlOFS: 0x0300, padCtl: 0, muxCtlOFS: 0x0074, muxMode: 4, selInputOFS: 0x069C, selInput: 0}
	MX6_PAD_GPIO1_IO06__GPIO1_IO06       = &pad{padCtlOFS: 0x0300, padCtl: 0, muxCtlOFS: 0x0074, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_GPIO1_IO06__CCM_WAIT         = &pad{padCtlOFS: 0x0300, padCtl: 0, muxCtlOFS: 0x0074, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_GPIO1_IO06__CCM_REF_EN_B     = &pad{padCtlOFS: 0x0300, padCtl: 0, muxCtlOFS: 0x0074, muxMode: 7, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_GPIO1_IO06__UART1_DCE_CTS    = &pad{padCtlOFS: 0x0300, padCtl: 0, muxCtlOFS: 0x0074, muxMode: 8, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_GPIO1_IO06__UART1_DTE_RTS    = &pad{padCtlOFS: 0x0300, padCtl: 0, muxCtlOFS: 0x0074, muxMode: 8, selInputOFS: 0x0620, selInput: 0}

	MX6_PAD_GPIO1_IO07__ENET1_MDC         = &pad{padCtlOFS: 0x0304, padCtl: 0, muxCtlOFS: 0x0078, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_GPIO1_IO07__ENET2_MDC         = &pad{padCtlOFS: 0x0304, padCtl: 0, muxCtlOFS: 0x0078, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_GPIO1_IO07__USB_OTG_HOST_MODE = &pad{padCtlOFS: 0x0304, padCtl: 0, muxCtlOFS: 0x0078, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_GPIO1_IO07__CSI_PIXCLK        = &pad{padCtlOFS: 0x0304, padCtl: 0, muxCtlOFS: 0x0078, muxMode: 3, selInputOFS: 0x0528, selInput: 0}
	MX6_PAD_GPIO1_IO07__USDHC2_CD_B       = &pad{padCtlOFS: 0x0304, padCtl: 0, muxCtlOFS: 0x0078, muxMode: 4, selInputOFS: 0x0674, selInput: 1}
	MX6_PAD_GPIO1_IO07__GPIO1_IO07        = &pad{padCtlOFS: 0x0304, padCtl: 0, muxCtlOFS: 0x0078, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_GPIO1_IO07__CCM_STOP          = &pad{padCtlOFS: 0x0304, padCtl: 0, muxCtlOFS: 0x0078, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_GPIO1_IO07__UART1_DCE_RTS     = &pad{padCtlOFS: 0x0304, padCtl: 0, muxCtlOFS: 0x0078, muxMode: 8, selInputOFS: 0x0620, selInput: 1}
	MX6_PAD_GPIO1_IO07__UART1_DTE_CTS     = &pad{padCtlOFS: 0x0304, padCtl: 0, muxCtlOFS: 0x0078, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_GPIO1_IO08__PWM1_OUT       = &pad{padCtlOFS: 0x0308, padCtl: 0, muxCtlOFS: 0x007C, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_GPIO1_IO08__WDOG1_WDOG_B   = &pad{padCtlOFS: 0x0308, padCtl: 0, muxCtlOFS: 0x007C, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_GPIO1_IO08__SPDIF_OUT      = &pad{padCtlOFS: 0x0308, padCtl: 0, muxCtlOFS: 0x007C, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_GPIO1_IO08__CSI_VSYNC      = &pad{padCtlOFS: 0x0308, padCtl: 0, muxCtlOFS: 0x007C, muxMode: 3, selInputOFS: 0x052C, selInput: 1}
	MX6_PAD_GPIO1_IO08__USDHC2_VSELECT = &pad{padCtlOFS: 0x0308, padCtl: 0, muxCtlOFS: 0x007C, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_GPIO1_IO08__GPIO1_IO08     = &pad{padCtlOFS: 0x0308, padCtl: 0, muxCtlOFS: 0x007C, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_GPIO1_IO08__CCM_PMIC_RDY   = &pad{padCtlOFS: 0x0308, padCtl: 0, muxCtlOFS: 0x007C, muxMode: 6, selInputOFS: 0x04C0, selInput: 1}
	MX6_PAD_GPIO1_IO08__UART5_DCE_RTS  = &pad{padCtlOFS: 0x0308, padCtl: 0, muxCtlOFS: 0x007C, muxMode: 8, selInputOFS: 0x0640, selInput: 1}
	MX6_PAD_GPIO1_IO08__UART5_DTE_CTS  = &pad{padCtlOFS: 0x0308, padCtl: 0, muxCtlOFS: 0x007C, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_GPIO1_IO09__PWM2_OUT       = &pad{padCtlOFS: 0x030C, padCtl: 0, muxCtlOFS: 0x0080, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_GPIO1_IO09__WDOG1_WDOG_ANY = &pad{padCtlOFS: 0x030C, padCtl: 0, muxCtlOFS: 0x0080, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_GPIO1_IO09__SPDIF_IN       = &pad{padCtlOFS: 0x030C, padCtl: 0, muxCtlOFS: 0x0080, muxMode: 2, selInputOFS: 0x0618, selInput: 0}
	MX6_PAD_GPIO1_IO09__CSI_HSYNC      = &pad{padCtlOFS: 0x030C, padCtl: 0, muxCtlOFS: 0x0080, muxMode: 3, selInputOFS: 0x0524, selInput: 1}
	MX6_PAD_GPIO1_IO09__USDHC2_RESET_B = &pad{padCtlOFS: 0x030C, padCtl: 0, muxCtlOFS: 0x0080, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_GPIO1_IO09__GPIO1_IO09     = &pad{padCtlOFS: 0x030C, padCtl: 0, muxCtlOFS: 0x0080, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_GPIO1_IO09__USDHC1_RESET_B = &pad{padCtlOFS: 0x030C, padCtl: 0, muxCtlOFS: 0x0080, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_GPIO1_IO09__UART5_DCE_CTS  = &pad{padCtlOFS: 0x030C, padCtl: 0, muxCtlOFS: 0x0080, muxMode: 8, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_GPIO1_IO09__UART5_DTE_RTS  = &pad{padCtlOFS: 0x030C, padCtl: 0, muxCtlOFS: 0x0080, muxMode: 8, selInputOFS: 0x0640, selInput: 2}

	MX6_PAD_UART1_TX_DATA__UART1_DCE_TX = &pad{padCtlOFS: 0x0310, padCtl: 0, muxCtlOFS: 0x0084, muxMode: 0, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_UART1_TX_DATA__UART1_DTE_RX  = &pad{padCtlOFS: 0x0310, padCtl: 0, muxCtlOFS: 0x0084, muxMode: 0, selInputOFS: 0x0624, selInput: 2}
	MX6_PAD_UART1_TX_DATA__ENET1_RDATA02 = &pad{padCtlOFS: 0x0310, padCtl: 0, muxCtlOFS: 0x0084, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART1_TX_DATA__I2C3_SCL      = &pad{padCtlOFS: 0x0310, padCtl: 0, muxCtlOFS: 0x0084, muxMode: IOMUX_CONFIG_SION | 2, selInputOFS: 0x05B4, selInput: 0}
	MX6_PAD_UART1_TX_DATA__CSI_DATA02    = &pad{padCtlOFS: 0x0310, padCtl: 0, muxCtlOFS: 0x0084, muxMode: 3, selInputOFS: 0x04C4, selInput: 1}
	MX6_PAD_UART1_TX_DATA__GPT1_COMPARE1 = &pad{padCtlOFS: 0x0310, padCtl: 0, muxCtlOFS: 0x0084, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART1_TX_DATA__GPIO1_IO16    = &pad{padCtlOFS: 0x0310, padCtl: 0, muxCtlOFS: 0x0084, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART1_TX_DATA__SPDIF_OUT     = &pad{padCtlOFS: 0x0310, padCtl: 0, muxCtlOFS: 0x0084, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_UART1_RX_DATA__UART1_DCE_RX = &pad{padCtlOFS: 0x0314, padCtl: 0, muxCtlOFS: 0x0088, muxMode: 0, selInputOFS: 0x0624, selInput: 3}

	MX6_PAD_UART1_RX_DATA__UART1_DTE_TX  = &pad{padCtlOFS: 0x0314, padCtl: 0, muxCtlOFS: 0x0088, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART1_RX_DATA__ENET1_RDATA03 = &pad{padCtlOFS: 0x0314, padCtl: 0, muxCtlOFS: 0x0088, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART1_RX_DATA__I2C3_SDA      = &pad{padCtlOFS: 0x0314, padCtl: 0, muxCtlOFS: 0x0088, muxMode: IOMUX_CONFIG_SION | 2, selInputOFS: 0x05B8, selInput: 0}
	MX6_PAD_UART1_RX_DATA__CSI_DATA03    = &pad{padCtlOFS: 0x0314, padCtl: 0, muxCtlOFS: 0x0088, muxMode: 3, selInputOFS: 0x04C8, selInput: 1}
	MX6_PAD_UART1_RX_DATA__GPT1_CLK      = &pad{padCtlOFS: 0x0314, padCtl: 0, muxCtlOFS: 0x0088, muxMode: 4, selInputOFS: 0x0594, selInput: 0}
	MX6_PAD_UART1_RX_DATA__GPIO1_IO17    = &pad{padCtlOFS: 0x0314, padCtl: 0, muxCtlOFS: 0x0088, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART1_RX_DATA__SPDIF_IN      = &pad{padCtlOFS: 0x0314, padCtl: 0, muxCtlOFS: 0x0088, muxMode: 8, selInputOFS: 0x0618, selInput: 1}

	MX6_PAD_UART1_CTS_B__UART1_DCE_CTS = &pad{padCtlOFS: 0x0318, padCtl: 0, muxCtlOFS: 0x008C, muxMode: 0, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_UART1_CTS_B__UART1_DTE_RTS        = &pad{padCtlOFS: 0x0318, padCtl: 0, muxCtlOFS: 0x008C, muxMode: 0, selInputOFS: 0x0620, selInput: 2}
	MX6_PAD_UART1_CTS_B__ENET1_RX_CLK         = &pad{padCtlOFS: 0x0318, padCtl: 0, muxCtlOFS: 0x008C, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART1_CTS_B__USDHC1_WP            = &pad{padCtlOFS: 0x0318, padCtl: 0, muxCtlOFS: 0x008C, muxMode: 2, selInputOFS: 0x066C, selInput: 1}
	MX6_PAD_UART1_CTS_B__CSI_DATA04           = &pad{padCtlOFS: 0x0318, padCtl: 0, muxCtlOFS: 0x008C, muxMode: 3, selInputOFS: 0x04D8, selInput: 0}
	MX6_PAD_UART1_CTS_B__ENET2_1588_EVENT1_IN = &pad{padCtlOFS: 0x0318, padCtl: 0, muxCtlOFS: 0x008C, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART1_CTS_B__GPIO1_IO18           = &pad{padCtlOFS: 0x0318, padCtl: 0, muxCtlOFS: 0x008C, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART1_CTS_B__USDHC2_WP            = &pad{padCtlOFS: 0x0318, padCtl: 0, muxCtlOFS: 0x008C, muxMode: 8, selInputOFS: 0x069C, selInput: 1}

	MX6_PAD_UART1_RTS_B__UART1_DCE_RTS = &pad{padCtlOFS: 0x031C, padCtl: 0, muxCtlOFS: 0x0090, muxMode: 0, selInputOFS: 0x0620, selInput: 3}

	MX6_PAD_UART1_RTS_B__UART1_DTE_CTS         = &pad{padCtlOFS: 0x031C, padCtl: 0, muxCtlOFS: 0x0090, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART1_RTS_B__ENET1_TX_ER           = &pad{padCtlOFS: 0x031C, padCtl: 0, muxCtlOFS: 0x0090, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART1_RTS_B__USDHC1_CD_B           = &pad{padCtlOFS: 0x031C, padCtl: 0, muxCtlOFS: 0x0090, muxMode: 2, selInputOFS: 0x0668, selInput: 1}
	MX6_PAD_UART1_RTS_B__CSI_DATA05            = &pad{padCtlOFS: 0x031C, padCtl: 0, muxCtlOFS: 0x0090, muxMode: 3, selInputOFS: 0x04CC, selInput: 1}
	MX6_PAD_UART1_RTS_B__ENET2_1588_EVENT1_OUT = &pad{padCtlOFS: 0x031C, padCtl: 0, muxCtlOFS: 0x0090, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART1_RTS_B__GPIO1_IO19            = &pad{padCtlOFS: 0x031C, padCtl: 0, muxCtlOFS: 0x0090, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART1_RTS_B__USDHC2_CD_B           = &pad{padCtlOFS: 0x031C, padCtl: 0, muxCtlOFS: 0x0090, muxMode: 8, selInputOFS: 0x0674, selInput: 2}

	MX6_PAD_UART2_TX_DATA__UART2_DCE_TX = &pad{padCtlOFS: 0x0320, padCtl: 0, muxCtlOFS: 0x0094, muxMode: 0, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_UART2_TX_DATA__UART2_DTE_RX  = &pad{padCtlOFS: 0x0320, padCtl: 0, muxCtlOFS: 0x0094, muxMode: 0, selInputOFS: 0x062C, selInput: 0}
	MX6_PAD_UART2_TX_DATA__ENET1_TDATA02 = &pad{padCtlOFS: 0x0320, padCtl: 0, muxCtlOFS: 0x0094, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART2_TX_DATA__I2C4_SCL      = &pad{padCtlOFS: 0x0320, padCtl: 0, muxCtlOFS: 0x0094, muxMode: IOMUX_CONFIG_SION | 2, selInputOFS: 0x05BC, selInput: 0}
	MX6_PAD_UART2_TX_DATA__CSI_DATA06    = &pad{padCtlOFS: 0x0320, padCtl: 0, muxCtlOFS: 0x0094, muxMode: 3, selInputOFS: 0x04DC, selInput: 0}
	MX6_PAD_UART2_TX_DATA__GPT1_CAPTURE1 = &pad{padCtlOFS: 0x0320, padCtl: 0, muxCtlOFS: 0x0094, muxMode: 4, selInputOFS: 0x058C, selInput: 1}
	MX6_PAD_UART2_TX_DATA__GPIO1_IO20    = &pad{padCtlOFS: 0x0320, padCtl: 0, muxCtlOFS: 0x0094, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART2_TX_DATA__ECSPI3_SS0    = &pad{padCtlOFS: 0x0320, padCtl: 0, muxCtlOFS: 0x0094, muxMode: 8, selInputOFS: 0x0560, selInput: 0}

	MX6_PAD_UART2_RX_DATA__UART2_DCE_RX = &pad{padCtlOFS: 0x0324, padCtl: 0, muxCtlOFS: 0x0098, muxMode: 0, selInputOFS: 0x062C, selInput: 1}

	MX6_PAD_UART2_RX_DATA__UART2_DTE_TX  = &pad{padCtlOFS: 0x0324, padCtl: 0, muxCtlOFS: 0x0098, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART2_RX_DATA__ENET1_TDATA03 = &pad{padCtlOFS: 0x0324, padCtl: 0, muxCtlOFS: 0x0098, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART2_RX_DATA__I2C4_SDA      = &pad{padCtlOFS: 0x0324, padCtl: 0, muxCtlOFS: 0x0098, muxMode: IOMUX_CONFIG_SION | 2, selInputOFS: 0x05C0, selInput: 0}
	MX6_PAD_UART2_RX_DATA__CSI_DATA07    = &pad{padCtlOFS: 0x0324, padCtl: 0, muxCtlOFS: 0x0098, muxMode: 3, selInputOFS: 0x04E0, selInput: 0}
	MX6_PAD_UART2_RX_DATA__GPT1_CAPTURE2 = &pad{padCtlOFS: 0x0324, padCtl: 0, muxCtlOFS: 0x0098, muxMode: 4, selInputOFS: 0x0590, selInput: 0}
	MX6_PAD_UART2_RX_DATA__GPIO1_IO21    = &pad{padCtlOFS: 0x0324, padCtl: 0, muxCtlOFS: 0x0098, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART2_RX_DATA__SJC_DONE      = &pad{padCtlOFS: 0x0324, padCtl: 0, muxCtlOFS: 0x0098, muxMode: 7, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART2_RX_DATA__ECSPI3_SCLK   = &pad{padCtlOFS: 0x0324, padCtl: 0, muxCtlOFS: 0x0098, muxMode: 8, selInputOFS: 0x0554, selInput: 0}

	MX6_PAD_UART2_CTS_B__UART2_DCE_CTS = &pad{padCtlOFS: 0x0328, padCtl: 0, muxCtlOFS: 0x009C, muxMode: 0, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_UART2_CTS_B__UART2_DTE_RTS = &pad{padCtlOFS: 0x0328, padCtl: 0, muxCtlOFS: 0x009C, muxMode: 0, selInputOFS: 0x0628, selInput: 0}
	MX6_PAD_UART2_CTS_B__ENET1_CRS     = &pad{padCtlOFS: 0x0328, padCtl: 0, muxCtlOFS: 0x009C, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART2_CTS_B__FLEXCAN2_TX   = &pad{padCtlOFS: 0x0328, padCtl: 0, muxCtlOFS: 0x009C, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART2_CTS_B__CSI_DATA08    = &pad{padCtlOFS: 0x0328, padCtl: 0, muxCtlOFS: 0x009C, muxMode: 3, selInputOFS: 0x04E4, selInput: 0}
	MX6_PAD_UART2_CTS_B__GPT1_COMPARE2 = &pad{padCtlOFS: 0x0328, padCtl: 0, muxCtlOFS: 0x009C, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART2_CTS_B__GPIO1_IO22    = &pad{padCtlOFS: 0x0328, padCtl: 0, muxCtlOFS: 0x009C, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART2_CTS_B__SJC_DE_B      = &pad{padCtlOFS: 0x0328, padCtl: 0, muxCtlOFS: 0x009C, muxMode: 7, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART2_CTS_B__ECSPI3_MOSI   = &pad{padCtlOFS: 0x0328, padCtl: 0, muxCtlOFS: 0x009C, muxMode: 8, selInputOFS: 0x055C, selInput: 0}

	MX6_PAD_UART2_RTS_B__UART2_DCE_RTS = &pad{padCtlOFS: 0x032C, padCtl: 0, muxCtlOFS: 0x00A0, muxMode: 0, selInputOFS: 0x0628, selInput: 1}

	MX6_PAD_UART2_RTS_B__UART2_DTE_CTS = &pad{padCtlOFS: 0x032C, padCtl: 0, muxCtlOFS: 0x00A0, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART2_RTS_B__ENET1_COL     = &pad{padCtlOFS: 0x032C, padCtl: 0, muxCtlOFS: 0x00A0, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART2_RTS_B__FLEXCAN2_RX   = &pad{padCtlOFS: 0x032C, padCtl: 0, muxCtlOFS: 0x00A0, muxMode: 2, selInputOFS: 0x0588, selInput: 0}
	MX6_PAD_UART2_RTS_B__CSI_DATA09    = &pad{padCtlOFS: 0x032C, padCtl: 0, muxCtlOFS: 0x00A0, muxMode: 3, selInputOFS: 0x04E8, selInput: 0}
	MX6_PAD_UART2_RTS_B__GPT1_COMPARE3 = &pad{padCtlOFS: 0x032C, padCtl: 0, muxCtlOFS: 0x00A0, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART2_RTS_B__GPIO1_IO23    = &pad{padCtlOFS: 0x032C, padCtl: 0, muxCtlOFS: 0x00A0, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART2_RTS_B__SJC_FAIL      = &pad{padCtlOFS: 0x032C, padCtl: 0, muxCtlOFS: 0x00A0, muxMode: 7, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART2_RTS_B__ECSPI3_MISO   = &pad{padCtlOFS: 0x032C, padCtl: 0, muxCtlOFS: 0x00A0, muxMode: 8, selInputOFS: 0x0558, selInput: 0}

	MX6_PAD_UART3_TX_DATA__UART3_DCE_TX = &pad{padCtlOFS: 0x0330, padCtl: 0, muxCtlOFS: 0x00A4, muxMode: 0, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_UART3_TX_DATA__UART3_DTE_RX   = &pad{padCtlOFS: 0x0330, padCtl: 0, muxCtlOFS: 0x00A4, muxMode: 0, selInputOFS: 0x0634, selInput: 0}
	MX6_PAD_UART3_TX_DATA__ENET2_RDATA02  = &pad{padCtlOFS: 0x0330, padCtl: 0, muxCtlOFS: 0x00A4, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART3_TX_DATA__SIM1_PORT0_PD  = &pad{padCtlOFS: 0x0330, padCtl: 0, muxCtlOFS: 0x00A4, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART3_TX_DATA__CSI_DATA01     = &pad{padCtlOFS: 0x0330, padCtl: 0, muxCtlOFS: 0x00A4, muxMode: 3, selInputOFS: 0x04D4, selInput: 0}
	MX6_PAD_UART3_TX_DATA__UART2_DCE_CTS  = &pad{padCtlOFS: 0x0330, padCtl: 0, muxCtlOFS: 0x00A4, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART3_TX_DATA__UART2_DTE_RTS  = &pad{padCtlOFS: 0x0330, padCtl: 0, muxCtlOFS: 0x00A4, muxMode: 4, selInputOFS: 0x0628, selInput: 2}
	MX6_PAD_UART3_TX_DATA__GPIO1_IO24     = &pad{padCtlOFS: 0x0330, padCtl: 0, muxCtlOFS: 0x00A4, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART3_TX_DATA__SJC_JTAG_ACT   = &pad{padCtlOFS: 0x0330, padCtl: 0, muxCtlOFS: 0x00A4, muxMode: 7, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART3_TX_DATA__ANATOP_OTG1_ID = &pad{padCtlOFS: 0x0330, padCtl: 0, muxCtlOFS: 0x00A4, muxMode: 8, selInputOFS: 0x04B8, selInput: 1}

	MX6_PAD_UART3_RX_DATA__UART3_DCE_RX = &pad{padCtlOFS: 0x0334, padCtl: 0, muxCtlOFS: 0x00A8, muxMode: 0, selInputOFS: 0x0634, selInput: 1}

	MX6_PAD_UART3_RX_DATA__UART3_DTE_TX  = &pad{padCtlOFS: 0x0334, padCtl: 0, muxCtlOFS: 0x00A8, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART3_RX_DATA__ENET2_RDATA03 = &pad{padCtlOFS: 0x0334, padCtl: 0, muxCtlOFS: 0x00A8, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART3_RX_DATA__SIM2_PORT0_PD = &pad{padCtlOFS: 0x0334, padCtl: 0, muxCtlOFS: 0x00A8, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART3_RX_DATA__CSI_DATA00    = &pad{padCtlOFS: 0x0334, padCtl: 0, muxCtlOFS: 0x00A8, muxMode: 3, selInputOFS: 0x04D0, selInput: 0}
	MX6_PAD_UART3_RX_DATA__UART2_DCE_RTS = &pad{padCtlOFS: 0x0334, padCtl: 0, muxCtlOFS: 0x00A8, muxMode: 4, selInputOFS: 0x0628, selInput: 3}
	MX6_PAD_UART3_RX_DATA__UART2_DTE_CTS = &pad{padCtlOFS: 0x0334, padCtl: 0, muxCtlOFS: 0x00A8, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART3_RX_DATA__GPIO1_IO25    = &pad{padCtlOFS: 0x0334, padCtl: 0, muxCtlOFS: 0x00A8, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART3_RX_DATA__EPIT1_OUT     = &pad{padCtlOFS: 0x0334, padCtl: 0, muxCtlOFS: 0x00A8, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_UART3_CTS_B__UART3_DCE_CTS = &pad{padCtlOFS: 0x0338, padCtl: 0, muxCtlOFS: 0x00AC, muxMode: 0, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_UART3_CTS_B__UART3_DTE_RTS        = &pad{padCtlOFS: 0x0338, padCtl: 0, muxCtlOFS: 0x00AC, muxMode: 0, selInputOFS: 0x0630, selInput: 0}
	MX6_PAD_UART3_CTS_B__ENET2_RX_CLK         = &pad{padCtlOFS: 0x0338, padCtl: 0, muxCtlOFS: 0x00AC, muxMode: IOMUX_CONFIG_SION | 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART3_CTS_B__FLEXCAN1_TX          = &pad{padCtlOFS: 0x0338, padCtl: 0, muxCtlOFS: 0x00AC, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART3_CTS_B__CSI_DATA10           = &pad{padCtlOFS: 0x0338, padCtl: 0, muxCtlOFS: 0x00AC, muxMode: 3, selInputOFS: 0x04EC, selInput: 0}
	MX6_PAD_UART3_CTS_B__ENET1_1588_EVENT1_IN = &pad{padCtlOFS: 0x0338, padCtl: 0, muxCtlOFS: 0x00AC, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART3_CTS_B__GPIO1_IO26           = &pad{padCtlOFS: 0x0338, padCtl: 0, muxCtlOFS: 0x00AC, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART3_CTS_B__EPIT2_OUT            = &pad{padCtlOFS: 0x0338, padCtl: 0, muxCtlOFS: 0x00AC, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_UART3_RTS_B__UART3_DCE_RTS = &pad{padCtlOFS: 0x033C, padCtl: 0, muxCtlOFS: 0x00B0, muxMode: 0, selInputOFS: 0x0630, selInput: 1}

	MX6_PAD_UART3_RTS_B__UART3_DTE_CTS         = &pad{padCtlOFS: 0x033C, padCtl: 0, muxCtlOFS: 0x00B0, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART3_RTS_B__ENET2_TX_ER           = &pad{padCtlOFS: 0x033C, padCtl: 0, muxCtlOFS: 0x00B0, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART3_RTS_B__FLEXCAN1_RX           = &pad{padCtlOFS: 0x033C, padCtl: 0, muxCtlOFS: 0x00B0, muxMode: 2, selInputOFS: 0x0584, selInput: 0}
	MX6_PAD_UART3_RTS_B__CSI_DATA11            = &pad{padCtlOFS: 0x033C, padCtl: 0, muxCtlOFS: 0x00B0, muxMode: 3, selInputOFS: 0x04F0, selInput: 0}
	MX6_PAD_UART3_RTS_B__ENET1_1588_EVENT1_OUT = &pad{padCtlOFS: 0x033C, padCtl: 0, muxCtlOFS: 0x00B0, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART3_RTS_B__GPIO1_IO27            = &pad{padCtlOFS: 0x033C, padCtl: 0, muxCtlOFS: 0x00B0, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART3_RTS_B__WDOG1_WDOG_B          = &pad{padCtlOFS: 0x033C, padCtl: 0, muxCtlOFS: 0x00B0, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_UART4_TX_DATA__UART4_DCE_TX = &pad{padCtlOFS: 0x0340, padCtl: 0, muxCtlOFS: 0x00B4, muxMode: 0, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_UART4_TX_DATA__UART4_DTE_RX        = &pad{padCtlOFS: 0x0340, padCtl: 0, muxCtlOFS: 0x00B4, muxMode: 0, selInputOFS: 0x063C, selInput: 0}
	MX6_PAD_UART4_TX_DATA__ENET2_TDATA02       = &pad{padCtlOFS: 0x0340, padCtl: 0, muxCtlOFS: 0x00B4, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART4_TX_DATA__I2C1_SCL            = &pad{padCtlOFS: 0x0340, padCtl: 0, muxCtlOFS: 0x00B4, muxMode: IOMUX_CONFIG_SION | 2, selInputOFS: 0x05A4, selInput: 1}
	MX6_PAD_UART4_TX_DATA__CSI_DATA12          = &pad{padCtlOFS: 0x0340, padCtl: 0, muxCtlOFS: 0x00B4, muxMode: 3, selInputOFS: 0x04F4, selInput: 0}
	MX6_PAD_UART4_TX_DATA__CSU_CSU_ALARM_AUT02 = &pad{padCtlOFS: 0x0340, padCtl: 0, muxCtlOFS: 0x00B4, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART4_TX_DATA__GPIO1_IO28          = &pad{padCtlOFS: 0x0340, padCtl: 0, muxCtlOFS: 0x00B4, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART4_TX_DATA__ECSPI2_SCLK         = &pad{padCtlOFS: 0x0340, padCtl: 0, muxCtlOFS: 0x00B4, muxMode: 8, selInputOFS: 0x0544, selInput: 1}

	MX6_PAD_UART4_RX_DATA__UART4_DCE_RX = &pad{padCtlOFS: 0x0344, padCtl: 0, muxCtlOFS: 0x00B8, muxMode: 0, selInputOFS: 0x063C, selInput: 1}

	MX6_PAD_UART4_RX_DATA__UART4_DTE_TX        = &pad{padCtlOFS: 0x0344, padCtl: 0, muxCtlOFS: 0x00B8, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART4_RX_DATA__ENET2_TDATA03       = &pad{padCtlOFS: 0x0344, padCtl: 0, muxCtlOFS: 0x00B8, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART4_RX_DATA__I2C1_SDA            = &pad{padCtlOFS: 0x0344, padCtl: 0, muxCtlOFS: 0x00B8, muxMode: IOMUX_CONFIG_SION | 2, selInputOFS: 0x05A8, selInput: 2}
	MX6_PAD_UART4_RX_DATA__CSI_DATA13          = &pad{padCtlOFS: 0x0344, padCtl: 0, muxCtlOFS: 0x00B8, muxMode: 3, selInputOFS: 0x04F8, selInput: 0}
	MX6_PAD_UART4_RX_DATA__CSU_CSU_ALARM_AUT01 = &pad{padCtlOFS: 0x0344, padCtl: 0, muxCtlOFS: 0x00B8, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART4_RX_DATA__GPIO1_IO29          = &pad{padCtlOFS: 0x0344, padCtl: 0, muxCtlOFS: 0x00B8, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART4_RX_DATA__ECSPI2_SS0          = &pad{padCtlOFS: 0x0344, padCtl: 0, muxCtlOFS: 0x00B8, muxMode: 8, selInputOFS: 0x0550, selInput: 1}
	MX6_PAD_UART5_TX_DATA__GPIO1_IO30          = &pad{padCtlOFS: 0x0348, padCtl: 0, muxCtlOFS: 0x00BC, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART5_TX_DATA__ECSPI2_MOSI         = &pad{padCtlOFS: 0x0348, padCtl: 0, muxCtlOFS: 0x00BC, muxMode: 8, selInputOFS: 0x054C, selInput: 0}

	MX6_PAD_UART5_TX_DATA__UART5_DCE_TX = &pad{padCtlOFS: 0x0348, padCtl: 0, muxCtlOFS: 0x00BC, muxMode: 0, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_UART5_TX_DATA__UART5_DTE_RX        = &pad{padCtlOFS: 0x0348, padCtl: 0, muxCtlOFS: 0x00BC, muxMode: 0, selInputOFS: 0x0644, selInput: 4}
	MX6_PAD_UART5_TX_DATA__ENET2_CRS           = &pad{padCtlOFS: 0x0348, padCtl: 0, muxCtlOFS: 0x00BC, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART5_TX_DATA__I2C2_SCL            = &pad{padCtlOFS: 0x0348, padCtl: 0, muxCtlOFS: 0x00BC, muxMode: IOMUX_CONFIG_SION | 2, selInputOFS: 0x05AC, selInput: 2}
	MX6_PAD_UART5_TX_DATA__CSI_DATA14          = &pad{padCtlOFS: 0x0348, padCtl: 0, muxCtlOFS: 0x00BC, muxMode: 3, selInputOFS: 0x04FC, selInput: 0}
	MX6_PAD_UART5_TX_DATA__CSU_CSU_ALARM_AUT00 = &pad{padCtlOFS: 0x0348, padCtl: 0, muxCtlOFS: 0x00BC, muxMode: 4, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_UART5_RX_DATA__UART5_DCE_RX = &pad{padCtlOFS: 0x034C, padCtl: 0, muxCtlOFS: 0x00C0, muxMode: 0, selInputOFS: 0x0644, selInput: 5}

	MX6_PAD_UART5_RX_DATA__UART5_DTE_TX    = &pad{padCtlOFS: 0x034C, padCtl: 0, muxCtlOFS: 0x00C0, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART5_RX_DATA__ENET2_COL       = &pad{padCtlOFS: 0x034C, padCtl: 0, muxCtlOFS: 0x00C0, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART5_RX_DATA__I2C2_SDA        = &pad{padCtlOFS: 0x034C, padCtl: 0, muxCtlOFS: 0x00C0, muxMode: IOMUX_CONFIG_SION | 2, selInputOFS: 0x05B0, selInput: 2}
	MX6_PAD_UART5_RX_DATA__CSI_DATA15      = &pad{padCtlOFS: 0x034C, padCtl: 0, muxCtlOFS: 0x00C0, muxMode: 3, selInputOFS: 0x0500, selInput: 0}
	MX6_PAD_UART5_RX_DATA__CSU_CSU_INT_DEB = &pad{padCtlOFS: 0x034C, padCtl: 0, muxCtlOFS: 0x00C0, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART5_RX_DATA__GPIO1_IO31      = &pad{padCtlOFS: 0x034C, padCtl: 0, muxCtlOFS: 0x00C0, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_UART5_RX_DATA__ECSPI2_MISO     = &pad{padCtlOFS: 0x034C, padCtl: 0, muxCtlOFS: 0x00C0, muxMode: 8, selInputOFS: 0x0548, selInput: 1}

	MX6_PAD_ENET1_RX_DATA0__ENET1_RDATA00 = &pad{padCtlOFS: 0x0350, padCtl: 0, muxCtlOFS: 0x00C4, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET1_RX_DATA0__UART4_DCE_RTS = &pad{padCtlOFS: 0x0350, padCtl: 0, muxCtlOFS: 0x00C4, muxMode: 1, selInputOFS: 0x0638, selInput: 0}
	MX6_PAD_ENET1_RX_DATA0__UART4_DTE_CTS = &pad{padCtlOFS: 0x0350, padCtl: 0, muxCtlOFS: 0x00C4, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET1_RX_DATA0__PWM1_OUT      = &pad{padCtlOFS: 0x0350, padCtl: 0, muxCtlOFS: 0x00C4, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET1_RX_DATA0__CSI_DATA16    = &pad{padCtlOFS: 0x0350, padCtl: 0, muxCtlOFS: 0x00C4, muxMode: 3, selInputOFS: 0x0504, selInput: 0}
	MX6_PAD_ENET1_RX_DATA0__FLEXCAN1_TX   = &pad{padCtlOFS: 0x0350, padCtl: 0, muxCtlOFS: 0x00C4, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET1_RX_DATA0__GPIO2_IO00    = &pad{padCtlOFS: 0x0350, padCtl: 0, muxCtlOFS: 0x00C4, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET1_RX_DATA0__KPP_ROW00     = &pad{padCtlOFS: 0x0350, padCtl: 0, muxCtlOFS: 0x00C4, muxMode: 6, selInputOFS: 0x05D0, selInput: 0}
	MX6_PAD_ENET1_RX_DATA0__USDHC1_LCTL   = &pad{padCtlOFS: 0x0350, padCtl: 0, muxCtlOFS: 0x00C4, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_ENET1_RX_DATA1__ENET1_RDATA01 = &pad{padCtlOFS: 0x0354, padCtl: 0, muxCtlOFS: 0x00C8, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET1_RX_DATA1__UART4_DCE_CTS = &pad{padCtlOFS: 0x0354, padCtl: 0, muxCtlOFS: 0x00C8, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET1_RX_DATA1__UART4_DTE_RTS = &pad{padCtlOFS: 0x0354, padCtl: 0, muxCtlOFS: 0x00C8, muxMode: 1, selInputOFS: 0x0638, selInput: 1}
	MX6_PAD_ENET1_RX_DATA1__PWM2_OUT      = &pad{padCtlOFS: 0x0354, padCtl: 0, muxCtlOFS: 0x00C8, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET1_RX_DATA1__CSI_DATA17    = &pad{padCtlOFS: 0x0354, padCtl: 0, muxCtlOFS: 0x00C8, muxMode: 3, selInputOFS: 0x0508, selInput: 0}
	MX6_PAD_ENET1_RX_DATA1__FLEXCAN1_RX   = &pad{padCtlOFS: 0x0354, padCtl: 0, muxCtlOFS: 0x00C8, muxMode: 4, selInputOFS: 0x0584, selInput: 1}
	MX6_PAD_ENET1_RX_DATA1__GPIO2_IO01    = &pad{padCtlOFS: 0x0354, padCtl: 0, muxCtlOFS: 0x00C8, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET1_RX_DATA1__KPP_COL00     = &pad{padCtlOFS: 0x0354, padCtl: 0, muxCtlOFS: 0x00C8, muxMode: 6, selInputOFS: 0x05C4, selInput: 0}
	MX6_PAD_ENET1_RX_DATA1__USDHC2_LCTL   = &pad{padCtlOFS: 0x0354, padCtl: 0, muxCtlOFS: 0x00C8, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_ENET1_RX_EN__ENET1_RX_EN    = &pad{padCtlOFS: 0x0358, padCtl: 0, muxCtlOFS: 0x00CC, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET1_RX_EN__UART5_DCE_RTS  = &pad{padCtlOFS: 0x0358, padCtl: 0, muxCtlOFS: 0x00CC, muxMode: 1, selInputOFS: 0x0640, selInput: 3}
	MX6_PAD_ENET1_RX_EN__UART5_DTE_CTS  = &pad{padCtlOFS: 0x0358, padCtl: 0, muxCtlOFS: 0x00CC, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET1_RX_EN__CSI_DATA18     = &pad{padCtlOFS: 0x0358, padCtl: 0, muxCtlOFS: 0x00CC, muxMode: 3, selInputOFS: 0x050C, selInput: 0}
	MX6_PAD_ENET1_RX_EN__FLEXCAN2_TX    = &pad{padCtlOFS: 0x0358, padCtl: 0, muxCtlOFS: 0x00CC, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET1_RX_EN__GPIO2_IO02     = &pad{padCtlOFS: 0x0358, padCtl: 0, muxCtlOFS: 0x00CC, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET1_RX_EN__KPP_ROW01      = &pad{padCtlOFS: 0x0358, padCtl: 0, muxCtlOFS: 0x00CC, muxMode: 6, selInputOFS: 0x05D4, selInput: 0}
	MX6_PAD_ENET1_RX_EN__USDHC1_VSELECT = &pad{padCtlOFS: 0x0358, padCtl: 0, muxCtlOFS: 0x00CC, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_ENET1_TX_DATA0__ENET1_TDATA00  = &pad{padCtlOFS: 0x035C, padCtl: 0, muxCtlOFS: 0x00D0, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET1_TX_DATA0__UART5_DCE_CTS  = &pad{padCtlOFS: 0x035C, padCtl: 0, muxCtlOFS: 0x00D0, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET1_TX_DATA0__UART5_DTE_RTS  = &pad{padCtlOFS: 0x035C, padCtl: 0, muxCtlOFS: 0x00D0, muxMode: 1, selInputOFS: 0x0640, selInput: 4}
	MX6_PAD_ENET1_TX_DATA0__CSI_DATA19     = &pad{padCtlOFS: 0x035C, padCtl: 0, muxCtlOFS: 0x00D0, muxMode: 3, selInputOFS: 0x0510, selInput: 0}
	MX6_PAD_ENET1_TX_DATA0__FLEXCAN2_RX    = &pad{padCtlOFS: 0x035C, padCtl: 0, muxCtlOFS: 0x00D0, muxMode: 4, selInputOFS: 0x0588, selInput: 1}
	MX6_PAD_ENET1_TX_DATA0__GPIO2_IO03     = &pad{padCtlOFS: 0x035C, padCtl: 0, muxCtlOFS: 0x00D0, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET1_TX_DATA0__KPP_COL01      = &pad{padCtlOFS: 0x035C, padCtl: 0, muxCtlOFS: 0x00D0, muxMode: 6, selInputOFS: 0x05C8, selInput: 0}
	MX6_PAD_ENET1_TX_DATA0__USDHC2_VSELECT = &pad{padCtlOFS: 0x035C, padCtl: 0, muxCtlOFS: 0x00D0, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_ENET1_TX_DATA1__ENET1_TDATA01        = &pad{padCtlOFS: 0x0360, padCtl: 0, muxCtlOFS: 0x00D4, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET1_TX_DATA1__UART6_DCE_CTS        = &pad{padCtlOFS: 0x0360, padCtl: 0, muxCtlOFS: 0x00D4, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET1_TX_DATA1__UART6_DTE_RTS        = &pad{padCtlOFS: 0x0360, padCtl: 0, muxCtlOFS: 0x00D4, muxMode: 1, selInputOFS: 0x0648, selInput: 2}
	MX6_PAD_ENET1_TX_DATA1__PWM5_OUT             = &pad{padCtlOFS: 0x0360, padCtl: 0, muxCtlOFS: 0x00D4, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET1_TX_DATA1__CSI_DATA20           = &pad{padCtlOFS: 0x0360, padCtl: 0, muxCtlOFS: 0x00D4, muxMode: 3, selInputOFS: 0x0514, selInput: 0}
	MX6_PAD_ENET1_TX_DATA1__ENET2_MDIO           = &pad{padCtlOFS: 0x0360, padCtl: 0, muxCtlOFS: 0x00D4, muxMode: 4, selInputOFS: 0x0580, selInput: 1}
	MX6_PAD_ENET1_TX_DATA1__GPIO2_IO04           = &pad{padCtlOFS: 0x0360, padCtl: 0, muxCtlOFS: 0x00D4, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET1_TX_DATA1__KPP_ROW02            = &pad{padCtlOFS: 0x0360, padCtl: 0, muxCtlOFS: 0x00D4, muxMode: 6, selInputOFS: 0x05D8, selInput: 0}
	MX6_PAD_ENET1_TX_DATA1__WDOG1_WDOG_RST_B_DEB = &pad{padCtlOFS: 0x0360, padCtl: 0, muxCtlOFS: 0x00D4, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_ENET1_TX_EN__ENET1_TX_EN          = &pad{padCtlOFS: 0x0364, padCtl: 0, muxCtlOFS: 0x00D8, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET1_TX_EN__UART6_DCE_RTS        = &pad{padCtlOFS: 0x0364, padCtl: 0, muxCtlOFS: 0x00D8, muxMode: 1, selInputOFS: 0x0648, selInput: 3}
	MX6_PAD_ENET1_TX_EN__UART6_DTE_CTS        = &pad{padCtlOFS: 0x0364, padCtl: 0, muxCtlOFS: 0x00D8, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET1_TX_EN__PWM6_OUT             = &pad{padCtlOFS: 0x0364, padCtl: 0, muxCtlOFS: 0x00D8, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET1_TX_EN__CSI_DATA21           = &pad{padCtlOFS: 0x0364, padCtl: 0, muxCtlOFS: 0x00D8, muxMode: 3, selInputOFS: 0x0518, selInput: 0}
	MX6_PAD_ENET1_TX_EN__ENET2_MDC            = &pad{padCtlOFS: 0x0364, padCtl: 0, muxCtlOFS: 0x00D8, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET1_TX_EN__GPIO2_IO05           = &pad{padCtlOFS: 0x0364, padCtl: 0, muxCtlOFS: 0x00D8, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET1_TX_EN__KPP_COL02            = &pad{padCtlOFS: 0x0364, padCtl: 0, muxCtlOFS: 0x00D8, muxMode: 6, selInputOFS: 0x05CC, selInput: 0}
	MX6_PAD_ENET1_TX_EN__WDOG2_WDOG_RST_B_DEB = &pad{padCtlOFS: 0x0364, padCtl: 0, muxCtlOFS: 0x00D8, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_ENET1_TX_CLK__ENET1_TX_CLK   = &pad{padCtlOFS: 0x0368, padCtl: 0, muxCtlOFS: 0x00DC, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET1_TX_CLK__UART7_DCE_CTS  = &pad{padCtlOFS: 0x0368, padCtl: 0, muxCtlOFS: 0x00DC, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET1_TX_CLK__UART7_DTE_RTS  = &pad{padCtlOFS: 0x0368, padCtl: 0, muxCtlOFS: 0x00DC, muxMode: 1, selInputOFS: 0x0650, selInput: 0}
	MX6_PAD_ENET1_TX_CLK__PWM7_OUT       = &pad{padCtlOFS: 0x0368, padCtl: 0, muxCtlOFS: 0x00DC, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET1_TX_CLK__CSI_DATA22     = &pad{padCtlOFS: 0x0368, padCtl: 0, muxCtlOFS: 0x00DC, muxMode: 3, selInputOFS: 0x051C, selInput: 0}
	MX6_PAD_ENET1_TX_CLK__ENET1_REF_CLK1 = &pad{padCtlOFS: 0x0368, padCtl: 0, muxCtlOFS: 0x00DC, muxMode: IOMUX_CONFIG_SION | 4, selInputOFS: 0x0574, selInput: 2}
	MX6_PAD_ENET1_TX_CLK__GPIO2_IO06     = &pad{padCtlOFS: 0x0368, padCtl: 0, muxCtlOFS: 0x00DC, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET1_TX_CLK__KPP_ROW03      = &pad{padCtlOFS: 0x0368, padCtl: 0, muxCtlOFS: 0x00DC, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET1_TX_CLK__GPT1_CLK       = &pad{padCtlOFS: 0x0368, padCtl: 0, muxCtlOFS: 0x00DC, muxMode: 8, selInputOFS: 0x0594, selInput: 1}

	MX6_PAD_ENET1_RX_ER__ENET1_RX_ER   = &pad{padCtlOFS: 0x036C, padCtl: 0, muxCtlOFS: 0x00E0, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET1_RX_ER__UART7_DCE_RTS = &pad{padCtlOFS: 0x036C, padCtl: 0, muxCtlOFS: 0x00E0, muxMode: 1, selInputOFS: 0x0650, selInput: 1}
	MX6_PAD_ENET1_RX_ER__UART7_DTE_CTS = &pad{padCtlOFS: 0x036C, padCtl: 0, muxCtlOFS: 0x00E0, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET1_RX_ER__PWM8_OUT      = &pad{padCtlOFS: 0x036C, padCtl: 0, muxCtlOFS: 0x00E0, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET1_RX_ER__CSI_DATA23    = &pad{padCtlOFS: 0x036C, padCtl: 0, muxCtlOFS: 0x00E0, muxMode: 3, selInputOFS: 0x0520, selInput: 0}
	MX6_PAD_ENET1_RX_ER__EIM_CRE       = &pad{padCtlOFS: 0x036C, padCtl: 0, muxCtlOFS: 0x00E0, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET1_RX_ER__GPIO2_IO07    = &pad{padCtlOFS: 0x036C, padCtl: 0, muxCtlOFS: 0x00E0, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET1_RX_ER__KPP_COL03     = &pad{padCtlOFS: 0x036C, padCtl: 0, muxCtlOFS: 0x00E0, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET1_RX_ER__GPT1_CAPTURE2 = &pad{padCtlOFS: 0x036C, padCtl: 0, muxCtlOFS: 0x00E0, muxMode: 8, selInputOFS: 0x0590, selInput: 1}

	MX6_PAD_ENET2_RX_DATA0__ENET2_RDATA00   = &pad{padCtlOFS: 0x0370, padCtl: 0, muxCtlOFS: 0x00E4, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET2_RX_DATA0__UART6_DCE_TX    = &pad{padCtlOFS: 0x0370, padCtl: 0, muxCtlOFS: 0x00E4, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET2_RX_DATA0__UART6_DTE_RX    = &pad{padCtlOFS: 0x0370, padCtl: 0, muxCtlOFS: 0x00E4, muxMode: 1, selInputOFS: 0x064C, selInput: 1}
	MX6_PAD_ENET2_RX_DATA0__SIM1_PORT0_TRXD = &pad{padCtlOFS: 0x0370, padCtl: 0, muxCtlOFS: 0x00E4, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET2_RX_DATA0__I2C3_SCL        = &pad{padCtlOFS: 0x0370, padCtl: 0, muxCtlOFS: 0x00E4, muxMode: IOMUX_CONFIG_SION | 3, selInputOFS: 0x05B4, selInput: 1}
	MX6_PAD_ENET2_RX_DATA0__ENET1_MDIO      = &pad{padCtlOFS: 0x0370, padCtl: 0, muxCtlOFS: 0x00E4, muxMode: 4, selInputOFS: 0x0578, selInput: 1}
	MX6_PAD_ENET2_RX_DATA0__GPIO2_IO08      = &pad{padCtlOFS: 0x0370, padCtl: 0, muxCtlOFS: 0x00E4, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET2_RX_DATA0__KPP_ROW04       = &pad{padCtlOFS: 0x0370, padCtl: 0, muxCtlOFS: 0x00E4, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET2_RX_DATA0__USB_OTG1_PWR    = &pad{padCtlOFS: 0x0370, padCtl: 0, muxCtlOFS: 0x00E4, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_ENET2_RX_DATA1__ENET2_RDATA01  = &pad{padCtlOFS: 0x0374, padCtl: 0, muxCtlOFS: 0x00E8, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET2_RX_DATA1__UART6_DCE_RX   = &pad{padCtlOFS: 0x0374, padCtl: 0, muxCtlOFS: 0x00E8, muxMode: 1, selInputOFS: 0x064C, selInput: 2}
	MX6_PAD_ENET2_RX_DATA1__UART6_DTE_TX   = &pad{padCtlOFS: 0x0374, padCtl: 0, muxCtlOFS: 0x00E8, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET2_RX_DATA1__SIM1_PORT0_CLK = &pad{padCtlOFS: 0x0374, padCtl: 0, muxCtlOFS: 0x00E8, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET2_RX_DATA1__I2C3_SDA       = &pad{padCtlOFS: 0x0374, padCtl: 0, muxCtlOFS: 0x00E8, muxMode: IOMUX_CONFIG_SION | 3, selInputOFS: 0x05B8, selInput: 1}
	MX6_PAD_ENET2_RX_DATA1__ENET1_MDC      = &pad{padCtlOFS: 0x0374, padCtl: 0, muxCtlOFS: 0x00E8, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET2_RX_DATA1__GPIO2_IO09     = &pad{padCtlOFS: 0x0374, padCtl: 0, muxCtlOFS: 0x00E8, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET2_RX_DATA1__KPP_COL04      = &pad{padCtlOFS: 0x0374, padCtl: 0, muxCtlOFS: 0x00E8, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET2_RX_DATA1__USB_OTG1_OC    = &pad{padCtlOFS: 0x0374, padCtl: 0, muxCtlOFS: 0x00E8, muxMode: 8, selInputOFS: 0x0664, selInput: 1}

	MX6_PAD_ENET2_RX_EN__ENET2_RX_EN       = &pad{padCtlOFS: 0x0378, padCtl: 0, muxCtlOFS: 0x00EC, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET2_RX_EN__UART7_DCE_TX      = &pad{padCtlOFS: 0x0378, padCtl: 0, muxCtlOFS: 0x00EC, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET2_RX_EN__UART7_DTE_RX      = &pad{padCtlOFS: 0x0378, padCtl: 0, muxCtlOFS: 0x00EC, muxMode: 1, selInputOFS: 0x0654, selInput: 0}
	MX6_PAD_ENET2_RX_EN__SIM1_PORT0_RST_B  = &pad{padCtlOFS: 0x0378, padCtl: 0, muxCtlOFS: 0x00EC, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET2_RX_EN__I2C4_SCL          = &pad{padCtlOFS: 0x0378, padCtl: 0, muxCtlOFS: 0x00EC, muxMode: IOMUX_CONFIG_SION | 3, selInputOFS: 0x05BC, selInput: 1}
	MX6_PAD_ENET2_RX_EN__EIM_ADDR26        = &pad{padCtlOFS: 0x0378, padCtl: 0, muxCtlOFS: 0x00EC, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET2_RX_EN__GPIO2_IO10        = &pad{padCtlOFS: 0x0378, padCtl: 0, muxCtlOFS: 0x00EC, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET2_RX_EN__KPP_ROW05         = &pad{padCtlOFS: 0x0378, padCtl: 0, muxCtlOFS: 0x00EC, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET2_RX_EN__ENET1_REF_CLK_25M = &pad{padCtlOFS: 0x0378, padCtl: 0, muxCtlOFS: 0x00EC, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_ENET2_TX_DATA0__ENET2_TDATA00   = &pad{padCtlOFS: 0x037C, padCtl: 0, muxCtlOFS: 0x00F0, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET2_TX_DATA0__UART7_DCE_RX    = &pad{padCtlOFS: 0x037C, padCtl: 0, muxCtlOFS: 0x00F0, muxMode: 1, selInputOFS: 0x0654, selInput: 1}
	MX6_PAD_ENET2_TX_DATA0__UART7_DTE_TX    = &pad{padCtlOFS: 0x037C, padCtl: 0, muxCtlOFS: 0x00F0, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET2_TX_DATA0__SIM1_PORT0_SVEN = &pad{padCtlOFS: 0x037C, padCtl: 0, muxCtlOFS: 0x00F0, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET2_TX_DATA0__I2C4_SDA        = &pad{padCtlOFS: 0x037C, padCtl: 0, muxCtlOFS: 0x00F0, muxMode: IOMUX_CONFIG_SION | 3, selInputOFS: 0x05C0, selInput: 1}
	MX6_PAD_ENET2_TX_DATA0__EIM_EB_B02      = &pad{padCtlOFS: 0x037C, padCtl: 0, muxCtlOFS: 0x00F0, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET2_TX_DATA0__GPIO2_IO11      = &pad{padCtlOFS: 0x037C, padCtl: 0, muxCtlOFS: 0x00F0, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET2_TX_DATA0__KPP_COL05       = &pad{padCtlOFS: 0x037C, padCtl: 0, muxCtlOFS: 0x00F0, muxMode: 6, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_ENET2_TX_DATA1__ENET2_TDATA01   = &pad{padCtlOFS: 0x0380, padCtl: 0, muxCtlOFS: 0x00F4, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET2_TX_DATA1__UART8_DCE_TX    = &pad{padCtlOFS: 0x0380, padCtl: 0, muxCtlOFS: 0x00F4, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET2_TX_DATA1__UART8_DTE_RX    = &pad{padCtlOFS: 0x0380, padCtl: 0, muxCtlOFS: 0x00F4, muxMode: 1, selInputOFS: 0x065C, selInput: 0}
	MX6_PAD_ENET2_TX_DATA1__SIM2_PORT0_TRXD = &pad{padCtlOFS: 0x0380, padCtl: 0, muxCtlOFS: 0x00F4, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET2_TX_DATA1__ECSPI4_SCLK     = &pad{padCtlOFS: 0x0380, padCtl: 0, muxCtlOFS: 0x00F4, muxMode: 3, selInputOFS: 0x0564, selInput: 0}
	MX6_PAD_ENET2_TX_DATA1__EIM_EB_B03      = &pad{padCtlOFS: 0x0380, padCtl: 0, muxCtlOFS: 0x00F4, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET2_TX_DATA1__GPIO2_IO12      = &pad{padCtlOFS: 0x0380, padCtl: 0, muxCtlOFS: 0x00F4, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET2_TX_DATA1__KPP_ROW06       = &pad{padCtlOFS: 0x0380, padCtl: 0, muxCtlOFS: 0x00F4, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET2_TX_DATA1__USB_OTG2_PWR    = &pad{padCtlOFS: 0x0380, padCtl: 0, muxCtlOFS: 0x00F4, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_ENET2_TX_EN__ENET2_TX_EN      = &pad{padCtlOFS: 0x0384, padCtl: 0, muxCtlOFS: 0x00F8, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET2_TX_EN__UART8_DCE_RX     = &pad{padCtlOFS: 0x0384, padCtl: 0, muxCtlOFS: 0x00F8, muxMode: 1, selInputOFS: 0x065C, selInput: 1}
	MX6_PAD_ENET2_TX_EN__UART8_DTE_TX     = &pad{padCtlOFS: 0x0384, padCtl: 0, muxCtlOFS: 0x00F8, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET2_TX_EN__SIM2_PORT0_CLK   = &pad{padCtlOFS: 0x0384, padCtl: 0, muxCtlOFS: 0x00F8, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET2_TX_EN__ECSPI4_MOSI      = &pad{padCtlOFS: 0x0384, padCtl: 0, muxCtlOFS: 0x00F8, muxMode: 3, selInputOFS: 0x056C, selInput: 0}
	MX6_PAD_ENET2_TX_EN__EIM_ACLK_FREERUN = &pad{padCtlOFS: 0x0384, padCtl: 0, muxCtlOFS: 0x00F8, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET2_TX_EN__GPIO2_IO13       = &pad{padCtlOFS: 0x0384, padCtl: 0, muxCtlOFS: 0x00F8, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET2_TX_EN__KPP_COL06        = &pad{padCtlOFS: 0x0384, padCtl: 0, muxCtlOFS: 0x00F8, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET2_TX_EN__USB_OTG2_OC      = &pad{padCtlOFS: 0x0384, padCtl: 0, muxCtlOFS: 0x00F8, muxMode: 8, selInputOFS: 0x0660, selInput: 1}

	MX6_PAD_ENET2_TX_CLK__ENET2_TX_CLK     = &pad{padCtlOFS: 0x0388, padCtl: 0, muxCtlOFS: 0x00FC, muxMode: IOMUX_CONFIG_SION | 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET2_TX_CLK__UART8_DCE_CTS    = &pad{padCtlOFS: 0x0388, padCtl: 0, muxCtlOFS: 0x00FC, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET2_TX_CLK__UART8_DTE_RTS    = &pad{padCtlOFS: 0x0388, padCtl: 0, muxCtlOFS: 0x00FC, muxMode: 1, selInputOFS: 0x0658, selInput: 0}
	MX6_PAD_ENET2_TX_CLK__SIM2_PORT0_RST_B = &pad{padCtlOFS: 0x0388, padCtl: 0, muxCtlOFS: 0x00FC, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET2_TX_CLK__ECSPI4_MISO      = &pad{padCtlOFS: 0x0388, padCtl: 0, muxCtlOFS: 0x00FC, muxMode: 3, selInputOFS: 0x0568, selInput: 0}
	MX6_PAD_ENET2_TX_CLK__ENET2_REF_CLK2   = &pad{padCtlOFS: 0x0388, padCtl: 0, muxCtlOFS: 0x00FC, muxMode: IOMUX_CONFIG_SION | 4, selInputOFS: 0x057C, selInput: 2}
	MX6_PAD_ENET2_TX_CLK__GPIO2_IO14       = &pad{padCtlOFS: 0x0388, padCtl: 0, muxCtlOFS: 0x00FC, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET2_TX_CLK__KPP_ROW07        = &pad{padCtlOFS: 0x0388, padCtl: 0, muxCtlOFS: 0x00FC, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET2_TX_CLK__ANATOP_OTG2_ID   = &pad{padCtlOFS: 0x0388, padCtl: 0, muxCtlOFS: 0x00FC, muxMode: 8, selInputOFS: 0x04BC, selInput: 1}

	MX6_PAD_ENET2_RX_ER__ENET2_RX_ER     = &pad{padCtlOFS: 0x038C, padCtl: 0, muxCtlOFS: 0x0100, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET2_RX_ER__UART8_DCE_RTS   = &pad{padCtlOFS: 0x038C, padCtl: 0, muxCtlOFS: 0x0100, muxMode: 1, selInputOFS: 0x0658, selInput: 1}
	MX6_PAD_ENET2_RX_ER__UART8_DTE_CTS   = &pad{padCtlOFS: 0x038C, padCtl: 0, muxCtlOFS: 0x0100, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET2_RX_ER__SIM2_PORT0_SVEN = &pad{padCtlOFS: 0x038C, padCtl: 0, muxCtlOFS: 0x0100, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET2_RX_ER__ECSPI4_SS0      = &pad{padCtlOFS: 0x038C, padCtl: 0, muxCtlOFS: 0x0100, muxMode: 3, selInputOFS: 0x0570, selInput: 0}
	MX6_PAD_ENET2_RX_ER__EIM_ADDR25      = &pad{padCtlOFS: 0x038C, padCtl: 0, muxCtlOFS: 0x0100, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET2_RX_ER__GPIO2_IO15      = &pad{padCtlOFS: 0x038C, padCtl: 0, muxCtlOFS: 0x0100, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET2_RX_ER__KPP_COL07       = &pad{padCtlOFS: 0x038C, padCtl: 0, muxCtlOFS: 0x0100, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_ENET2_RX_ER__WDOG1_WDOG_ANY  = &pad{padCtlOFS: 0x038C, padCtl: 0, muxCtlOFS: 0x0100, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_LCD_CLK__LCDIF_CLK            = &pad{padCtlOFS: 0x0390, padCtl: 0, muxCtlOFS: 0x0104, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_CLK__LCDIF_WR_RWN         = &pad{padCtlOFS: 0x0390, padCtl: 0, muxCtlOFS: 0x0104, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_CLK__UART4_DCE_TX         = &pad{padCtlOFS: 0x0390, padCtl: 0, muxCtlOFS: 0x0104, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_CLK__UART4_DTE_RX         = &pad{padCtlOFS: 0x0390, padCtl: 0, muxCtlOFS: 0x0104, muxMode: 2, selInputOFS: 0x063C, selInput: 2}
	MX6_PAD_LCD_CLK__SAI3_MCLK            = &pad{padCtlOFS: 0x0390, padCtl: 0, muxCtlOFS: 0x0104, muxMode: 3, selInputOFS: 0x0600, selInput: 0}
	MX6_PAD_LCD_CLK__EIM_CS2_B            = &pad{padCtlOFS: 0x0390, padCtl: 0, muxCtlOFS: 0x0104, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_CLK__GPIO3_IO00           = &pad{padCtlOFS: 0x0390, padCtl: 0, muxCtlOFS: 0x0104, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_CLK__WDOG1_WDOG_RST_B_DEB = &pad{padCtlOFS: 0x0390, padCtl: 0, muxCtlOFS: 0x0104, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_LCD_ENABLE__LCDIF_ENABLE = &pad{padCtlOFS: 0x0394, padCtl: 0, muxCtlOFS: 0x0108, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_ENABLE__LCDIF_RD_E   = &pad{padCtlOFS: 0x0394, padCtl: 0, muxCtlOFS: 0x0108, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_ENABLE__UART4_DCE_RX = &pad{padCtlOFS: 0x0394, padCtl: 0, muxCtlOFS: 0x0108, muxMode: 2, selInputOFS: 0x063C, selInput: 3}
	MX6_PAD_LCD_ENABLE__UART4_DTE_TX = &pad{padCtlOFS: 0x0394, padCtl: 0, muxCtlOFS: 0x0108, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_ENABLE__SAI3_TX_SYNC = &pad{padCtlOFS: 0x0394, padCtl: 0, muxCtlOFS: 0x0108, muxMode: 3, selInputOFS: 0x060C, selInput: 0}
	MX6_PAD_LCD_ENABLE__EIM_CS3_B    = &pad{padCtlOFS: 0x0394, padCtl: 0, muxCtlOFS: 0x0108, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_ENABLE__GPIO3_IO01   = &pad{padCtlOFS: 0x0394, padCtl: 0, muxCtlOFS: 0x0108, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_ENABLE__ECSPI2_RDY   = &pad{padCtlOFS: 0x0394, padCtl: 0, muxCtlOFS: 0x0108, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_LCD_HSYNC__LCDIF_HSYNC          = &pad{padCtlOFS: 0x0398, padCtl: 0, muxCtlOFS: 0x010C, muxMode: 0, selInputOFS: 0x05DC, selInput: 0}
	MX6_PAD_LCD_HSYNC__LCDIF_RS             = &pad{padCtlOFS: 0x0398, padCtl: 0, muxCtlOFS: 0x010C, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_HSYNC__UART4_DCE_CTS        = &pad{padCtlOFS: 0x0398, padCtl: 0, muxCtlOFS: 0x010C, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_HSYNC__UART4_DTE_RTS        = &pad{padCtlOFS: 0x0398, padCtl: 0, muxCtlOFS: 0x010C, muxMode: 2, selInputOFS: 0x0638, selInput: 2}
	MX6_PAD_LCD_HSYNC__SAI3_TX_BCLK         = &pad{padCtlOFS: 0x0398, padCtl: 0, muxCtlOFS: 0x010C, muxMode: 3, selInputOFS: 0x0608, selInput: 0}
	MX6_PAD_LCD_HSYNC__WDOG3_WDOG_RST_B_DEB = &pad{padCtlOFS: 0x0398, padCtl: 0, muxCtlOFS: 0x010C, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_HSYNC__GPIO3_IO02           = &pad{padCtlOFS: 0x0398, padCtl: 0, muxCtlOFS: 0x010C, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_HSYNC__ECSPI2_SS1           = &pad{padCtlOFS: 0x0398, padCtl: 0, muxCtlOFS: 0x010C, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_LCD_VSYNC__LCDIF_VSYNC   = &pad{padCtlOFS: 0x039C, padCtl: 0, muxCtlOFS: 0x0110, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_VSYNC__LCDIF_BUSY    = &pad{padCtlOFS: 0x039C, padCtl: 0, muxCtlOFS: 0x0110, muxMode: 1, selInputOFS: 0x05DC, selInput: 1}
	MX6_PAD_LCD_VSYNC__UART4_DCE_RTS = &pad{padCtlOFS: 0x039C, padCtl: 0, muxCtlOFS: 0x0110, muxMode: 2, selInputOFS: 0x0638, selInput: 3}
	MX6_PAD_LCD_VSYNC__UART4_DTE_CTS = &pad{padCtlOFS: 0x039C, padCtl: 0, muxCtlOFS: 0x0110, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_VSYNC__SAI3_RX_DATA  = &pad{padCtlOFS: 0x039C, padCtl: 0, muxCtlOFS: 0x0110, muxMode: 3, selInputOFS: 0x0604, selInput: 0}
	MX6_PAD_LCD_VSYNC__WDOG2_WDOG_B  = &pad{padCtlOFS: 0x039C, padCtl: 0, muxCtlOFS: 0x0110, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_VSYNC__GPIO3_IO03    = &pad{padCtlOFS: 0x039C, padCtl: 0, muxCtlOFS: 0x0110, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_VSYNC__ECSPI2_SS2    = &pad{padCtlOFS: 0x039C, padCtl: 0, muxCtlOFS: 0x0110, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_LCD_RESET__LCDIF_RESET      = &pad{padCtlOFS: 0x03A0, padCtl: 0, muxCtlOFS: 0x0114, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_RESET__LCDIF_CS         = &pad{padCtlOFS: 0x03A0, padCtl: 0, muxCtlOFS: 0x0114, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_RESET__CA7_MX6UL_EVENTI = &pad{padCtlOFS: 0x03A0, padCtl: 0, muxCtlOFS: 0x0114, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_RESET__SAI3_TX_DATA     = &pad{padCtlOFS: 0x03A0, padCtl: 0, muxCtlOFS: 0x0114, muxMode: 3, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_RESET__WDOG1_WDOG_ANY   = &pad{padCtlOFS: 0x03A0, padCtl: 0, muxCtlOFS: 0x0114, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_RESET__GPIO3_IO04       = &pad{padCtlOFS: 0x03A0, padCtl: 0, muxCtlOFS: 0x0114, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_RESET__ECSPI2_SS3       = &pad{padCtlOFS: 0x03A0, padCtl: 0, muxCtlOFS: 0x0114, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_LCD_DATA00__LCDIF_DATA00         = &pad{padCtlOFS: 0x03A4, padCtl: 0, muxCtlOFS: 0x0118, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA00__PWM1_OUT             = &pad{padCtlOFS: 0x03A4, padCtl: 0, muxCtlOFS: 0x0118, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA00__ENET1_1588_EVENT2_IN = &pad{padCtlOFS: 0x03A4, padCtl: 0, muxCtlOFS: 0x0118, muxMode: 3, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA00__I2C3_SDA             = &pad{padCtlOFS: 0x03A4, padCtl: 0, muxCtlOFS: 0x0118, muxMode: IOMUX_CONFIG_SION | 4, selInputOFS: 0x05B8, selInput: 2}
	MX6_PAD_LCD_DATA00__GPIO3_IO05           = &pad{padCtlOFS: 0x03A4, padCtl: 0, muxCtlOFS: 0x0118, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA00__SRC_BT_CFG00         = &pad{padCtlOFS: 0x03A4, padCtl: 0, muxCtlOFS: 0x0118, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA00__SAI1_MCLK            = &pad{padCtlOFS: 0x03A4, padCtl: 0, muxCtlOFS: 0x0118, muxMode: 8, selInputOFS: 0x05E0, selInput: 1}

	MX6_PAD_LCD_DATA01__LCDIF_DATA01          = &pad{padCtlOFS: 0x03A8, padCtl: 0, muxCtlOFS: 0x011C, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA01__PWM2_OUT              = &pad{padCtlOFS: 0x03A8, padCtl: 0, muxCtlOFS: 0x011C, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA01__ENET1_1588_EVENT2_OUT = &pad{padCtlOFS: 0x03A8, padCtl: 0, muxCtlOFS: 0x011C, muxMode: 3, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA01__I2C3_SCL              = &pad{padCtlOFS: 0x03A8, padCtl: 0, muxCtlOFS: 0x011C, muxMode: IOMUX_CONFIG_SION | 4, selInputOFS: 0x05B4, selInput: 2}
	MX6_PAD_LCD_DATA01__GPIO3_IO06            = &pad{padCtlOFS: 0x03A8, padCtl: 0, muxCtlOFS: 0x011C, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA01__SRC_BT_CFG01          = &pad{padCtlOFS: 0x03A8, padCtl: 0, muxCtlOFS: 0x011C, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA01__SAI1_TX_SYNC          = &pad{padCtlOFS: 0x03A8, padCtl: 0, muxCtlOFS: 0x011C, muxMode: 8, selInputOFS: 0x05EC, selInput: 0}

	MX6_PAD_LCD_DATA02__LCDIF_DATA02         = &pad{padCtlOFS: 0x03AC, padCtl: 0, muxCtlOFS: 0x0120, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA02__PWM3_OUT             = &pad{padCtlOFS: 0x03AC, padCtl: 0, muxCtlOFS: 0x0120, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA02__ENET1_1588_EVENT3_IN = &pad{padCtlOFS: 0x03AC, padCtl: 0, muxCtlOFS: 0x0120, muxMode: 3, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA02__I2C4_SDA             = &pad{padCtlOFS: 0x03AC, padCtl: 0, muxCtlOFS: 0x0120, muxMode: IOMUX_CONFIG_SION | 4, selInputOFS: 0x05C0, selInput: 2}
	MX6_PAD_LCD_DATA02__GPIO3_IO07           = &pad{padCtlOFS: 0x03AC, padCtl: 0, muxCtlOFS: 0x0120, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA02__SRC_BT_CFG02         = &pad{padCtlOFS: 0x03AC, padCtl: 0, muxCtlOFS: 0x0120, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA02__SAI1_TX_BCLK         = &pad{padCtlOFS: 0x03AC, padCtl: 0, muxCtlOFS: 0x0120, muxMode: 8, selInputOFS: 0x05E8, selInput: 0}

	MX6_PAD_LCD_DATA03__LCDIF_DATA03          = &pad{padCtlOFS: 0x03B0, padCtl: 0, muxCtlOFS: 0x0124, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA03__PWM4_OUT              = &pad{padCtlOFS: 0x03B0, padCtl: 0, muxCtlOFS: 0x0124, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA03__ENET1_1588_EVENT3_OUT = &pad{padCtlOFS: 0x03B0, padCtl: 0, muxCtlOFS: 0x0124, muxMode: 3, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA03__I2C4_SCL              = &pad{padCtlOFS: 0x03B0, padCtl: 0, muxCtlOFS: 0x0124, muxMode: IOMUX_CONFIG_SION | 4, selInputOFS: 0x05BC, selInput: 2}
	MX6_PAD_LCD_DATA03__GPIO3_IO08            = &pad{padCtlOFS: 0x03B0, padCtl: 0, muxCtlOFS: 0x0124, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA03__SRC_BT_CFG03          = &pad{padCtlOFS: 0x03B0, padCtl: 0, muxCtlOFS: 0x0124, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA03__SAI1_RX_DATA          = &pad{padCtlOFS: 0x03B0, padCtl: 0, muxCtlOFS: 0x0124, muxMode: 8, selInputOFS: 0x05E4, selInput: 0}

	MX6_PAD_LCD_DATA04__LCDIF_DATA04         = &pad{padCtlOFS: 0x03B4, padCtl: 0, muxCtlOFS: 0x0128, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA04__UART8_DCE_CTS        = &pad{padCtlOFS: 0x03B4, padCtl: 0, muxCtlOFS: 0x0128, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA04__UART8_DTE_RTS        = &pad{padCtlOFS: 0x03B4, padCtl: 0, muxCtlOFS: 0x0128, muxMode: 1, selInputOFS: 0x0658, selInput: 2}
	MX6_PAD_LCD_DATA04__ENET2_1588_EVENT2_IN = &pad{padCtlOFS: 0x03B4, padCtl: 0, muxCtlOFS: 0x0128, muxMode: 3, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA04__SPDIF_SR_CLK         = &pad{padCtlOFS: 0x03B4, padCtl: 0, muxCtlOFS: 0x0128, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA04__GPIO3_IO09           = &pad{padCtlOFS: 0x03B4, padCtl: 0, muxCtlOFS: 0x0128, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA04__SRC_BT_CFG04         = &pad{padCtlOFS: 0x03B4, padCtl: 0, muxCtlOFS: 0x0128, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA04__SAI1_TX_DATA         = &pad{padCtlOFS: 0x03B4, padCtl: 0, muxCtlOFS: 0x0128, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_LCD_DATA05__LCDIF_DATA05          = &pad{padCtlOFS: 0x03B8, padCtl: 0, muxCtlOFS: 0x012C, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA05__UART8_DCE_RTS         = &pad{padCtlOFS: 0x03B8, padCtl: 0, muxCtlOFS: 0x012C, muxMode: 1, selInputOFS: 0x0658, selInput: 3}
	MX6_PAD_LCD_DATA05__UART8_DTE_CTS         = &pad{padCtlOFS: 0x03B8, padCtl: 0, muxCtlOFS: 0x012C, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA05__ENET2_1588_EVENT2_OUT = &pad{padCtlOFS: 0x03B8, padCtl: 0, muxCtlOFS: 0x012C, muxMode: 3, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA05__SPDIF_OUT             = &pad{padCtlOFS: 0x03B8, padCtl: 0, muxCtlOFS: 0x012C, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA05__GPIO3_IO10            = &pad{padCtlOFS: 0x03B8, padCtl: 0, muxCtlOFS: 0x012C, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA05__SRC_BT_CFG05          = &pad{padCtlOFS: 0x03B8, padCtl: 0, muxCtlOFS: 0x012C, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA05__ECSPI1_SS1            = &pad{padCtlOFS: 0x03B8, padCtl: 0, muxCtlOFS: 0x012C, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_LCD_DATA06__LCDIF_DATA06         = &pad{padCtlOFS: 0x03BC, padCtl: 0, muxCtlOFS: 0x0130, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA06__UART7_DCE_CTS        = &pad{padCtlOFS: 0x03BC, padCtl: 0, muxCtlOFS: 0x0130, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA06__UART7_DTE_RTS        = &pad{padCtlOFS: 0x03BC, padCtl: 0, muxCtlOFS: 0x0130, muxMode: 1, selInputOFS: 0x0650, selInput: 2}
	MX6_PAD_LCD_DATA06__ENET2_1588_EVENT3_IN = &pad{padCtlOFS: 0x03BC, padCtl: 0, muxCtlOFS: 0x0130, muxMode: 3, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA06__SPDIF_LOCK           = &pad{padCtlOFS: 0x03BC, padCtl: 0, muxCtlOFS: 0x0130, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA06__GPIO3_IO11           = &pad{padCtlOFS: 0x03BC, padCtl: 0, muxCtlOFS: 0x0130, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA06__SRC_BT_CFG06         = &pad{padCtlOFS: 0x03BC, padCtl: 0, muxCtlOFS: 0x0130, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA06__ECSPI1_SS2           = &pad{padCtlOFS: 0x03BC, padCtl: 0, muxCtlOFS: 0x0130, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_LCD_DATA07__LCDIF_DATA07          = &pad{padCtlOFS: 0x03C0, padCtl: 0, muxCtlOFS: 0x0134, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA07__UART7_DCE_RTS         = &pad{padCtlOFS: 0x03C0, padCtl: 0, muxCtlOFS: 0x0134, muxMode: 1, selInputOFS: 0x0650, selInput: 3}
	MX6_PAD_LCD_DATA07__UART7_DTE_CTS         = &pad{padCtlOFS: 0x03C0, padCtl: 0, muxCtlOFS: 0x0134, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA07__ENET2_1588_EVENT3_OUT = &pad{padCtlOFS: 0x03C0, padCtl: 0, muxCtlOFS: 0x0134, muxMode: 3, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA07__SPDIF_EXT_CLK         = &pad{padCtlOFS: 0x03C0, padCtl: 0, muxCtlOFS: 0x0134, muxMode: 4, selInputOFS: 0x061C, selInput: 0}
	MX6_PAD_LCD_DATA07__GPIO3_IO12            = &pad{padCtlOFS: 0x03C0, padCtl: 0, muxCtlOFS: 0x0134, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA07__SRC_BT_CFG07          = &pad{padCtlOFS: 0x03C0, padCtl: 0, muxCtlOFS: 0x0134, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA07__ECSPI1_SS3            = &pad{padCtlOFS: 0x03C0, padCtl: 0, muxCtlOFS: 0x0134, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_LCD_DATA08__LCDIF_DATA08 = &pad{padCtlOFS: 0x03C4, padCtl: 0, muxCtlOFS: 0x0138, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA08__SPDIF_IN     = &pad{padCtlOFS: 0x03C4, padCtl: 0, muxCtlOFS: 0x0138, muxMode: 1, selInputOFS: 0x0618, selInput: 2}
	MX6_PAD_LCD_DATA08__CSI_DATA16   = &pad{padCtlOFS: 0x03C4, padCtl: 0, muxCtlOFS: 0x0138, muxMode: 3, selInputOFS: 0x0504, selInput: 1}
	MX6_PAD_LCD_DATA08__EIM_DATA00   = &pad{padCtlOFS: 0x03C4, padCtl: 0, muxCtlOFS: 0x0138, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA08__GPIO3_IO13   = &pad{padCtlOFS: 0x03C4, padCtl: 0, muxCtlOFS: 0x0138, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA08__SRC_BT_CFG08 = &pad{padCtlOFS: 0x03C4, padCtl: 0, muxCtlOFS: 0x0138, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA08__FLEXCAN1_TX  = &pad{padCtlOFS: 0x03C4, padCtl: 0, muxCtlOFS: 0x0138, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_LCD_DATA09__LCDIF_DATA09 = &pad{padCtlOFS: 0x03C8, padCtl: 0, muxCtlOFS: 0x013C, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA09__SAI3_MCLK    = &pad{padCtlOFS: 0x03C8, padCtl: 0, muxCtlOFS: 0x013C, muxMode: 1, selInputOFS: 0x0600, selInput: 1}
	MX6_PAD_LCD_DATA09__CSI_DATA17   = &pad{padCtlOFS: 0x03C8, padCtl: 0, muxCtlOFS: 0x013C, muxMode: 3, selInputOFS: 0x0508, selInput: 1}
	MX6_PAD_LCD_DATA09__EIM_DATA01   = &pad{padCtlOFS: 0x03C8, padCtl: 0, muxCtlOFS: 0x013C, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA09__GPIO3_IO14   = &pad{padCtlOFS: 0x03C8, padCtl: 0, muxCtlOFS: 0x013C, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA09__SRC_BT_CFG09 = &pad{padCtlOFS: 0x03C8, padCtl: 0, muxCtlOFS: 0x013C, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA09__FLEXCAN1_RX  = &pad{padCtlOFS: 0x03C8, padCtl: 0, muxCtlOFS: 0x013C, muxMode: 8, selInputOFS: 0x0584, selInput: 2}

	MX6_PAD_LCD_DATA10__LCDIF_DATA10 = &pad{padCtlOFS: 0x03CC, padCtl: 0, muxCtlOFS: 0x0140, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA10__SAI3_RX_SYNC = &pad{padCtlOFS: 0x03CC, padCtl: 0, muxCtlOFS: 0x0140, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA10__CSI_DATA18   = &pad{padCtlOFS: 0x03CC, padCtl: 0, muxCtlOFS: 0x0140, muxMode: 3, selInputOFS: 0x050C, selInput: 1}
	MX6_PAD_LCD_DATA10__EIM_DATA02   = &pad{padCtlOFS: 0x03CC, padCtl: 0, muxCtlOFS: 0x0140, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA10__GPIO3_IO15   = &pad{padCtlOFS: 0x03CC, padCtl: 0, muxCtlOFS: 0x0140, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA10__SRC_BT_CFG10 = &pad{padCtlOFS: 0x03CC, padCtl: 0, muxCtlOFS: 0x0140, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA10__FLEXCAN2_TX  = &pad{padCtlOFS: 0x03CC, padCtl: 0, muxCtlOFS: 0x0140, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_LCD_DATA11__LCDIF_DATA11 = &pad{padCtlOFS: 0x03D0, padCtl: 0, muxCtlOFS: 0x0144, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA11__SAI3_RX_BCLK = &pad{padCtlOFS: 0x03D0, padCtl: 0, muxCtlOFS: 0x0144, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA11__CSI_DATA19   = &pad{padCtlOFS: 0x03D0, padCtl: 0, muxCtlOFS: 0x0144, muxMode: 3, selInputOFS: 0x0510, selInput: 1}
	MX6_PAD_LCD_DATA11__EIM_DATA03   = &pad{padCtlOFS: 0x03D0, padCtl: 0, muxCtlOFS: 0x0144, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA11__GPIO3_IO16   = &pad{padCtlOFS: 0x03D0, padCtl: 0, muxCtlOFS: 0x0144, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA11__SRC_BT_CFG11 = &pad{padCtlOFS: 0x03D0, padCtl: 0, muxCtlOFS: 0x0144, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA11__FLEXCAN2_RX  = &pad{padCtlOFS: 0x03D0, padCtl: 0, muxCtlOFS: 0x0144, muxMode: 8, selInputOFS: 0x0588, selInput: 2}

	MX6_PAD_LCD_DATA12__LCDIF_DATA12 = &pad{padCtlOFS: 0x03D4, padCtl: 0, muxCtlOFS: 0x0148, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA12__SAI3_TX_SYNC = &pad{padCtlOFS: 0x03D4, padCtl: 0, muxCtlOFS: 0x0148, muxMode: 1, selInputOFS: 0x060C, selInput: 1}
	MX6_PAD_LCD_DATA12__CSI_DATA20   = &pad{padCtlOFS: 0x03D4, padCtl: 0, muxCtlOFS: 0x0148, muxMode: 3, selInputOFS: 0x0514, selInput: 1}
	MX6_PAD_LCD_DATA12__EIM_DATA04   = &pad{padCtlOFS: 0x03D4, padCtl: 0, muxCtlOFS: 0x0148, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA12__GPIO3_IO17   = &pad{padCtlOFS: 0x03D4, padCtl: 0, muxCtlOFS: 0x0148, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA12__SRC_BT_CFG12 = &pad{padCtlOFS: 0x03D4, padCtl: 0, muxCtlOFS: 0x0148, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA12__ECSPI1_RDY   = &pad{padCtlOFS: 0x03D4, padCtl: 0, muxCtlOFS: 0x0148, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_LCD_DATA13__LCDIF_DATA13   = &pad{padCtlOFS: 0x03D8, padCtl: 0, muxCtlOFS: 0x014C, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA13__SAI3_TX_BCLK   = &pad{padCtlOFS: 0x03D8, padCtl: 0, muxCtlOFS: 0x014C, muxMode: 1, selInputOFS: 0x0608, selInput: 1}
	MX6_PAD_LCD_DATA13__CSI_DATA21     = &pad{padCtlOFS: 0x03D8, padCtl: 0, muxCtlOFS: 0x014C, muxMode: 3, selInputOFS: 0x0518, selInput: 1}
	MX6_PAD_LCD_DATA13__EIM_DATA05     = &pad{padCtlOFS: 0x03D8, padCtl: 0, muxCtlOFS: 0x014C, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA13__GPIO3_IO18     = &pad{padCtlOFS: 0x03D8, padCtl: 0, muxCtlOFS: 0x014C, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA13__SRC_BT_CFG13   = &pad{padCtlOFS: 0x03D8, padCtl: 0, muxCtlOFS: 0x014C, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA13__USDHC2_RESET_B = &pad{padCtlOFS: 0x03D8, padCtl: 0, muxCtlOFS: 0x014C, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_LCD_DATA14__LCDIF_DATA14 = &pad{padCtlOFS: 0x03DC, padCtl: 0, muxCtlOFS: 0x0150, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA14__SAI3_RX_DATA = &pad{padCtlOFS: 0x03DC, padCtl: 0, muxCtlOFS: 0x0150, muxMode: 1, selInputOFS: 0x0604, selInput: 1}
	MX6_PAD_LCD_DATA14__CSI_DATA22   = &pad{padCtlOFS: 0x03DC, padCtl: 0, muxCtlOFS: 0x0150, muxMode: 3, selInputOFS: 0x051C, selInput: 1}
	MX6_PAD_LCD_DATA14__EIM_DATA06   = &pad{padCtlOFS: 0x03DC, padCtl: 0, muxCtlOFS: 0x0150, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA14__GPIO3_IO19   = &pad{padCtlOFS: 0x03DC, padCtl: 0, muxCtlOFS: 0x0150, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA14__SRC_BT_CFG14 = &pad{padCtlOFS: 0x03DC, padCtl: 0, muxCtlOFS: 0x0150, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA14__USDHC2_DATA4 = &pad{padCtlOFS: 0x03DC, padCtl: 0, muxCtlOFS: 0x0150, muxMode: 8, selInputOFS: 0x068C, selInput: 0}

	MX6_PAD_LCD_DATA15__LCDIF_DATA15 = &pad{padCtlOFS: 0x03E0, padCtl: 0, muxCtlOFS: 0x0154, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA15__SAI3_TX_DATA = &pad{padCtlOFS: 0x03E0, padCtl: 0, muxCtlOFS: 0x0154, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA15__CSI_DATA23   = &pad{padCtlOFS: 0x03E0, padCtl: 0, muxCtlOFS: 0x0154, muxMode: 3, selInputOFS: 0x0520, selInput: 1}
	MX6_PAD_LCD_DATA15__EIM_DATA07   = &pad{padCtlOFS: 0x03E0, padCtl: 0, muxCtlOFS: 0x0154, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA15__GPIO3_IO20   = &pad{padCtlOFS: 0x03E0, padCtl: 0, muxCtlOFS: 0x0154, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA15__SRC_BT_CFG15 = &pad{padCtlOFS: 0x03E0, padCtl: 0, muxCtlOFS: 0x0154, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA15__USDHC2_DATA5 = &pad{padCtlOFS: 0x03E0, padCtl: 0, muxCtlOFS: 0x0154, muxMode: 8, selInputOFS: 0x0690, selInput: 0}

	MX6_PAD_LCD_DATA16__LCDIF_DATA16 = &pad{padCtlOFS: 0x03E4, padCtl: 0, muxCtlOFS: 0x0158, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA16__UART7_DCE_TX = &pad{padCtlOFS: 0x03E4, padCtl: 0, muxCtlOFS: 0x0158, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA16__UART7_DTE_RX = &pad{padCtlOFS: 0x03E4, padCtl: 0, muxCtlOFS: 0x0158, muxMode: 1, selInputOFS: 0x0654, selInput: 2}
	MX6_PAD_LCD_DATA16__CSI_DATA01   = &pad{padCtlOFS: 0x03E4, padCtl: 0, muxCtlOFS: 0x0158, muxMode: 3, selInputOFS: 0x04D4, selInput: 1}
	MX6_PAD_LCD_DATA16__EIM_DATA08   = &pad{padCtlOFS: 0x03E4, padCtl: 0, muxCtlOFS: 0x0158, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA16__GPIO3_IO21   = &pad{padCtlOFS: 0x03E4, padCtl: 0, muxCtlOFS: 0x0158, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA16__SRC_BT_CFG24 = &pad{padCtlOFS: 0x03E4, padCtl: 0, muxCtlOFS: 0x0158, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA16__USDHC2_DATA6 = &pad{padCtlOFS: 0x03E4, padCtl: 0, muxCtlOFS: 0x0158, muxMode: 8, selInputOFS: 0x0694, selInput: 0}

	MX6_PAD_LCD_DATA17__LCDIF_DATA17 = &pad{padCtlOFS: 0x03E8, padCtl: 0, muxCtlOFS: 0x015C, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA17__UART7_DCE_RX = &pad{padCtlOFS: 0x03E8, padCtl: 0, muxCtlOFS: 0x015C, muxMode: 1, selInputOFS: 0x0654, selInput: 3}
	MX6_PAD_LCD_DATA17__UART7_DTE_TX = &pad{padCtlOFS: 0x03E8, padCtl: 0, muxCtlOFS: 0x015C, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA17__CSI_DATA00   = &pad{padCtlOFS: 0x03E8, padCtl: 0, muxCtlOFS: 0x015C, muxMode: 3, selInputOFS: 0x04D0, selInput: 1}
	MX6_PAD_LCD_DATA17__EIM_DATA09   = &pad{padCtlOFS: 0x03E8, padCtl: 0, muxCtlOFS: 0x015C, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA17__GPIO3_IO22   = &pad{padCtlOFS: 0x03E8, padCtl: 0, muxCtlOFS: 0x015C, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA17__SRC_BT_CFG25 = &pad{padCtlOFS: 0x03E8, padCtl: 0, muxCtlOFS: 0x015C, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA17__USDHC2_DATA7 = &pad{padCtlOFS: 0x03E8, padCtl: 0, muxCtlOFS: 0x015C, muxMode: 8, selInputOFS: 0x0698, selInput: 0}

	MX6_PAD_LCD_DATA18__LCDIF_DATA18     = &pad{padCtlOFS: 0x03EC, padCtl: 0, muxCtlOFS: 0x0160, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA18__PWM5_OUT         = &pad{padCtlOFS: 0x03EC, padCtl: 0, muxCtlOFS: 0x0160, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA18__CA7_MX6UL_EVENTO = &pad{padCtlOFS: 0x03EC, padCtl: 0, muxCtlOFS: 0x0160, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA18__CSI_DATA10       = &pad{padCtlOFS: 0x03EC, padCtl: 0, muxCtlOFS: 0x0160, muxMode: 3, selInputOFS: 0x04EC, selInput: 1}
	MX6_PAD_LCD_DATA18__EIM_DATA10       = &pad{padCtlOFS: 0x03EC, padCtl: 0, muxCtlOFS: 0x0160, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA18__GPIO3_IO23       = &pad{padCtlOFS: 0x03EC, padCtl: 0, muxCtlOFS: 0x0160, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA18__SRC_BT_CFG26     = &pad{padCtlOFS: 0x03EC, padCtl: 0, muxCtlOFS: 0x0160, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA18__USDHC2_CMD       = &pad{padCtlOFS: 0x03EC, padCtl: 0, muxCtlOFS: 0x0160, muxMode: 8, selInputOFS: 0x0678, selInput: 1}
	MX6_PAD_LCD_DATA19__EIM_DATA11       = &pad{padCtlOFS: 0x03F0, padCtl: 0, muxCtlOFS: 0x0164, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA19__GPIO3_IO24       = &pad{padCtlOFS: 0x03F0, padCtl: 0, muxCtlOFS: 0x0164, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA19__SRC_BT_CFG27     = &pad{padCtlOFS: 0x03F0, padCtl: 0, muxCtlOFS: 0x0164, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA19__USDHC2_CLK       = &pad{padCtlOFS: 0x03F0, padCtl: 0, muxCtlOFS: 0x0164, muxMode: 8, selInputOFS: 0x0670, selInput: 1}

	MX6_PAD_LCD_DATA19__LCDIF_DATA19   = &pad{padCtlOFS: 0x03F0, padCtl: 0, muxCtlOFS: 0x0164, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA19__PWM6_OUT       = &pad{padCtlOFS: 0x03F0, padCtl: 0, muxCtlOFS: 0x0164, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA19__WDOG1_WDOG_ANY = &pad{padCtlOFS: 0x03F0, padCtl: 0, muxCtlOFS: 0x0164, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA19__CSI_DATA11     = &pad{padCtlOFS: 0x03F0, padCtl: 0, muxCtlOFS: 0x0164, muxMode: 3, selInputOFS: 0x04F0, selInput: 1}
	MX6_PAD_LCD_DATA20__EIM_DATA12     = &pad{padCtlOFS: 0x03F4, padCtl: 0, muxCtlOFS: 0x0168, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA20__GPIO3_IO25     = &pad{padCtlOFS: 0x03F4, padCtl: 0, muxCtlOFS: 0x0168, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA20__SRC_BT_CFG28   = &pad{padCtlOFS: 0x03F4, padCtl: 0, muxCtlOFS: 0x0168, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA20__USDHC2_DATA0   = &pad{padCtlOFS: 0x03F4, padCtl: 0, muxCtlOFS: 0x0168, muxMode: 8, selInputOFS: 0x067C, selInput: 1}

	MX6_PAD_LCD_DATA20__LCDIF_DATA20 = &pad{padCtlOFS: 0x03F4, padCtl: 0, muxCtlOFS: 0x0168, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA20__UART8_DCE_TX = &pad{padCtlOFS: 0x03F4, padCtl: 0, muxCtlOFS: 0x0168, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA20__UART8_DTE_RX = &pad{padCtlOFS: 0x03F4, padCtl: 0, muxCtlOFS: 0x0168, muxMode: 1, selInputOFS: 0x065C, selInput: 2}
	MX6_PAD_LCD_DATA20__ECSPI1_SCLK  = &pad{padCtlOFS: 0x03F4, padCtl: 0, muxCtlOFS: 0x0168, muxMode: 2, selInputOFS: 0x0534, selInput: 0}
	MX6_PAD_LCD_DATA20__CSI_DATA12   = &pad{padCtlOFS: 0x03F4, padCtl: 0, muxCtlOFS: 0x0168, muxMode: 3, selInputOFS: 0x04F4, selInput: 1}

	MX6_PAD_LCD_DATA21__LCDIF_DATA21 = &pad{padCtlOFS: 0x03F8, padCtl: 0, muxCtlOFS: 0x016C, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA21__UART8_DCE_RX = &pad{padCtlOFS: 0x03F8, padCtl: 0, muxCtlOFS: 0x016C, muxMode: 1, selInputOFS: 0x065C, selInput: 3}
	MX6_PAD_LCD_DATA21__UART8_DTE_TX = &pad{padCtlOFS: 0x03F8, padCtl: 0, muxCtlOFS: 0x016C, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA21__ECSPI1_SS0   = &pad{padCtlOFS: 0x03F8, padCtl: 0, muxCtlOFS: 0x016C, muxMode: 2, selInputOFS: 0x0540, selInput: 0}
	MX6_PAD_LCD_DATA21__CSI_DATA13   = &pad{padCtlOFS: 0x03F8, padCtl: 0, muxCtlOFS: 0x016C, muxMode: 3, selInputOFS: 0x04F8, selInput: 1}
	MX6_PAD_LCD_DATA21__EIM_DATA13   = &pad{padCtlOFS: 0x03F8, padCtl: 0, muxCtlOFS: 0x016C, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA21__GPIO3_IO26   = &pad{padCtlOFS: 0x03F8, padCtl: 0, muxCtlOFS: 0x016C, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA21__SRC_BT_CFG29 = &pad{padCtlOFS: 0x03F8, padCtl: 0, muxCtlOFS: 0x016C, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA21__USDHC2_DATA1 = &pad{padCtlOFS: 0x03F8, padCtl: 0, muxCtlOFS: 0x016C, muxMode: 8, selInputOFS: 0x0680, selInput: 1}

	MX6_PAD_LCD_DATA22__LCDIF_DATA22 = &pad{padCtlOFS: 0x03FC, padCtl: 0, muxCtlOFS: 0x0170, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA22__MQS_RIGHT    = &pad{padCtlOFS: 0x03FC, padCtl: 0, muxCtlOFS: 0x0170, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA22__ECSPI1_MOSI  = &pad{padCtlOFS: 0x03FC, padCtl: 0, muxCtlOFS: 0x0170, muxMode: 2, selInputOFS: 0x053C, selInput: 0}
	MX6_PAD_LCD_DATA22__CSI_DATA14   = &pad{padCtlOFS: 0x03FC, padCtl: 0, muxCtlOFS: 0x0170, muxMode: 3, selInputOFS: 0x04FC, selInput: 1}
	MX6_PAD_LCD_DATA22__EIM_DATA14   = &pad{padCtlOFS: 0x03FC, padCtl: 0, muxCtlOFS: 0x0170, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA22__GPIO3_IO27   = &pad{padCtlOFS: 0x03FC, padCtl: 0, muxCtlOFS: 0x0170, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA22__SRC_BT_CFG30 = &pad{padCtlOFS: 0x03FC, padCtl: 0, muxCtlOFS: 0x0170, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA22__USDHC2_DATA2 = &pad{padCtlOFS: 0x03FC, padCtl: 0, muxCtlOFS: 0x0170, muxMode: 8, selInputOFS: 0x0684, selInput: 0}

	MX6_PAD_LCD_DATA23__LCDIF_DATA23 = &pad{padCtlOFS: 0x0400, padCtl: 0, muxCtlOFS: 0x0174, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA23__MQS_LEFT     = &pad{padCtlOFS: 0x0400, padCtl: 0, muxCtlOFS: 0x0174, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA23__ECSPI1_MISO  = &pad{padCtlOFS: 0x0400, padCtl: 0, muxCtlOFS: 0x0174, muxMode: 2, selInputOFS: 0x0538, selInput: 0}
	MX6_PAD_LCD_DATA23__CSI_DATA15   = &pad{padCtlOFS: 0x0400, padCtl: 0, muxCtlOFS: 0x0174, muxMode: 3, selInputOFS: 0x0500, selInput: 1}
	MX6_PAD_LCD_DATA23__EIM_DATA15   = &pad{padCtlOFS: 0x0400, padCtl: 0, muxCtlOFS: 0x0174, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA23__GPIO3_IO28   = &pad{padCtlOFS: 0x0400, padCtl: 0, muxCtlOFS: 0x0174, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA23__SRC_BT_CFG31 = &pad{padCtlOFS: 0x0400, padCtl: 0, muxCtlOFS: 0x0174, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_LCD_DATA23__USDHC2_DATA3 = &pad{padCtlOFS: 0x0400, padCtl: 0, muxCtlOFS: 0x0174, muxMode: 8, selInputOFS: 0x0688, selInput: 1}

	MX6_PAD_NAND_RE_B__RAWNAND_RE_B = &pad{padCtlOFS: 0x0404, padCtl: 0, muxCtlOFS: 0x0178, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_RE_B__USDHC2_CLK   = &pad{padCtlOFS: 0x0404, padCtl: 0, muxCtlOFS: 0x0178, muxMode: 1, selInputOFS: 0x0670, selInput: 2}
	MX6_PAD_NAND_RE_B__QSPI_B_SCLK  = &pad{padCtlOFS: 0x0404, padCtl: 0, muxCtlOFS: 0x0178, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_RE_B__KPP_ROW00    = &pad{padCtlOFS: 0x0404, padCtl: 0, muxCtlOFS: 0x0178, muxMode: 3, selInputOFS: 0x05D0, selInput: 1}
	MX6_PAD_NAND_RE_B__EIM_EB_B00   = &pad{padCtlOFS: 0x0404, padCtl: 0, muxCtlOFS: 0x0178, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_RE_B__GPIO4_IO00   = &pad{padCtlOFS: 0x0404, padCtl: 0, muxCtlOFS: 0x0178, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_RE_B__ECSPI3_SS2   = &pad{padCtlOFS: 0x0404, padCtl: 0, muxCtlOFS: 0x0178, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_NAND_WE_B__RAWNAND_WE_B = &pad{padCtlOFS: 0x0408, padCtl: 0, muxCtlOFS: 0x017C, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_WE_B__USDHC2_CMD   = &pad{padCtlOFS: 0x0408, padCtl: 0, muxCtlOFS: 0x017C, muxMode: 1, selInputOFS: 0x0678, selInput: 2}
	MX6_PAD_NAND_WE_B__QSPI_B_SS0_B = &pad{padCtlOFS: 0x0408, padCtl: 0, muxCtlOFS: 0x017C, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_WE_B__KPP_COL00    = &pad{padCtlOFS: 0x0408, padCtl: 0, muxCtlOFS: 0x017C, muxMode: 3, selInputOFS: 0x05C4, selInput: 1}
	MX6_PAD_NAND_WE_B__EIM_EB_B01   = &pad{padCtlOFS: 0x0408, padCtl: 0, muxCtlOFS: 0x017C, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_WE_B__GPIO4_IO01   = &pad{padCtlOFS: 0x0408, padCtl: 0, muxCtlOFS: 0x017C, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_WE_B__ECSPI3_SS3   = &pad{padCtlOFS: 0x0408, padCtl: 0, muxCtlOFS: 0x017C, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_NAND_DATA00__RAWNAND_DATA00 = &pad{padCtlOFS: 0x040C, padCtl: 0, muxCtlOFS: 0x0180, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_DATA00__USDHC2_DATA0   = &pad{padCtlOFS: 0x040C, padCtl: 0, muxCtlOFS: 0x0180, muxMode: 1, selInputOFS: 0x067C, selInput: 2}
	MX6_PAD_NAND_DATA00__QSPI_B_SS1_B   = &pad{padCtlOFS: 0x040C, padCtl: 0, muxCtlOFS: 0x0180, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_DATA00__KPP_ROW01      = &pad{padCtlOFS: 0x040C, padCtl: 0, muxCtlOFS: 0x0180, muxMode: 3, selInputOFS: 0x05D4, selInput: 1}
	MX6_PAD_NAND_DATA00__EIM_AD08       = &pad{padCtlOFS: 0x040C, padCtl: 0, muxCtlOFS: 0x0180, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_DATA00__GPIO4_IO02     = &pad{padCtlOFS: 0x040C, padCtl: 0, muxCtlOFS: 0x0180, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_DATA00__ECSPI4_RDY     = &pad{padCtlOFS: 0x040C, padCtl: 0, muxCtlOFS: 0x0180, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_NAND_DATA01__RAWNAND_DATA01 = &pad{padCtlOFS: 0x0410, padCtl: 0, muxCtlOFS: 0x0184, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_DATA01__USDHC2_DATA1   = &pad{padCtlOFS: 0x0410, padCtl: 0, muxCtlOFS: 0x0184, muxMode: 1, selInputOFS: 0x0680, selInput: 2}
	MX6_PAD_NAND_DATA01__QSPI_B_DQS     = &pad{padCtlOFS: 0x0410, padCtl: 0, muxCtlOFS: 0x0184, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_DATA01__KPP_COL01      = &pad{padCtlOFS: 0x0410, padCtl: 0, muxCtlOFS: 0x0184, muxMode: 3, selInputOFS: 0x05C8, selInput: 1}
	MX6_PAD_NAND_DATA01__EIM_AD09       = &pad{padCtlOFS: 0x0410, padCtl: 0, muxCtlOFS: 0x0184, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_DATA01__GPIO4_IO03     = &pad{padCtlOFS: 0x0410, padCtl: 0, muxCtlOFS: 0x0184, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_DATA01__ECSPI4_SS1     = &pad{padCtlOFS: 0x0410, padCtl: 0, muxCtlOFS: 0x0184, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_NAND_DATA02__RAWNAND_DATA02 = &pad{padCtlOFS: 0x0414, padCtl: 0, muxCtlOFS: 0x0188, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_DATA02__USDHC2_DATA2   = &pad{padCtlOFS: 0x0414, padCtl: 0, muxCtlOFS: 0x0188, muxMode: 1, selInputOFS: 0x0684, selInput: 1}
	MX6_PAD_NAND_DATA02__QSPI_B_DATA00  = &pad{padCtlOFS: 0x0414, padCtl: 0, muxCtlOFS: 0x0188, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_DATA02__KPP_ROW02      = &pad{padCtlOFS: 0x0414, padCtl: 0, muxCtlOFS: 0x0188, muxMode: 3, selInputOFS: 0x05D8, selInput: 1}
	MX6_PAD_NAND_DATA02__EIM_AD10       = &pad{padCtlOFS: 0x0414, padCtl: 0, muxCtlOFS: 0x0188, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_DATA02__GPIO4_IO04     = &pad{padCtlOFS: 0x0414, padCtl: 0, muxCtlOFS: 0x0188, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_DATA02__ECSPI4_SS2     = &pad{padCtlOFS: 0x0414, padCtl: 0, muxCtlOFS: 0x0188, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_NAND_DATA03__RAWNAND_DATA03 = &pad{padCtlOFS: 0x0418, padCtl: 0, muxCtlOFS: 0x018C, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_DATA03__USDHC2_DATA3   = &pad{padCtlOFS: 0x0418, padCtl: 0, muxCtlOFS: 0x018C, muxMode: 1, selInputOFS: 0x0688, selInput: 2}
	MX6_PAD_NAND_DATA03__QSPI_B_DATA01  = &pad{padCtlOFS: 0x0418, padCtl: 0, muxCtlOFS: 0x018C, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_DATA03__KPP_COL02      = &pad{padCtlOFS: 0x0418, padCtl: 0, muxCtlOFS: 0x018C, muxMode: 3, selInputOFS: 0x05CC, selInput: 1}
	MX6_PAD_NAND_DATA03__EIM_AD11       = &pad{padCtlOFS: 0x0418, padCtl: 0, muxCtlOFS: 0x018C, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_DATA03__GPIO4_IO05     = &pad{padCtlOFS: 0x0418, padCtl: 0, muxCtlOFS: 0x018C, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_DATA03__ECSPI4_SS3     = &pad{padCtlOFS: 0x0418, padCtl: 0, muxCtlOFS: 0x018C, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_NAND_DATA04__RAWNAND_DATA04 = &pad{padCtlOFS: 0x041C, padCtl: 0, muxCtlOFS: 0x0190, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_DATA04__USDHC2_DATA4   = &pad{padCtlOFS: 0x041C, padCtl: 0, muxCtlOFS: 0x0190, muxMode: 1, selInputOFS: 0x068C, selInput: 1}
	MX6_PAD_NAND_DATA04__QSPI_B_DATA02  = &pad{padCtlOFS: 0x041C, padCtl: 0, muxCtlOFS: 0x0190, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_DATA04__ECSPI4_SCLK    = &pad{padCtlOFS: 0x041C, padCtl: 0, muxCtlOFS: 0x0190, muxMode: 3, selInputOFS: 0x0564, selInput: 1}
	MX6_PAD_NAND_DATA04__EIM_AD12       = &pad{padCtlOFS: 0x041C, padCtl: 0, muxCtlOFS: 0x0190, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_DATA04__GPIO4_IO06     = &pad{padCtlOFS: 0x041C, padCtl: 0, muxCtlOFS: 0x0190, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_DATA04__UART2_DCE_TX   = &pad{padCtlOFS: 0x041C, padCtl: 0, muxCtlOFS: 0x0190, muxMode: 8, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_DATA04__UART2_DTE_RX   = &pad{padCtlOFS: 0x041C, padCtl: 0, muxCtlOFS: 0x0190, muxMode: 8, selInputOFS: 0x062C, selInput: 2}

	MX6_PAD_NAND_DATA05__RAWNAND_DATA05 = &pad{padCtlOFS: 0x0420, padCtl: 0, muxCtlOFS: 0x0194, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_DATA05__USDHC2_DATA5   = &pad{padCtlOFS: 0x0420, padCtl: 0, muxCtlOFS: 0x0194, muxMode: 1, selInputOFS: 0x0690, selInput: 1}
	MX6_PAD_NAND_DATA05__QSPI_B_DATA03  = &pad{padCtlOFS: 0x0420, padCtl: 0, muxCtlOFS: 0x0194, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_DATA05__ECSPI4_MOSI    = &pad{padCtlOFS: 0x0420, padCtl: 0, muxCtlOFS: 0x0194, muxMode: 3, selInputOFS: 0x056C, selInput: 1}
	MX6_PAD_NAND_DATA05__EIM_AD13       = &pad{padCtlOFS: 0x0420, padCtl: 0, muxCtlOFS: 0x0194, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_DATA05__GPIO4_IO07     = &pad{padCtlOFS: 0x0420, padCtl: 0, muxCtlOFS: 0x0194, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_DATA05__UART2_DCE_RX   = &pad{padCtlOFS: 0x0420, padCtl: 0, muxCtlOFS: 0x0194, muxMode: 8, selInputOFS: 0x062C, selInput: 3}
	MX6_PAD_NAND_DATA05__UART2_DTE_TX   = &pad{padCtlOFS: 0x0420, padCtl: 0, muxCtlOFS: 0x0194, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_NAND_DATA06__RAWNAND_DATA06 = &pad{padCtlOFS: 0x0424, padCtl: 0, muxCtlOFS: 0x0198, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_DATA06__USDHC2_DATA6   = &pad{padCtlOFS: 0x0424, padCtl: 0, muxCtlOFS: 0x0198, muxMode: 1, selInputOFS: 0x0694, selInput: 1}
	MX6_PAD_NAND_DATA06__SAI2_RX_BCLK   = &pad{padCtlOFS: 0x0424, padCtl: 0, muxCtlOFS: 0x0198, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_DATA06__ECSPI4_MISO    = &pad{padCtlOFS: 0x0424, padCtl: 0, muxCtlOFS: 0x0198, muxMode: 3, selInputOFS: 0x0568, selInput: 1}
	MX6_PAD_NAND_DATA06__EIM_AD14       = &pad{padCtlOFS: 0x0424, padCtl: 0, muxCtlOFS: 0x0198, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_DATA06__GPIO4_IO08     = &pad{padCtlOFS: 0x0424, padCtl: 0, muxCtlOFS: 0x0198, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_DATA06__UART2_DCE_CTS  = &pad{padCtlOFS: 0x0424, padCtl: 0, muxCtlOFS: 0x0198, muxMode: 8, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_DATA06__UART2_DTE_RTS  = &pad{padCtlOFS: 0x0424, padCtl: 0, muxCtlOFS: 0x0198, muxMode: 8, selInputOFS: 0x0628, selInput: 4}

	MX6_PAD_NAND_DATA07__RAWNAND_DATA07 = &pad{padCtlOFS: 0x0428, padCtl: 0, muxCtlOFS: 0x019C, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_DATA07__USDHC2_DATA7   = &pad{padCtlOFS: 0x0428, padCtl: 0, muxCtlOFS: 0x019C, muxMode: 1, selInputOFS: 0x0698, selInput: 1}
	MX6_PAD_NAND_DATA07__QSPI_A_SS1_B   = &pad{padCtlOFS: 0x0428, padCtl: 0, muxCtlOFS: 0x019C, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_DATA07__ECSPI4_SS0     = &pad{padCtlOFS: 0x0428, padCtl: 0, muxCtlOFS: 0x019C, muxMode: 3, selInputOFS: 0x0570, selInput: 1}
	MX6_PAD_NAND_DATA07__EIM_AD15       = &pad{padCtlOFS: 0x0428, padCtl: 0, muxCtlOFS: 0x019C, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_DATA07__GPIO4_IO09     = &pad{padCtlOFS: 0x0428, padCtl: 0, muxCtlOFS: 0x019C, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_DATA07__UART2_DCE_RTS  = &pad{padCtlOFS: 0x0428, padCtl: 0, muxCtlOFS: 0x019C, muxMode: 8, selInputOFS: 0x0628, selInput: 5}
	MX6_PAD_NAND_DATA07__UART2_DTE_CTS  = &pad{padCtlOFS: 0x0428, padCtl: 0, muxCtlOFS: 0x019C, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_NAND_ALE__RAWNAND_ALE    = &pad{padCtlOFS: 0x042C, padCtl: 0, muxCtlOFS: 0x01A0, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_ALE__USDHC2_RESET_B = &pad{padCtlOFS: 0x042C, padCtl: 0, muxCtlOFS: 0x01A0, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_ALE__QSPI_A_DQS     = &pad{padCtlOFS: 0x042C, padCtl: 0, muxCtlOFS: 0x01A0, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_ALE__PWM3_OUT       = &pad{padCtlOFS: 0x042C, padCtl: 0, muxCtlOFS: 0x01A0, muxMode: 3, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_ALE__EIM_ADDR17     = &pad{padCtlOFS: 0x042C, padCtl: 0, muxCtlOFS: 0x01A0, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_ALE__GPIO4_IO10     = &pad{padCtlOFS: 0x042C, padCtl: 0, muxCtlOFS: 0x01A0, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_ALE__ECSPI3_SS1     = &pad{padCtlOFS: 0x042C, padCtl: 0, muxCtlOFS: 0x01A0, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_NAND_WP_B__RAWNAND_WP_B   = &pad{padCtlOFS: 0x0430, padCtl: 0, muxCtlOFS: 0x01A4, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_WP_B__USDHC1_RESET_B = &pad{padCtlOFS: 0x0430, padCtl: 0, muxCtlOFS: 0x01A4, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_WP_B__QSPI_A_SCLK    = &pad{padCtlOFS: 0x0430, padCtl: 0, muxCtlOFS: 0x01A4, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_WP_B__PWM4_OUT       = &pad{padCtlOFS: 0x0430, padCtl: 0, muxCtlOFS: 0x01A4, muxMode: 3, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_WP_B__EIM_BCLK       = &pad{padCtlOFS: 0x0430, padCtl: 0, muxCtlOFS: 0x01A4, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_WP_B__GPIO4_IO11     = &pad{padCtlOFS: 0x0430, padCtl: 0, muxCtlOFS: 0x01A4, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_WP_B__ECSPI3_RDY     = &pad{padCtlOFS: 0x0430, padCtl: 0, muxCtlOFS: 0x01A4, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_NAND_READY_B__RAWNAND_READY_B = &pad{padCtlOFS: 0x0434, padCtl: 0, muxCtlOFS: 0x01A8, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_READY_B__USDHC1_DATA4    = &pad{padCtlOFS: 0x0434, padCtl: 0, muxCtlOFS: 0x01A8, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_READY_B__QSPI_A_DATA00   = &pad{padCtlOFS: 0x0434, padCtl: 0, muxCtlOFS: 0x01A8, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_READY_B__ECSPI3_SS0      = &pad{padCtlOFS: 0x0434, padCtl: 0, muxCtlOFS: 0x01A8, muxMode: 3, selInputOFS: 0x0560, selInput: 1}
	MX6_PAD_NAND_READY_B__EIM_CS1_B       = &pad{padCtlOFS: 0x0434, padCtl: 0, muxCtlOFS: 0x01A8, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_READY_B__GPIO4_IO12      = &pad{padCtlOFS: 0x0434, padCtl: 0, muxCtlOFS: 0x01A8, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_READY_B__UART3_DCE_TX    = &pad{padCtlOFS: 0x0434, padCtl: 0, muxCtlOFS: 0x01A8, muxMode: 8, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_READY_B__UART3_DTE_RX    = &pad{padCtlOFS: 0x0434, padCtl: 0, muxCtlOFS: 0x01A8, muxMode: 8, selInputOFS: 0x0634, selInput: 2}

	MX6_PAD_NAND_CE0_B__RAWNAND_CE0_B = &pad{padCtlOFS: 0x0438, padCtl: 0, muxCtlOFS: 0x01AC, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_CE0_B__USDHC1_DATA5  = &pad{padCtlOFS: 0x0438, padCtl: 0, muxCtlOFS: 0x01AC, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_CE0_B__QSPI_A_DATA01 = &pad{padCtlOFS: 0x0438, padCtl: 0, muxCtlOFS: 0x01AC, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_CE0_B__ECSPI3_SCLK   = &pad{padCtlOFS: 0x0438, padCtl: 0, muxCtlOFS: 0x01AC, muxMode: 3, selInputOFS: 0x0554, selInput: 1}
	MX6_PAD_NAND_CE0_B__EIM_DTACK_B   = &pad{padCtlOFS: 0x0438, padCtl: 0, muxCtlOFS: 0x01AC, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_CE0_B__GPIO4_IO13    = &pad{padCtlOFS: 0x0438, padCtl: 0, muxCtlOFS: 0x01AC, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_CE0_B__UART3_DCE_RX  = &pad{padCtlOFS: 0x0438, padCtl: 0, muxCtlOFS: 0x01AC, muxMode: 8, selInputOFS: 0x0634, selInput: 3}
	MX6_PAD_NAND_CE0_B__UART3_DTE_TX  = &pad{padCtlOFS: 0x0438, padCtl: 0, muxCtlOFS: 0x01AC, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_NAND_CE1_B__RAWNAND_CE1_B = &pad{padCtlOFS: 0x043C, padCtl: 0, muxCtlOFS: 0x01B0, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_CE1_B__USDHC1_DATA6  = &pad{padCtlOFS: 0x043C, padCtl: 0, muxCtlOFS: 0x01B0, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_CE1_B__QSPI_A_DATA02 = &pad{padCtlOFS: 0x043C, padCtl: 0, muxCtlOFS: 0x01B0, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_CE1_B__ECSPI3_MOSI   = &pad{padCtlOFS: 0x043C, padCtl: 0, muxCtlOFS: 0x01B0, muxMode: 3, selInputOFS: 0x055C, selInput: 1}
	MX6_PAD_NAND_CE1_B__EIM_ADDR18    = &pad{padCtlOFS: 0x043C, padCtl: 0, muxCtlOFS: 0x01B0, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_CE1_B__GPIO4_IO14    = &pad{padCtlOFS: 0x043C, padCtl: 0, muxCtlOFS: 0x01B0, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_CE1_B__UART3_DCE_CTS = &pad{padCtlOFS: 0x043C, padCtl: 0, muxCtlOFS: 0x01B0, muxMode: 8, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_CE1_B__UART3_DTE_RTS = &pad{padCtlOFS: 0x043C, padCtl: 0, muxCtlOFS: 0x01B0, muxMode: 8, selInputOFS: 0x0630, selInput: 2}

	MX6_PAD_NAND_CLE__RAWNAND_CLE   = &pad{padCtlOFS: 0x0440, padCtl: 0, muxCtlOFS: 0x01B4, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_CLE__USDHC1_DATA7  = &pad{padCtlOFS: 0x0440, padCtl: 0, muxCtlOFS: 0x01B4, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_CLE__QSPI_A_DATA03 = &pad{padCtlOFS: 0x0440, padCtl: 0, muxCtlOFS: 0x01B4, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_CLE__ECSPI3_MISO   = &pad{padCtlOFS: 0x0440, padCtl: 0, muxCtlOFS: 0x01B4, muxMode: 3, selInputOFS: 0x0558, selInput: 1}
	MX6_PAD_NAND_CLE__EIM_ADDR16    = &pad{padCtlOFS: 0x0440, padCtl: 0, muxCtlOFS: 0x01B4, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_CLE__GPIO4_IO15    = &pad{padCtlOFS: 0x0440, padCtl: 0, muxCtlOFS: 0x01B4, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_CLE__UART3_DCE_RTS = &pad{padCtlOFS: 0x0440, padCtl: 0, muxCtlOFS: 0x01B4, muxMode: 8, selInputOFS: 0x0630, selInput: 3}
	MX6_PAD_NAND_CLE__UART3_DTE_CTS = &pad{padCtlOFS: 0x0440, padCtl: 0, muxCtlOFS: 0x01B4, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_NAND_DQS__RAWNAND_DQS      = &pad{padCtlOFS: 0x0444, padCtl: 0, muxCtlOFS: 0x01B8, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_DQS__CSI_FIELD        = &pad{padCtlOFS: 0x0444, padCtl: 0, muxCtlOFS: 0x01B8, muxMode: 1, selInputOFS: 0x0530, selInput: 1}
	MX6_PAD_NAND_DQS__QSPI_A_SS0_B     = &pad{padCtlOFS: 0x0444, padCtl: 0, muxCtlOFS: 0x01B8, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_DQS__PWM5_OUT         = &pad{padCtlOFS: 0x0444, padCtl: 0, muxCtlOFS: 0x01B8, muxMode: 3, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_DQS__EIM_WAIT         = &pad{padCtlOFS: 0x0444, padCtl: 0, muxCtlOFS: 0x01B8, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_DQS__GPIO4_IO16       = &pad{padCtlOFS: 0x0444, padCtl: 0, muxCtlOFS: 0x01B8, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_NAND_DQS__SDMA_EXT_EVENT01 = &pad{padCtlOFS: 0x0444, padCtl: 0, muxCtlOFS: 0x01B8, muxMode: 6, selInputOFS: 0x0614, selInput: 1}
	MX6_PAD_NAND_DQS__SPDIF_EXT_CLK    = &pad{padCtlOFS: 0x0444, padCtl: 0, muxCtlOFS: 0x01B8, muxMode: 8, selInputOFS: 0x061C, selInput: 1}

	MX6_PAD_SD1_CMD__USDHC1_CMD       = &pad{padCtlOFS: 0x0448, padCtl: 0, muxCtlOFS: 0x01BC, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_SD1_CMD__GPT2_COMPARE1    = &pad{padCtlOFS: 0x0448, padCtl: 0, muxCtlOFS: 0x01BC, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_SD1_CMD__SAI2_RX_SYNC     = &pad{padCtlOFS: 0x0448, padCtl: 0, muxCtlOFS: 0x01BC, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_SD1_CMD__SPDIF_OUT        = &pad{padCtlOFS: 0x0448, padCtl: 0, muxCtlOFS: 0x01BC, muxMode: 3, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_SD1_CMD__EIM_ADDR19       = &pad{padCtlOFS: 0x0448, padCtl: 0, muxCtlOFS: 0x01BC, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_SD1_CMD__GPIO2_IO16       = &pad{padCtlOFS: 0x0448, padCtl: 0, muxCtlOFS: 0x01BC, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_SD1_CMD__SDMA_EXT_EVENT00 = &pad{padCtlOFS: 0x0448, padCtl: 0, muxCtlOFS: 0x01BC, muxMode: 6, selInputOFS: 0x0610, selInput: 2}
	MX6_PAD_SD1_CMD__USB_OTG1_PWR     = &pad{padCtlOFS: 0x0448, padCtl: 0, muxCtlOFS: 0x01BC, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_SD1_CLK__USDHC1_CLK    = &pad{padCtlOFS: 0x044C, padCtl: 0, muxCtlOFS: 0x01C0, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_SD1_CLK__GPT2_COMPARE2 = &pad{padCtlOFS: 0x044C, padCtl: 0, muxCtlOFS: 0x01C0, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_SD1_CLK__SAI2_MCLK     = &pad{padCtlOFS: 0x044C, padCtl: 0, muxCtlOFS: 0x01C0, muxMode: 2, selInputOFS: 0x05F0, selInput: 1}
	MX6_PAD_SD1_CLK__SPDIF_IN      = &pad{padCtlOFS: 0x044C, padCtl: 0, muxCtlOFS: 0x01C0, muxMode: 3, selInputOFS: 0x0618, selInput: 3}
	MX6_PAD_SD1_CLK__EIM_ADDR20    = &pad{padCtlOFS: 0x044C, padCtl: 0, muxCtlOFS: 0x01C0, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_SD1_CLK__GPIO2_IO17    = &pad{padCtlOFS: 0x044C, padCtl: 0, muxCtlOFS: 0x01C0, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_SD1_CLK__USB_OTG1_OC   = &pad{padCtlOFS: 0x044C, padCtl: 0, muxCtlOFS: 0x01C0, muxMode: 8, selInputOFS: 0x0664, selInput: 2}

	MX6_PAD_SD1_DATA0__USDHC1_DATA0   = &pad{padCtlOFS: 0x0450, padCtl: 0, muxCtlOFS: 0x01C4, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_SD1_DATA0__GPT2_COMPARE3  = &pad{padCtlOFS: 0x0450, padCtl: 0, muxCtlOFS: 0x01C4, muxMode: 1, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_SD1_DATA0__SAI2_TX_SYNC   = &pad{padCtlOFS: 0x0450, padCtl: 0, muxCtlOFS: 0x01C4, muxMode: 2, selInputOFS: 0x05FC, selInput: 1}
	MX6_PAD_SD1_DATA0__FLEXCAN1_TX    = &pad{padCtlOFS: 0x0450, padCtl: 0, muxCtlOFS: 0x01C4, muxMode: 3, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_SD1_DATA0__EIM_ADDR21     = &pad{padCtlOFS: 0x0450, padCtl: 0, muxCtlOFS: 0x01C4, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_SD1_DATA0__GPIO2_IO18     = &pad{padCtlOFS: 0x0450, padCtl: 0, muxCtlOFS: 0x01C4, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_SD1_DATA0__ANATOP_OTG1_ID = &pad{padCtlOFS: 0x0450, padCtl: 0, muxCtlOFS: 0x01C4, muxMode: 8, selInputOFS: 0x04B8, selInput: 2}

	MX6_PAD_SD1_DATA1__USDHC1_DATA1 = &pad{padCtlOFS: 0x0454, padCtl: 0, muxCtlOFS: 0x01C8, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_SD1_DATA1__GPT2_CLK     = &pad{padCtlOFS: 0x0454, padCtl: 0, muxCtlOFS: 0x01C8, muxMode: 1, selInputOFS: 0x05A0, selInput: 1}
	MX6_PAD_SD1_DATA1__SAI2_TX_BCLK = &pad{padCtlOFS: 0x0454, padCtl: 0, muxCtlOFS: 0x01C8, muxMode: 2, selInputOFS: 0x05F8, selInput: 1}
	MX6_PAD_SD1_DATA1__FLEXCAN1_RX  = &pad{padCtlOFS: 0x0454, padCtl: 0, muxCtlOFS: 0x01C8, muxMode: 3, selInputOFS: 0x0584, selInput: 3}
	MX6_PAD_SD1_DATA1__EIM_ADDR22   = &pad{padCtlOFS: 0x0454, padCtl: 0, muxCtlOFS: 0x01C8, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_SD1_DATA1__GPIO2_IO19   = &pad{padCtlOFS: 0x0454, padCtl: 0, muxCtlOFS: 0x01C8, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_SD1_DATA1__USB_OTG2_PWR = &pad{padCtlOFS: 0x0454, padCtl: 0, muxCtlOFS: 0x01C8, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_SD1_DATA2__USDHC1_DATA2  = &pad{padCtlOFS: 0x0458, padCtl: 0, muxCtlOFS: 0x01CC, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_SD1_DATA2__GPT2_CAPTURE1 = &pad{padCtlOFS: 0x0458, padCtl: 0, muxCtlOFS: 0x01CC, muxMode: 1, selInputOFS: 0x0598, selInput: 1}
	MX6_PAD_SD1_DATA2__SAI2_RX_DATA  = &pad{padCtlOFS: 0x0458, padCtl: 0, muxCtlOFS: 0x01CC, muxMode: 2, selInputOFS: 0x05F4, selInput: 1}
	MX6_PAD_SD1_DATA2__FLEXCAN2_TX   = &pad{padCtlOFS: 0x0458, padCtl: 0, muxCtlOFS: 0x01CC, muxMode: 3, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_SD1_DATA2__EIM_ADDR23    = &pad{padCtlOFS: 0x0458, padCtl: 0, muxCtlOFS: 0x01CC, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_SD1_DATA2__GPIO2_IO20    = &pad{padCtlOFS: 0x0458, padCtl: 0, muxCtlOFS: 0x01CC, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_SD1_DATA2__CCM_CLKO1     = &pad{padCtlOFS: 0x0458, padCtl: 0, muxCtlOFS: 0x01CC, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_SD1_DATA2__USB_OTG2_OC   = &pad{padCtlOFS: 0x0458, padCtl: 0, muxCtlOFS: 0x01CC, muxMode: 8, selInputOFS: 0x0660, selInput: 2}

	MX6_PAD_SD1_DATA3__USDHC1_DATA3   = &pad{padCtlOFS: 0x045C, padCtl: 0, muxCtlOFS: 0x01D0, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_SD1_DATA3__GPT2_CAPTURE2  = &pad{padCtlOFS: 0x045C, padCtl: 0, muxCtlOFS: 0x01D0, muxMode: 1, selInputOFS: 0x059C, selInput: 1}
	MX6_PAD_SD1_DATA3__SAI2_TX_DATA   = &pad{padCtlOFS: 0x045C, padCtl: 0, muxCtlOFS: 0x01D0, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_SD1_DATA3__FLEXCAN2_RX    = &pad{padCtlOFS: 0x045C, padCtl: 0, muxCtlOFS: 0x01D0, muxMode: 3, selInputOFS: 0x0588, selInput: 3}
	MX6_PAD_SD1_DATA3__EIM_ADDR24     = &pad{padCtlOFS: 0x045C, padCtl: 0, muxCtlOFS: 0x01D0, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_SD1_DATA3__GPIO2_IO21     = &pad{padCtlOFS: 0x045C, padCtl: 0, muxCtlOFS: 0x01D0, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_SD1_DATA3__CCM_CLKO2      = &pad{padCtlOFS: 0x045C, padCtl: 0, muxCtlOFS: 0x01D0, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_SD1_DATA3__ANATOP_OTG2_ID = &pad{padCtlOFS: 0x045C, padCtl: 0, muxCtlOFS: 0x01D0, muxMode: 8, selInputOFS: 0x04BC, selInput: 2}

	MX6_PAD_CSI_MCLK__CSI_MCLK          = &pad{padCtlOFS: 0x0460, padCtl: 0, muxCtlOFS: 0x01D4, muxMode: 0, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_CSI_MCLK__USDHC2_CD_B       = &pad{padCtlOFS: 0x0460, padCtl: 0, muxCtlOFS: 0x01D4, muxMode: 1, selInputOFS: 0x0674, selInput: 0}
	MX6_PAD_CSI_MCLK__RAWNAND_CE2_B     = &pad{padCtlOFS: 0x0460, padCtl: 0, muxCtlOFS: 0x01D4, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_CSI_MCLK__I2C1_SDA          = &pad{padCtlOFS: 0x0460, padCtl: 0, muxCtlOFS: 0x01D4, muxMode: IOMUX_CONFIG_SION | 3, selInputOFS: 0x05A8, selInput: 0}
	MX6_PAD_CSI_MCLK__EIM_CS0_B         = &pad{padCtlOFS: 0x0460, padCtl: 0, muxCtlOFS: 0x01D4, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_CSI_MCLK__GPIO4_IO17        = &pad{padCtlOFS: 0x0460, padCtl: 0, muxCtlOFS: 0x01D4, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_CSI_MCLK__SNVS_HP_VIO_5_CTL = &pad{padCtlOFS: 0x0460, padCtl: 0, muxCtlOFS: 0x01D4, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_CSI_MCLK__UART6_DCE_TX      = &pad{padCtlOFS: 0x0460, padCtl: 0, muxCtlOFS: 0x01D4, muxMode: 8, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_CSI_MCLK__UART6_DTE_RX      = &pad{padCtlOFS: 0x0460, padCtl: 0, muxCtlOFS: 0x01D4, muxMode: 8, selInputOFS: 0x064C, selInput: 0}

	MX6_PAD_CSI_PIXCLK__CSI_PIXCLK    = &pad{padCtlOFS: 0x0464, padCtl: 0, muxCtlOFS: 0x01D8, muxMode: 0, selInputOFS: 0x0528, selInput: 1}
	MX6_PAD_CSI_PIXCLK__USDHC2_WP     = &pad{padCtlOFS: 0x0464, padCtl: 0, muxCtlOFS: 0x01D8, muxMode: 1, selInputOFS: 0x069C, selInput: 2}
	MX6_PAD_CSI_PIXCLK__RAWNAND_CE3_B = &pad{padCtlOFS: 0x0464, padCtl: 0, muxCtlOFS: 0x01D8, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_CSI_PIXCLK__I2C1_SCL      = &pad{padCtlOFS: 0x0464, padCtl: 0, muxCtlOFS: 0x01D8, muxMode: IOMUX_CONFIG_SION | 3, selInputOFS: 0x05A4, selInput: 2}
	MX6_PAD_CSI_PIXCLK__EIM_OE        = &pad{padCtlOFS: 0x0464, padCtl: 0, muxCtlOFS: 0x01D8, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_CSI_PIXCLK__GPIO4_IO18    = &pad{padCtlOFS: 0x0464, padCtl: 0, muxCtlOFS: 0x01D8, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_CSI_PIXCLK__SNVS_HP_VIO_5 = &pad{padCtlOFS: 0x0464, padCtl: 0, muxCtlOFS: 0x01D8, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_CSI_PIXCLK__UART6_DCE_RX  = &pad{padCtlOFS: 0x0464, padCtl: 0, muxCtlOFS: 0x01D8, muxMode: 8, selInputOFS: 0x064C, selInput: 3}
	MX6_PAD_CSI_PIXCLK__UART6_DTE_TX  = &pad{padCtlOFS: 0x0464, padCtl: 0, muxCtlOFS: 0x01D8, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_CSI_VSYNC__CSI_VSYNC      = &pad{padCtlOFS: 0x0468, padCtl: 0, muxCtlOFS: 0x01DC, muxMode: 0, selInputOFS: 0x052C, selInput: 0}
	MX6_PAD_CSI_VSYNC__USDHC2_CLK     = &pad{padCtlOFS: 0x0468, padCtl: 0, muxCtlOFS: 0x01DC, muxMode: 1, selInputOFS: 0x0670, selInput: 0}
	MX6_PAD_CSI_VSYNC__SIM1_PORT1_CLK = &pad{padCtlOFS: 0x0468, padCtl: 0, muxCtlOFS: 0x01DC, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_CSI_VSYNC__I2C2_SDA       = &pad{padCtlOFS: 0x0468, padCtl: 0, muxCtlOFS: 0x01DC, muxMode: IOMUX_CONFIG_SION | 3, selInputOFS: 0x05B0, selInput: 0}
	MX6_PAD_CSI_VSYNC__EIM_RW         = &pad{padCtlOFS: 0x0468, padCtl: 0, muxCtlOFS: 0x01DC, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_CSI_VSYNC__GPIO4_IO19     = &pad{padCtlOFS: 0x0468, padCtl: 0, muxCtlOFS: 0x01DC, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_CSI_VSYNC__PWM7_OUT       = &pad{padCtlOFS: 0x0468, padCtl: 0, muxCtlOFS: 0x01DC, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_CSI_VSYNC__UART6_DCE_RTS  = &pad{padCtlOFS: 0x0468, padCtl: 0, muxCtlOFS: 0x01DC, muxMode: 8, selInputOFS: 0x0648, selInput: 0}
	MX6_PAD_CSI_VSYNC__UART6_DTE_CTS  = &pad{padCtlOFS: 0x0468, padCtl: 0, muxCtlOFS: 0x01DC, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_CSI_HSYNC__CSI_HSYNC     = &pad{padCtlOFS: 0x046C, padCtl: 0, muxCtlOFS: 0x01E0, muxMode: 0, selInputOFS: 0x0524, selInput: 0}
	MX6_PAD_CSI_HSYNC__USDHC2_CMD    = &pad{padCtlOFS: 0x046C, padCtl: 0, muxCtlOFS: 0x01E0, muxMode: 1, selInputOFS: 0x0678, selInput: 0}
	MX6_PAD_CSI_HSYNC__SIM1_PORT1_PD = &pad{padCtlOFS: 0x046C, padCtl: 0, muxCtlOFS: 0x01E0, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_CSI_HSYNC__I2C2_SCL      = &pad{padCtlOFS: 0x046C, padCtl: 0, muxCtlOFS: 0x01E0, muxMode: IOMUX_CONFIG_SION | 3, selInputOFS: 0x05AC, selInput: 0}
	MX6_PAD_CSI_HSYNC__EIM_LBA_B     = &pad{padCtlOFS: 0x046C, padCtl: 0, muxCtlOFS: 0x01E0, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_CSI_HSYNC__GPIO4_IO20    = &pad{padCtlOFS: 0x046C, padCtl: 0, muxCtlOFS: 0x01E0, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_CSI_HSYNC__PWM8_OUT      = &pad{padCtlOFS: 0x046C, padCtl: 0, muxCtlOFS: 0x01E0, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_CSI_HSYNC__UART6_DCE_CTS = &pad{padCtlOFS: 0x046C, padCtl: 0, muxCtlOFS: 0x01E0, muxMode: 8, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_CSI_HSYNC__UART6_DTE_RTS = &pad{padCtlOFS: 0x046C, padCtl: 0, muxCtlOFS: 0x01E0, muxMode: 8, selInputOFS: 0x0648, selInput: 1}

	MX6_PAD_CSI_DATA00__CSI_DATA02       = &pad{padCtlOFS: 0x0470, padCtl: 0, muxCtlOFS: 0x01E4, muxMode: 0, selInputOFS: 0x04C4, selInput: 0}
	MX6_PAD_CSI_DATA00__USDHC2_DATA0     = &pad{padCtlOFS: 0x0470, padCtl: 0, muxCtlOFS: 0x01E4, muxMode: 1, selInputOFS: 0x067C, selInput: 0}
	MX6_PAD_CSI_DATA00__SIM1_PORT1_RST_B = &pad{padCtlOFS: 0x0470, padCtl: 0, muxCtlOFS: 0x01E4, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_CSI_DATA00__ECSPI2_SCLK      = &pad{padCtlOFS: 0x0470, padCtl: 0, muxCtlOFS: 0x01E4, muxMode: 3, selInputOFS: 0x0544, selInput: 0}
	MX6_PAD_CSI_DATA00__EIM_AD00         = &pad{padCtlOFS: 0x0470, padCtl: 0, muxCtlOFS: 0x01E4, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_CSI_DATA00__GPIO4_IO21       = &pad{padCtlOFS: 0x0470, padCtl: 0, muxCtlOFS: 0x01E4, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_CSI_DATA00__SRC_INT_BOOT     = &pad{padCtlOFS: 0x0470, padCtl: 0, muxCtlOFS: 0x01E4, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_CSI_DATA00__UART5_DCE_TX     = &pad{padCtlOFS: 0x0470, padCtl: 0, muxCtlOFS: 0x01E4, muxMode: 8, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_CSI_DATA00__UART5_DTE_RX     = &pad{padCtlOFS: 0x0470, padCtl: 0, muxCtlOFS: 0x01E4, muxMode: 8, selInputOFS: 0x0644, selInput: 0}

	MX6_PAD_CSI_DATA01__CSI_DATA03      = &pad{padCtlOFS: 0x0474, padCtl: 0, muxCtlOFS: 0x01E8, muxMode: 0, selInputOFS: 0x04C8, selInput: 0}
	MX6_PAD_CSI_DATA01__USDHC2_DATA1    = &pad{padCtlOFS: 0x0474, padCtl: 0, muxCtlOFS: 0x01E8, muxMode: 1, selInputOFS: 0x0680, selInput: 0}
	MX6_PAD_CSI_DATA01__SIM1_PORT1_SVEN = &pad{padCtlOFS: 0x0474, padCtl: 0, muxCtlOFS: 0x01E8, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_CSI_DATA01__ECSPI2_SS0      = &pad{padCtlOFS: 0x0474, padCtl: 0, muxCtlOFS: 0x01E8, muxMode: 3, selInputOFS: 0x0550, selInput: 0}
	MX6_PAD_CSI_DATA01__EIM_AD01        = &pad{padCtlOFS: 0x0474, padCtl: 0, muxCtlOFS: 0x01E8, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_CSI_DATA01__GPIO4_IO22      = &pad{padCtlOFS: 0x0474, padCtl: 0, muxCtlOFS: 0x01E8, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_CSI_DATA01__SAI1_MCLK       = &pad{padCtlOFS: 0x0474, padCtl: 0, muxCtlOFS: 0x01E8, muxMode: 6, selInputOFS: 0x05E0, selInput: 0}
	MX6_PAD_CSI_DATA01__UART5_DCE_RX    = &pad{padCtlOFS: 0x0474, padCtl: 0, muxCtlOFS: 0x01E8, muxMode: 8, selInputOFS: 0x0644, selInput: 1}
	MX6_PAD_CSI_DATA01__UART5_DTE_TX    = &pad{padCtlOFS: 0x0474, padCtl: 0, muxCtlOFS: 0x01E8, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_CSI_DATA02__CSI_DATA04      = &pad{padCtlOFS: 0x0478, padCtl: 0, muxCtlOFS: 0x01EC, muxMode: 0, selInputOFS: 0x04D8, selInput: 1}
	MX6_PAD_CSI_DATA02__USDHC2_DATA2    = &pad{padCtlOFS: 0x0478, padCtl: 0, muxCtlOFS: 0x01EC, muxMode: 1, selInputOFS: 0x0684, selInput: 2}
	MX6_PAD_CSI_DATA02__SIM1_PORT1_TRXD = &pad{padCtlOFS: 0x0478, padCtl: 0, muxCtlOFS: 0x01EC, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_CSI_DATA02__ECSPI2_MOSI     = &pad{padCtlOFS: 0x0478, padCtl: 0, muxCtlOFS: 0x01EC, muxMode: 3, selInputOFS: 0x054C, selInput: 1}
	MX6_PAD_CSI_DATA02__EIM_AD02        = &pad{padCtlOFS: 0x0478, padCtl: 0, muxCtlOFS: 0x01EC, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_CSI_DATA02__GPIO4_IO23      = &pad{padCtlOFS: 0x0478, padCtl: 0, muxCtlOFS: 0x01EC, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_CSI_DATA02__SAI1_RX_SYNC    = &pad{padCtlOFS: 0x0478, padCtl: 0, muxCtlOFS: 0x01EC, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_CSI_DATA02__UART5_DCE_RTS   = &pad{padCtlOFS: 0x0478, padCtl: 0, muxCtlOFS: 0x01EC, muxMode: 8, selInputOFS: 0x0640, selInput: 5}
	MX6_PAD_CSI_DATA02__UART5_DTE_CTS   = &pad{padCtlOFS: 0x0478, padCtl: 0, muxCtlOFS: 0x01EC, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_CSI_DATA03__CSI_DATA05    = &pad{padCtlOFS: 0x047C, padCtl: 0, muxCtlOFS: 0x01F0, muxMode: 0, selInputOFS: 0x04CC, selInput: 0}
	MX6_PAD_CSI_DATA03__USDHC2_DATA3  = &pad{padCtlOFS: 0x047C, padCtl: 0, muxCtlOFS: 0x01F0, muxMode: 1, selInputOFS: 0x0688, selInput: 0}
	MX6_PAD_CSI_DATA03__SIM2_PORT1_PD = &pad{padCtlOFS: 0x047C, padCtl: 0, muxCtlOFS: 0x01F0, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_CSI_DATA03__ECSPI2_MISO   = &pad{padCtlOFS: 0x047C, padCtl: 0, muxCtlOFS: 0x01F0, muxMode: 3, selInputOFS: 0x0548, selInput: 0}
	MX6_PAD_CSI_DATA03__EIM_AD03      = &pad{padCtlOFS: 0x047C, padCtl: 0, muxCtlOFS: 0x01F0, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_CSI_DATA03__GPIO4_IO24    = &pad{padCtlOFS: 0x047C, padCtl: 0, muxCtlOFS: 0x01F0, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_CSI_DATA03__SAI1_RX_BCLK  = &pad{padCtlOFS: 0x047C, padCtl: 0, muxCtlOFS: 0x01F0, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_CSI_DATA03__UART5_DCE_CTS = &pad{padCtlOFS: 0x047C, padCtl: 0, muxCtlOFS: 0x01F0, muxMode: 8, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_CSI_DATA03__UART5_DTE_RTS = &pad{padCtlOFS: 0x047C, padCtl: 0, muxCtlOFS: 0x01F0, muxMode: 8, selInputOFS: 0x0640, selInput: 0}

	MX6_PAD_CSI_DATA04__CSI_DATA06     = &pad{padCtlOFS: 0x0480, padCtl: 0, muxCtlOFS: 0x01F4, muxMode: 0, selInputOFS: 0x04DC, selInput: 1}
	MX6_PAD_CSI_DATA04__USDHC2_DATA4   = &pad{padCtlOFS: 0x0480, padCtl: 0, muxCtlOFS: 0x01F4, muxMode: 1, selInputOFS: 0x068C, selInput: 2}
	MX6_PAD_CSI_DATA04__SIM2_PORT1_CLK = &pad{padCtlOFS: 0x0480, padCtl: 0, muxCtlOFS: 0x01F4, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_CSI_DATA04__ECSPI1_SCLK    = &pad{padCtlOFS: 0x0480, padCtl: 0, muxCtlOFS: 0x01F4, muxMode: 3, selInputOFS: 0x0534, selInput: 1}
	MX6_PAD_CSI_DATA04__EIM_AD04       = &pad{padCtlOFS: 0x0480, padCtl: 0, muxCtlOFS: 0x01F4, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_CSI_DATA04__GPIO4_IO25     = &pad{padCtlOFS: 0x0480, padCtl: 0, muxCtlOFS: 0x01F4, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_CSI_DATA04__SAI1_TX_SYNC   = &pad{padCtlOFS: 0x0480, padCtl: 0, muxCtlOFS: 0x01F4, muxMode: 6, selInputOFS: 0x05EC, selInput: 1}
	MX6_PAD_CSI_DATA04__USDHC1_WP      = &pad{padCtlOFS: 0x0480, padCtl: 0, muxCtlOFS: 0x01F4, muxMode: 8, selInputOFS: 0x066C, selInput: 2}

	MX6_PAD_CSI_DATA05__CSI_DATA07       = &pad{padCtlOFS: 0x0484, padCtl: 0, muxCtlOFS: 0x01F8, muxMode: 0, selInputOFS: 0x04E0, selInput: 1}
	MX6_PAD_CSI_DATA05__USDHC2_DATA5     = &pad{padCtlOFS: 0x0484, padCtl: 0, muxCtlOFS: 0x01F8, muxMode: 1, selInputOFS: 0x0690, selInput: 2}
	MX6_PAD_CSI_DATA05__SIM2_PORT1_RST_B = &pad{padCtlOFS: 0x0484, padCtl: 0, muxCtlOFS: 0x01F8, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_CSI_DATA05__ECSPI1_SS0       = &pad{padCtlOFS: 0x0484, padCtl: 0, muxCtlOFS: 0x01F8, muxMode: 3, selInputOFS: 0x0540, selInput: 1}
	MX6_PAD_CSI_DATA05__EIM_AD05         = &pad{padCtlOFS: 0x0484, padCtl: 0, muxCtlOFS: 0x01F8, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_CSI_DATA05__GPIO4_IO26       = &pad{padCtlOFS: 0x0484, padCtl: 0, muxCtlOFS: 0x01F8, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_CSI_DATA05__SAI1_TX_BCLK     = &pad{padCtlOFS: 0x0484, padCtl: 0, muxCtlOFS: 0x01F8, muxMode: 6, selInputOFS: 0x05E8, selInput: 1}
	MX6_PAD_CSI_DATA05__USDHC1_CD_B      = &pad{padCtlOFS: 0x0484, padCtl: 0, muxCtlOFS: 0x01F8, muxMode: 8, selInputOFS: 0x0668, selInput: 2}

	MX6_PAD_CSI_DATA06__CSI_DATA08      = &pad{padCtlOFS: 0x0488, padCtl: 0, muxCtlOFS: 0x01FC, muxMode: 0, selInputOFS: 0x04E4, selInput: 1}
	MX6_PAD_CSI_DATA06__USDHC2_DATA6    = &pad{padCtlOFS: 0x0488, padCtl: 0, muxCtlOFS: 0x01FC, muxMode: 1, selInputOFS: 0x0694, selInput: 2}
	MX6_PAD_CSI_DATA06__SIM2_PORT1_SVEN = &pad{padCtlOFS: 0x0488, padCtl: 0, muxCtlOFS: 0x01FC, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_CSI_DATA06__ECSPI1_MOSI     = &pad{padCtlOFS: 0x0488, padCtl: 0, muxCtlOFS: 0x01FC, muxMode: 3, selInputOFS: 0x053C, selInput: 1}
	MX6_PAD_CSI_DATA06__EIM_AD06        = &pad{padCtlOFS: 0x0488, padCtl: 0, muxCtlOFS: 0x01FC, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_CSI_DATA06__GPIO4_IO27      = &pad{padCtlOFS: 0x0488, padCtl: 0, muxCtlOFS: 0x01FC, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_CSI_DATA06__SAI1_RX_DATA    = &pad{padCtlOFS: 0x0488, padCtl: 0, muxCtlOFS: 0x01FC, muxMode: 6, selInputOFS: 0x05E4, selInput: 1}
	MX6_PAD_CSI_DATA06__USDHC1_RESET_B  = &pad{padCtlOFS: 0x0488, padCtl: 0, muxCtlOFS: 0x01FC, muxMode: 8, selInputOFS: 0x0000, selInput: 0}

	MX6_PAD_CSI_DATA07__CSI_DATA09      = &pad{padCtlOFS: 0x048C, padCtl: 0, muxCtlOFS: 0x0200, muxMode: 0, selInputOFS: 0x04E8, selInput: 1}
	MX6_PAD_CSI_DATA07__USDHC2_DATA7    = &pad{padCtlOFS: 0x048C, padCtl: 0, muxCtlOFS: 0x0200, muxMode: 1, selInputOFS: 0x0698, selInput: 2}
	MX6_PAD_CSI_DATA07__SIM2_PORT1_TRXD = &pad{padCtlOFS: 0x048C, padCtl: 0, muxCtlOFS: 0x0200, muxMode: 2, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_CSI_DATA07__ECSPI1_MISO     = &pad{padCtlOFS: 0x048C, padCtl: 0, muxCtlOFS: 0x0200, muxMode: 3, selInputOFS: 0x0538, selInput: 1}
	MX6_PAD_CSI_DATA07__EIM_AD07        = &pad{padCtlOFS: 0x048C, padCtl: 0, muxCtlOFS: 0x0200, muxMode: 4, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_CSI_DATA07__GPIO4_IO28      = &pad{padCtlOFS: 0x048C, padCtl: 0, muxCtlOFS: 0x0200, muxMode: 5, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_CSI_DATA07__SAI1_TX_DATA    = &pad{padCtlOFS: 0x048C, padCtl: 0, muxCtlOFS: 0x0200, muxMode: 6, selInputOFS: 0x0000, selInput: 0}
	MX6_PAD_CSI_DATA07__USDHC1_VSELECT  = &pad{padCtlOFS: 0x048C, padCtl: 0, muxCtlOFS: 0x0200, muxMode: 8, selInputOFS: 0x0000, selInput: 0}
)

type lcdPad struct {
	p *pad
	v uint32
}

var lcdPads = []lcdPad{
	{p: MX6_PAD_LCD_CLK__LCDIF_CLK, v: PAD_CTL_HYS | PAD_CTL_PUS_100K_UP | PAD_CTL_PUE | PAD_CTL_PKE | PAD_CTL_SPEED_MED | PAD_CTL_DSE_40ohm},
	{p: MX6_PAD_LCD_ENABLE__LCDIF_ENABLE, v: PAD_CTL_HYS | PAD_CTL_PUS_100K_UP | PAD_CTL_PUE | PAD_CTL_PKE | PAD_CTL_SPEED_MED | PAD_CTL_DSE_40ohm},
	{p: MX6_PAD_LCD_HSYNC__LCDIF_HSYNC, v: PAD_CTL_HYS | PAD_CTL_PUS_100K_UP | PAD_CTL_PUE | PAD_CTL_PKE | PAD_CTL_SPEED_MED | PAD_CTL_DSE_40ohm},
	{p: MX6_PAD_LCD_VSYNC__LCDIF_VSYNC, v: PAD_CTL_HYS | PAD_CTL_PUS_100K_UP | PAD_CTL_PUE | PAD_CTL_PKE | PAD_CTL_SPEED_MED | PAD_CTL_DSE_40ohm},
	{p: MX6_PAD_LCD_DATA00__LCDIF_DATA00, v: PAD_CTL_HYS | PAD_CTL_PUS_100K_UP | PAD_CTL_PUE | PAD_CTL_PKE | PAD_CTL_SPEED_MED | PAD_CTL_DSE_40ohm},
	{p: MX6_PAD_LCD_DATA01__LCDIF_DATA01, v: PAD_CTL_HYS | PAD_CTL_PUS_100K_UP | PAD_CTL_PUE | PAD_CTL_PKE | PAD_CTL_SPEED_MED | PAD_CTL_DSE_40ohm},
	{p: MX6_PAD_LCD_DATA02__LCDIF_DATA02, v: PAD_CTL_HYS | PAD_CTL_PUS_100K_UP | PAD_CTL_PUE | PAD_CTL_PKE | PAD_CTL_SPEED_MED | PAD_CTL_DSE_40ohm},
	{p: MX6_PAD_LCD_DATA03__LCDIF_DATA03, v: PAD_CTL_HYS | PAD_CTL_PUS_100K_UP | PAD_CTL_PUE | PAD_CTL_PKE | PAD_CTL_SPEED_MED | PAD_CTL_DSE_40ohm},
	{p: MX6_PAD_LCD_DATA04__LCDIF_DATA04, v: PAD_CTL_HYS | PAD_CTL_PUS_100K_UP | PAD_CTL_PUE | PAD_CTL_PKE | PAD_CTL_SPEED_MED | PAD_CTL_DSE_40ohm},
	{p: MX6_PAD_LCD_DATA05__LCDIF_DATA05, v: PAD_CTL_HYS | PAD_CTL_PUS_100K_UP | PAD_CTL_PUE | PAD_CTL_PKE | PAD_CTL_SPEED_MED | PAD_CTL_DSE_40ohm},
	{p: MX6_PAD_LCD_DATA06__LCDIF_DATA06, v: PAD_CTL_HYS | PAD_CTL_PUS_100K_UP | PAD_CTL_PUE | PAD_CTL_PKE | PAD_CTL_SPEED_MED | PAD_CTL_DSE_40ohm},
	{p: MX6_PAD_LCD_DATA07__LCDIF_DATA07, v: PAD_CTL_HYS | PAD_CTL_PUS_100K_UP | PAD_CTL_PUE | PAD_CTL_PKE | PAD_CTL_SPEED_MED | PAD_CTL_DSE_40ohm},
	{p: MX6_PAD_LCD_DATA08__LCDIF_DATA08, v: PAD_CTL_HYS | PAD_CTL_PUS_100K_UP | PAD_CTL_PUE | PAD_CTL_PKE | PAD_CTL_SPEED_MED | PAD_CTL_DSE_40ohm},
	{p: MX6_PAD_LCD_DATA09__LCDIF_DATA09, v: PAD_CTL_HYS | PAD_CTL_PUS_100K_UP | PAD_CTL_PUE | PAD_CTL_PKE | PAD_CTL_SPEED_MED | PAD_CTL_DSE_40ohm},
	{p: MX6_PAD_LCD_DATA10__LCDIF_DATA10, v: PAD_CTL_HYS | PAD_CTL_PUS_100K_UP | PAD_CTL_PUE | PAD_CTL_PKE | PAD_CTL_SPEED_MED | PAD_CTL_DSE_40ohm},
	{p: MX6_PAD_LCD_DATA11__LCDIF_DATA11, v: PAD_CTL_HYS | PAD_CTL_PUS_100K_UP | PAD_CTL_PUE | PAD_CTL_PKE | PAD_CTL_SPEED_MED | PAD_CTL_DSE_40ohm},
	{p: MX6_PAD_LCD_DATA12__LCDIF_DATA12, v: PAD_CTL_HYS | PAD_CTL_PUS_100K_UP | PAD_CTL_PUE | PAD_CTL_PKE | PAD_CTL_SPEED_MED | PAD_CTL_DSE_40ohm},
	{p: MX6_PAD_LCD_DATA13__LCDIF_DATA13, v: PAD_CTL_HYS | PAD_CTL_PUS_100K_UP | PAD_CTL_PUE | PAD_CTL_PKE | PAD_CTL_SPEED_MED | PAD_CTL_DSE_40ohm},
	{p: MX6_PAD_LCD_DATA14__LCDIF_DATA14, v: PAD_CTL_HYS | PAD_CTL_PUS_100K_UP | PAD_CTL_PUE | PAD_CTL_PKE | PAD_CTL_SPEED_MED | PAD_CTL_DSE_40ohm},
	{p: MX6_PAD_LCD_DATA15__LCDIF_DATA15, v: PAD_CTL_HYS | PAD_CTL_PUS_100K_UP | PAD_CTL_PUE | PAD_CTL_PKE | PAD_CTL_SPEED_MED | PAD_CTL_DSE_40ohm},
	{p: MX6_PAD_LCD_DATA16__LCDIF_DATA16, v: PAD_CTL_HYS | PAD_CTL_PUS_100K_UP | PAD_CTL_PUE | PAD_CTL_PKE | PAD_CTL_SPEED_MED | PAD_CTL_DSE_40ohm},
	{p: MX6_PAD_LCD_DATA17__LCDIF_DATA17, v: PAD_CTL_HYS | PAD_CTL_PUS_100K_UP | PAD_CTL_PUE | PAD_CTL_PKE | PAD_CTL_SPEED_MED | PAD_CTL_DSE_40ohm},
	{p: MX6_PAD_LCD_DATA18__LCDIF_DATA18, v: PAD_CTL_HYS | PAD_CTL_PUS_100K_UP | PAD_CTL_PUE | PAD_CTL_PKE | PAD_CTL_SPEED_MED | PAD_CTL_DSE_40ohm},
	{p: MX6_PAD_LCD_DATA19__LCDIF_DATA19, v: PAD_CTL_HYS | PAD_CTL_PUS_100K_UP | PAD_CTL_PUE | PAD_CTL_PKE | PAD_CTL_SPEED_MED | PAD_CTL_DSE_40ohm},
	{p: MX6_PAD_LCD_DATA20__LCDIF_DATA20, v: PAD_CTL_HYS | PAD_CTL_PUS_100K_UP | PAD_CTL_PUE | PAD_CTL_PKE | PAD_CTL_SPEED_MED | PAD_CTL_DSE_40ohm},
	{p: MX6_PAD_LCD_DATA21__LCDIF_DATA21, v: PAD_CTL_HYS | PAD_CTL_PUS_100K_UP | PAD_CTL_PUE | PAD_CTL_PKE | PAD_CTL_SPEED_MED | PAD_CTL_DSE_40ohm},
	{p: MX6_PAD_LCD_DATA22__LCDIF_DATA22, v: PAD_CTL_HYS | PAD_CTL_PUS_100K_UP | PAD_CTL_PUE | PAD_CTL_PKE | PAD_CTL_SPEED_MED | PAD_CTL_DSE_40ohm},
	{p: MX6_PAD_LCD_DATA23__LCDIF_DATA23, v: PAD_CTL_HYS | PAD_CTL_PUS_100K_UP | PAD_CTL_PUE | PAD_CTL_PKE | PAD_CTL_SPEED_MED | PAD_CTL_DSE_40ohm},

	{p: MX6_PAD_SNVS_TAMPER9__GPIO5_IO09, v: NO_PAD_CTRL},

	{p: MX6_PAD_GPIO1_IO08__GPIO1_IO08, v: NO_PAD_CTRL},
}
