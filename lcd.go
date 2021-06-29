package main

import (
	"flag"
	"fmt"
	"log"
	"os"
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
)

var (
	// These are used when run standalone not on the board.
	ccm = flag.String("ccm", "/dev/ccm", "Device to be used for the CCM")
)

func init() {
	flag.Parse()
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
	log.Print("Clear %#x", lcdif_ccgr3_mask)
	if err := bitclr(ccm, lcdif_ccgr3_mask, CCGR3); err != nil {
		return fmt.Errorf("Gate LCDIF clock step 1: %v", err)
	}
	r, err = readl(ccm, CCGR3)
	if err != nil {
		return err
	}
	log.Printf("CCG#3 is %#x", r)
	return nil

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
	return nil
}
