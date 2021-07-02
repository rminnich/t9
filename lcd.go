package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// these need to make it to tamago
const (
	CCGR3  = 0x020C4074
	CCGR2  = 0x020C4070
	cscdr2 = 0x020C4038
	/* i.MX6SX/UL LCD and PXP */
	MXC_CCM_CCGR2_LCD_OFFSET = 28
	MXC_CCM_CCGR2_LCD_MASK   = (3 << MXC_CCM_CCGR2_LCD_OFFSET)
	MXC_CCM_CCGR2_PXP_OFFSET = 30
	MXC_CCM_CCGR2_PXP_MASK   = (3 << MXC_CCM_CCGR2_PXP_OFFSET)

	MXC_CCM_CSCDR2_LCDIF1_CLK_SEL_MASK = (0x7 << 9)
	MXC_CCM_CCGR3_LCDIF_PIX_OFFSET     = 8
	MXC_CCM_CCGR3_LCDIF_PIX_MASK       = (3 << MXC_CCM_CCGR3_LCDIF_PIX_OFFSET)

	MXC_CCM_CCGR3_LCDIF1_PIX_OFFSET = 10
	MXC_CCM_CCGR3_LCDIF1_PIX_MASK   = (3 << MXC_CCM_CCGR3_LCDIF1_PIX_OFFSET)
	lcdbase                         = 0x20e_0000
)

var (
	// These are used when run standalone not on the board.
	ccm = flag.String("ccm", "/dev/ccm", "Device to be used for the CCM")
)

func init() {
	flag.Parse()
}

/*
 * configures a single pad in the iomuxer
 */
func onePad(w io.WriterAt, p *pad) {
	mux_ctrl_ofs := p.muxCtlOFS + lcdbase
	mux_mode := p.muxMode
	sel_input_ofs := p.selInputOFS + lcdbase
	sel_input := p.selInput
	pad_ctrl_ofs := p.padCtlOFS + lcdbase
	pad_ctrl := p.padCtl

	// no LPRR?lpsr := (pad & MUX_MODE_LPSR) >> MUX_MODE_SHIFT;

	// #ifdef CONFIG_MX7
	// 	if (lpsr == IOMUX_CONFIG_LPSR) {
	// 		base = (void *)IOMUXC_LPSR_BASE_ADDR;
	// 		mux_mode &= ~IOMUX_CONFIG_LPSR;
	// 		/* set daisy chain sel_input */
	// 		if (sel_input_ofs)
	// 			sel_input_ofs += IOMUX_LPSR_SEL_INPUT_OFS;
	// 	}
	// #else
	// 	if (is_mx6ull() || is_mx6sll()) {
	// 		if (lpsr == IOMUX_CONFIG_LPSR) {
	// 			base = (void *)IOMUXC_SNVS_BASE_ADDR;
	// 			mux_mode &= ~IOMUX_CONFIG_LPSR;
	// 		}
	// 	}
	// #endif
	// #endif

	//	if (is_mx7() || is_mx6ull() || is_mx6sll() || mux_ctrl_ofs)
	writel(w, mux_mode, mux_ctrl_ofs)

	if sel_input_ofs != 0 {
		writel(w, sel_input, sel_input_ofs)
	}

	// #ifdef CONFIG_IOMUX_SHARE_CONF_REG
	// 	if (!(pad_ctrl & NO_PAD_CTRL))
	// 		__raw_writel((mux_mode << PAD_MUX_MODE_SHIFT) | pad_ctrl,
	// 			base + pad_ctrl_ofs);
	// #else
	// whatevs	if (!(pad_ctrl & NO_PAD_CTRL) && pad_ctrl_ofs)
	writel(w, pad_ctrl, pad_ctrl_ofs)

	// #ifdef CONFIG_IOMUX_LPSR
	// 	if (lpsr == IOMUX_CONFIG_LPSR)
	// 		base = (void *)IOMUXC_BASE_ADDR;
	// #endif

}

/* configures a list of pads within declared with IOMUX_PADS macro */
func doPads(w io.WriterAt, pads []lcdPad) {
	for _, p := range pads {
		onePad(w, p.p)
	}
}

func NewLCD(enable bool) error {
	ccm, err := os.OpenFile(*ccm, os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer ccm.Close()
	var lcdif_clk_sel_mask uint32 = MXC_CCM_CSCDR2_LCDIF1_CLK_SEL_MASK
	var lcdif_ccgr3_mask uint32 = MXC_CCM_CCGR3_LCDIF1_PIX_MASK
	/* Gate LCDIF clock first */
	// reg = readl(&imx_ccm->CCGR3);
	// reg &= ~lcdif_ccgr3_mask;
	// writel(reg, &imx_ccm->CCGR3);

	r, err := readl(ccm, CCGR3)
	if err != nil {
		return err
	}
	log.Printf("CCG#3 is %#x", r)
	log.Printf("Clear %#x", lcdif_ccgr3_mask)
	if err := bitclr(ccm, lcdif_ccgr3_mask, CCGR3); err != nil {
		return fmt.Errorf("Gate LCDIF clock step 1: %v", err)
	}
	r, err = readl(ccm, CCGR3)
	if err != nil {
		return err
	}
	log.Printf("CCG#3 is %#x", r)

	// reg = readl(&imx_ccm->CCGR2);
	// reg &= ~MXC_CCM_CCGR2_LCD_MASK;
	// writel(reg, &imx_ccm->CCGR2);

	if err := bitclr(ccm, MXC_CCM_CCGR2_LCD_MASK, CCGR2); err != nil {
		return fmt.Errorf("Gate LCDIF clock step 1: %v", err)
	}
	if enable {
		/* Select pre-mux */
		if err := bitclr(ccm, lcdif_clk_sel_mask, cscdr2); err != nil {
			return fmt.Errorf("Select pre-mux: %v", err)
		}
		// reg = readl(ccm, cscdr2)
		// reg &= ^lcdif_clk_sel_mask
		// writel(ccm, reg, cscdr2)

		/* Enable the LCDIF pix clock */
		if err := bitset(ccm, lcdif_ccgr3_mask, CCGR3); err != nil {
			return fmt.Errorf("Setting LCDIF pix clock: %v", err)
		}

		// reg = readl(ccm, CCGR3)
		// reg |= lcdif_ccgr3_mask
		// writel(ccm, reg, CCGR3)

		if err := bitset(ccm, MXC_CCM_CCGR2_LCD_MASK, CCGR2); err != nil {
			return fmt.Errorf("Select pre-mux: %v", err)
		}
		// reg = readl(ccm, CCGR2)
		// reg |= MXC_CCM_CCGR2_LCD_MASK
		// writel(ccm, reg, CCGR2)
	}

	doPads(ccm, lcdPads)
	// /* Reset the LCD */
	// gpio_request(IMX_GPIO_NR(5, 9), "lcd reset")
	// gpio_direction_output(IMX_GPIO_NR(5, 9), 0)
	// udelay(500)
	// gpio_direction_output(IMX_GPIO_NR(5, 9), 1)
	// We're going 0 relative.
	if g := NewGPIO(ccm, 4, 9, "lcd reset").Set(0).Output().Delay(func() error {
		time.Sleep(500 * time.Microsecond)
		return nil
	}).Set(1); g.err != nil {
		return g.err
	}

	// /* Set Brightness to high */
	// gpio_request(IMX_GPIO_NR(1, 8), "backlight")
	// gpio_direction_output(IMX_GPIO_NR(1, 8), 1)
	if g := NewGPIO(ccm, 0, 8, "backlight").Set(1).Output(); err != nil {
		return g.err
	}
	return nil
}
