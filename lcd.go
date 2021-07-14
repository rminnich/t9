package main

// A lot of this file is from u-boot, and so we inherit the GPL-2.0+.
/*
 * (C) Copyright 2004
 * Pierre Aubert, Staubli Faverges , <p.aubert@staubli.com>
 * Copyright 2011 Freescale Semiconductor, Inc.
 */

/************************************************************************
  Get Parameters for the video mode:
  The default video mode can be defined in CONFIG_SYS_DEFAULT_VIDEO_MODE.
  If undefined, default video mode is set to 0x301
  Parameters can be set via the variable "videomode" in the environment.
  2 diferent ways are possible:
  "videomode=301"   - 301 is a hexadecimal number describing the VESA
		      mode. Following modes are implemented:

		      Colors	640x480 800x600 1024x768 1152x864 1280x1024
		     --------+---------------------------------------------
		      8 bits |	0x301	0x303	 0x305	  0x161	    0x307
		     15 bits |	0x310	0x313	 0x316	  0x162	    0x319
		     16 bits |	0x311	0x314	 0x317	  0x163	    0x31A
		     24 bits |	0x312	0x315	 0x318	    ?	    0x31B
		     --------+---------------------------------------------
  "videomode=bootargs"
		   - the parameters are parsed from the bootargs.
		      The format is "NAME:VALUE,NAME:VALUE" etc.
		      Ex.:
		      "bootargs=video=ctfb:x:800,y:600,depth:16,pclk:25000"
		      Parameters not included in the list will be taken from
		      the default mode, which is one of the following:
		      mode:0  640x480x24
		      mode:1  800x600x16
		      mode:2  1024x768x8
		      mode:3  960x720x24
		      mode:4  1152x864x16
		      mode:5  1280x1024x8

		      if "mode" is not provided within the parameter list,
		      mode:0 is assumed.
		      Following parameters are supported:
		      x	      xres = visible resolution horizontal
		      y	      yres = visible resolution vertical
		      pclk    pixelclocks in pico sec
		      le      left_marging time from sync to picture in pixelclocks
		      ri      right_marging time from picture to sync in pixelclocks
		      up      upper_margin time from sync to picture
		      lo      lower_margin
		      hs      hsync_len length of horizontal sync
		      vs      vsync_len length of vertical sync
		      sync    see FB_SYNC_*
		      vmode   see FB_VMODE_*
		      depth   Color depth in bits per pixel
		      All other parameters in the variable bootargs are ignored.
		      It is also possible to set the parameters direct in the
		      variable "videomode", or in another variable i.e.
		      "myvideo" and setting the variable "videomode=myvideo"..
****************************************************************************/

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

const (
	/* Some mode definitions */
	FB_SYNC_HOR_HIGH_ACT  = 1  /* horizontal sync high active	*/
	FB_SYNC_VERT_HIGH_ACT = 2  /* vertical sync high active	*/
	FB_SYNC_EXT           = 4  /* external sync		*/
	FB_SYNC_COMP_HIGH_ACT = 8  /* composite sync high active	*/
	FB_SYNC_BROADCAST     = 16 /* broadcast video timings	*/
	/* vtotal = 144d/288n/576i => PAL  */
	/* vtotal = 121d/242n/484i => NTSC */
	FB_SYNC_ON_GREEN       = 32 /* sync on green */
	FB_VMODE_NONINTERLACED = 0  /* non interlaced */
	FB_VMODE_INTERLACED    = 1  /* interlaced	*/
	FB_VMODE_DOUBLE        = 2  /* double scan */
	FB_VMODE_MASK          = 255

	FB_VMODE_YWRAP       = 256 /* ywrap instead of panning	*/
	FB_VMODE_SMOOTH_XPAN = 512 /* smooth xpan possible (internally used) */
	FB_VMODE_CONUPDATE   = 512 /* don't update x/yoffset	*/
)

/******************************************************************
 * Resolution Struct
 ******************************************************************/
type ctfb_res_modes struct {
	int xres /* visible resolution		*/
	int yres
	int refresh /* vertical refresh rate in hz  */
	/* Timing: All values in pixclocks, except pixclock (of course) */
	int pixclock     /* pixel clock in ps (pico seconds) */
	int pixclock_khz /* pixel clock in kHz           */
	int left_margin  /* time from sync to picture	*/
	int right_margin /* time from picture to sync	*/
	int upper_margin /* time from sync to picture	*/
	int lower_margin
	int hsync_len /* length of horizontal sync	*/
	int vsync_len /* length of vertical sync	*/
	int sync      /* see FB_SYNC_*		*/
	int vmode     /* see FB_VMODE_*		*/
}

/******************************************************************
 * Vesa Mode Struct
 ******************************************************************/
type ctfb_vesa_modes struct {
	vesanr         int /* Vesa number as in LILO (VESA Nr + 0x200} */
	resindex       int /* index to resolution struct */
	bits_per_pixel int /* bpp */
}

const (
	RES_MODE_640x480   = 0
	RES_MODE_800x600   = 1
	RES_MODE_1024x768  = 2
	RES_MODE_960_720   = 3
	RES_MODE_1152x864  = 4
	RES_MODE_1280x1024 = 5
	RES_MODE_1280x720  = 6
	RES_MODE_1360x768  = 7
	RES_MODE_1920x1080 = 8
	RES_MODE_1920x1200 = 9
	RES_MODES_COUNT    = 10

	VESA_MODES_COUNT = 19
)

// VESA eh? VESA will never die.

var vesa_modes = []ctfb_vesa_modes{
	{0x301, RES_MODE_640x480, 8},
	{0x310, RES_MODE_640x480, 15},
	{0x311, RES_MODE_640x480, 16},
	{0x312, RES_MODE_640x480, 24},
	{0x303, RES_MODE_800x600, 8},
	{0x313, RES_MODE_800x600, 15},
	{0x314, RES_MODE_800x600, 16},
	{0x315, RES_MODE_800x600, 24},
	{0x305, RES_MODE_1024x768, 8},
	{0x316, RES_MODE_1024x768, 15},
	{0x317, RES_MODE_1024x768, 16},
	{0x318, RES_MODE_1024x768, 24},
	{0x161, RES_MODE_1152x864, 8},
	{0x162, RES_MODE_1152x864, 15},
	{0x163, RES_MODE_1152x864, 16},
	{0x307, RES_MODE_1280x1024, 8},
	{0x319, RES_MODE_1280x1024, 15},
	{0x31A, RES_MODE_1280x1024, 16},
	{0x31B, RES_MODE_1280x1024, 24},
}

// This is a mess. Which one is it? STD?
var res_mode_init = []ctfb_res_modes{
	/*  x     y  hz  pixclk ps/kHz   le   ri  up  lo   hs vs  s  vmode */
	// #ifndef CONFIG_VIDEO_STD_TIMINGS
	{640, 480, 60, 39721, 25180, 40, 24, 32, 11, 96, 2, 0, FB_VMODE_NONINTERLACED},
	{800, 600, 60, 27778, 36000, 64, 24, 22, 1, 72, 2, 0, FB_VMODE_NONINTERLACED},
	{1024, 768, 60, 15384, 65000, 168, 8, 29, 3, 144, 4, 0, FB_VMODE_NONINTERLACED},
	{960, 720, 80, 13100, 76335, 160, 40, 32, 8, 80, 4, 0, FB_VMODE_NONINTERLACED},
	{1152, 864, 60, 12004, 83300, 200, 64, 32, 16, 80, 4, 0, FB_VMODE_NONINTERLACED},
	{1280, 1024, 60, 9090, 110000, 200, 48, 26, 1, 184, 3, 0, FB_VMODE_NONINTERLACED},
	//	{ 640,  480, 60, 39683,  25200,  48,  16, 33, 10,  96, 2, 0, FB_VMODE_NONINTERLACED},
	//	{ 800,  600, 60, 25000,  40000,  88,  40, 23,  1, 128, 4, FB_SYNC_HOR_HIGH_ACT | FB_SYNC_VERT_HIGH_ACT, FB_VMODE_NONINTERLACED},
	//	{1024,  768, 60, 15384,  65000, 160,  24, 29,  3, 136, 6, 0, FB_VMODE_NONINTERLACED},
	//	{ 960,  720, 75, 13468,  74250, 176,  72, 27,  1, 112, 2, 0, FB_VMODE_NONINTERLACED},
	//	{1152,  864, 75,  9259, 108000, 256,  64, 32,  1, 128, 3, FB_SYNC_HOR_HIGH_ACT | FB_SYNC_VERT_HIGH_ACT, FB_VMODE_NONINTERLACED},
	//	{1280, 1024, 60,  9259, 108000, 248,  48, 38,  1, 112, 3, FB_SYNC_HOR_HIGH_ACT | FB_SYNC_VERT_HIGH_ACT, FB_VMODE_NONINTERLACED},
	{1280, 720, 60, 13468, 74250, 220, 110, 20, 5, 40, 5, FB_SYNC_HOR_HIGH_ACT | FB_SYNC_VERT_HIGH_ACT, FB_VMODE_NONINTERLACED},
	{1360, 768, 60, 11696, 85500, 256, 64, 17, 3, 112, 7, 0, FB_VMODE_NONINTERLACED},
	{1920, 1080, 60, 6734, 148500, 148, 88, 36, 4, 44, 5, FB_SYNC_HOR_HIGH_ACT | FB_SYNC_VERT_HIGH_ACT, FB_VMODE_NONINTERLACED},
	{1920, 1200, 60, 6494, 154000, 80, 48, 26, 3, 32, 6, FB_SYNC_HOR_HIGH_ACT, FB_VMODE_NONINTERLACED},
}

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
	// They have this horrible oddball videomode variable, which I leave here, but we won't use.
	//mode = flag.String("videomode", "videomode=video=ctfb:x:480,y:272,depth:24,pclk:108695,le:8,ri:4,up:2,lo:4,hs:41,vs:10,sync:0,vmode:0", "video mode")
	x     = flag.Int("x", 480, "x")
	y     = flag.Int("y", 272, "y")
	depth = flag.Int("depth", 24, "depth")
	pclk  = flag.Int("pclk", 108695, "pclk")
	le    = flag.Int("le", 8, "le")
	ri    = flag.Int("ri", 4, "ri")
	up    = flag.Int("up", 2, "up")
	lo    = flag.Int("lo", 4, "lo")
	hs    = flag.Int("hs", 41, "hs")
	vs    = flag.Int("vs", 10, "vs")
	sync  = flag.Int("sync", 0, "sync")
	vmode = flag.Int("vmode", 0, "vmode")
	mode  = res_mode_init[0]
	bpp   int
)

func init() {
	flag.Parse()
	if *vmode > len(res_mode_init) {
		log.Fatalf("Mode %d is out of range 0 - %d", *vmode, len(res_mode_init))
	}
	mode = res_mode_init[*vmode]
	bpp = 24 - ((*vmode % 3) * 8)
	if false {
		for i := range res_mode_init {
			if res_mode_init[i].xres == xres &&
				res_mode_init[i].yres == yres &&
				res_mode_init[i].refresh == refresh {
				mode = res_mode_init[i]
				break
			}
		}
	}
}

/**
 * Convert an EDID detailed timing to a struct ctfb_res_modes
 *
 * @param t		The EDID detailed timing to be converted
 * @param mode		Returns the converted timing
 *
 * @return 0 on success, or a negative errno on error
 */
// int video_edid_dtd_to_ctfb_res_modes(struct edid_detailed_timing *t,
// 				     struct ctfb_res_modes *mode)
// {
// 	int margin, h_total, v_total;

// 	/* Check all timings are non 0 */
// 	if (EDID_DETAILED_TIMING_PIXEL_CLOCK(*t) == 0 ||
// 	    EDID_DETAILED_TIMING_HORIZONTAL_ACTIVE(*t) == 0 ||
// 	    EDID_DETAILED_TIMING_HORIZONTAL_BLANKING(*t) == 0 ||
// 	    EDID_DETAILED_TIMING_VERTICAL_ACTIVE(*t) == 0 ||
// 	    EDID_DETAILED_TIMING_VERTICAL_BLANKING(*t) == 0 ||
// 	    EDID_DETAILED_TIMING_HSYNC_OFFSET(*t) == 0 ||
// 	    EDID_DETAILED_TIMING_VSYNC_OFFSET(*t) == 0 ||
// 	    /* 3d formats are not supported*/
// 	    EDID_DETAILED_TIMING_FLAG_STEREO(*t) != 0)
// 		return -EINVAL;

// 	mode->xres = EDID_DETAILED_TIMING_HORIZONTAL_ACTIVE(*t);
// 	mode->yres = EDID_DETAILED_TIMING_VERTICAL_ACTIVE(*t);

// 	h_total = mode->xres + EDID_DETAILED_TIMING_HORIZONTAL_BLANKING(*t);
// 	v_total = mode->yres + EDID_DETAILED_TIMING_VERTICAL_BLANKING(*t);
// 	mode->refresh = EDID_DETAILED_TIMING_PIXEL_CLOCK(*t) /
// 			(h_total * v_total);

// 	mode->pixclock_khz = EDID_DETAILED_TIMING_PIXEL_CLOCK(*t) / 1000;
// 	mode->pixclock = 1000000000L / mode->pixclock_khz;

// 	mode->right_margin = EDID_DETAILED_TIMING_HSYNC_OFFSET(*t);
// 	mode->hsync_len = EDID_DETAILED_TIMING_HSYNC_PULSE_WIDTH(*t);
// 	margin = EDID_DETAILED_TIMING_HORIZONTAL_BLANKING(*t) -
// 			(mode->right_margin + mode->hsync_len);
// 	if (margin <= 0)
// 		return -EINVAL;

// 	mode->left_margin = margin;

// 	mode->lower_margin = EDID_DETAILED_TIMING_VSYNC_OFFSET(*t);
// 	mode->vsync_len = EDID_DETAILED_TIMING_VSYNC_PULSE_WIDTH(*t);
// 	margin = EDID_DETAILED_TIMING_VERTICAL_BLANKING(*t) -
// 			(mode->lower_margin + mode->vsync_len);
// 	if (margin <= 0)
// 		return -EINVAL;

// 	mode->upper_margin = margin;

// 	mode->sync = 0;
// 	if (EDID_DETAILED_TIMING_FLAG_HSYNC_POLARITY(*t))
// 		mode->sync |= FB_SYNC_HOR_HIGH_ACT;
// 	if (EDID_DETAILED_TIMING_FLAG_VSYNC_POLARITY(*t))
// 		mode->sync |= FB_SYNC_VERT_HIGH_ACT;

// 	if (EDID_DETAILED_TIMING_FLAG_INTERLACED(*t))
// 		mode->vmode = FB_VMODE_INTERLACED;
// 	else
// 		mode->vmode = FB_VMODE_NONINTERLACED;

// 	return 0;
// }

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
