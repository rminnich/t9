package main

// From various include hell

// /include/configs/mx6ul_14x14_evk.h:#define MXS_LCDIF_BASE MX6UL_LCDIF1_BASE_ADDR
// arch/arm/include/asm/arch-mx6/imx-regs.h:#define MX6UL_LCDIF1_BASE_ADDR      (AIPS2_OFF_BASE_ADDR + 0x48000)
// arch/arm/include/asm/arch-mx6/imx-regs.h:#define AIPS2_OFF_BASE_ADDR         (ATZ2_BASE_ADDR + 0x80000)
// arch/arm/include/asm/arch-mx6/imx-regs.h:#define ATZ2_BASE_ADDR              AIPS2_ARB_BASE_ADDR
// arch/arm/include/asm/arch-mx6/imx-regs.h:#define AIPS2_ARB_BASE_ADDR             0x02100000

// Somebody, somewhere, thinks this is cool :-)
// So for lcd the thing we mess with is 0x0120_0000 + 0x8_0000 + 0x4_8000 = 0x012c8000 ... ?
// The first thing the code does:
// mxsfb.c
// mxs_set_lcdclk(MXS_LCDIF_BASE, PS2KHZ(mode->pixclock));
// mode is the mess define below, which we declined to do.
// see lcd

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

// these need to make it to tamago
const (
	CCGR3  = 0x020C4074
	CCGR2  = 0x020C4070
	cscdr2 = 0x020C4038
	cbcmr  = 0x020C4018
	/* i.MX6SX/UL LCD and PXP */
	MXC_CCM_CCGR2_LCD_OFFSET = 28
	MXC_CCM_CCGR2_LCD_MASK   = (3 << MXC_CCM_CCGR2_LCD_OFFSET)
	MXC_CCM_CCGR2_PXP_OFFSET = 30
	MXC_CCM_CCGR2_PXP_MASK   = (3 << MXC_CCM_CCGR2_PXP_OFFSET)

	MXC_CCM_CCGR3_LCDIF_PIX_OFFSET = 8
	MXC_CCM_CCGR3_LCDIF_PIX_MASK   = (3 << MXC_CCM_CCGR3_LCDIF_PIX_OFFSET)

	MXC_CCM_CCGR3_LCDIF1_PIX_OFFSET = 10
	MXC_CCM_CCGR3_LCDIF1_PIX_MASK   = (3 << MXC_CCM_CCGR3_LCDIF1_PIX_OFFSET)
	lcdbase                         = 0x20e_0000
)

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

// #ifdef CONFIG_VIDEO
// #define CONFIG_VIDEO_MXS
// #define CONFIG_VIDEO_LOGO
// #define CONFIG_SPLASH_SCREEN
// #define CONFIG_SPLASH_SCREEN_ALIGN
// #define CONFIG_BMP_16BPP
// #define CONFIG_VIDEO_BMP_RLE8
// #define CONFIG_VIDEO_BMP_LOGO
// #define MXS_LCDIF_BASE MX6UL_LCDIF1_BASE_ADDR

/******************************************************************
 * Resolution Struct
 ******************************************************************/
type ctfb_res_modes struct {
	xres    uint32 /* visible resolution		*/
	yres    uint32
	refresh uint32 /* vertical refresh rate in hz  */
	/* Timing: All values in pixclocks, except pixclock (of course) */
	pixclock     uint32 /* pixel clock in ps (pico seconds) */
	pixclock_khz uint32 /* pixel clock in kHz           */
	left_margin  uint32 /* time from sync to picture	*/
	right_margin uint32 /* time from picture to sync	*/
	upper_margin uint32 /* time from sync to picture	*/
	lower_margin uint32
	hsync_len    uint32 /* length of horizontal sync	*/
	vsync_len    uint32 /* length of vertical sync	*/
	sync         uint32 /* see FB_SYNC_*		*/
	vmode        uint32 /* see FB_VMODE_*		*/
}

/******************************************************************
 * Vesa Mode Struct
 ******************************************************************/
type ctfb_vesa_modes struct {
	vesanr         uint32 /* Vesa number as in LILO (VESA Nr + 0x200} */
	resindex       uint32 /* index to resolution struct */
	bits_per_pixel uint32 /* bpp */
}

type panel struct {
	winSizeX   uint32
	winSizeY   uint32
	plnSizeX   uint32
	plnSizeY   uint32
	gdfBytesPP uint32
	gdfIndex   uint32
	memSize    uint32
	frameAdrs  uint32
}

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

var (
	// These are used when run standalone not on the board.
	ccm = flag.String("ccm", "/dev/ccm", "Device to be used for the CCM")
	// They have this horrible oddball videomode variable, which I leave here, but we won't use.
	//mode = flag.String("videomode", "videomode=video=ctfb:x:480,y:272,depth:24,pclk:108695,le:8,ri:4,up:2,lo:4,hs:41,vs:10,sync:0,vmode:0", "video mode")
	xres  = flag.Int("xres", 480, "xres")
	yres  = flag.Int("*yres", 272, "*yres")
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
	log.Printf("Got mode #v", mode)
	if false {
		for i := range res_mode_init {
			if res_mode_init[i].xres == uint32(*xres) &&
				res_mode_init[i].yres == uint32(*yres) {
				//			&&	res_mode_init[i].refresh == refresh
				mode = res_mode_init[i]
				log.Printf("Got mode #v", mode)
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

// 	mode.xres = EDID_DETAILED_TIMING_HORIZONTAL_ACTIVE(*t);
// 	mode.yres = EDID_DETAILED_TIMING_VERTICAL_ACTIVE(*t);

// 	h_total = mode.xres + EDID_DETAILED_TIMING_HORIZONTAL_BLANKING(*t);
// 	v_total = mode.yres + EDID_DETAILED_TIMING_VERTICAL_BLANKING(*t);
// 	mode.refresh = EDID_DETAILED_TIMING_PIXEL_CLOCK(*t) /
// 			(h_total * v_total);

// 	mode.pixclock_khz = EDID_DETAILED_TIMING_PIXEL_CLOCK(*t) / 1000;
// 	mode.pixclock = 1000000000L / mode.pixclock_khz;

// 	mode.right_margin = EDID_DETAILED_TIMING_HSYNC_OFFSET(*t);
// 	mode.hsync_len = EDID_DETAILED_TIMING_HSYNC_PULSE_WIDTH(*t);
// 	margin = EDID_DETAILED_TIMING_HORIZONTAL_BLANKING(*t) -
// 			(mode.right_margin + mode.hsync_len);
// 	if (margin <= 0)
// 		return -EINVAL;

// 	mode.left_margin = margin;

// 	mode.lower_margin = EDID_DETAILED_TIMING_VSYNC_OFFSET(*t);
// 	mode.vsync_len = EDID_DETAILED_TIMING_VSYNC_PULSE_WIDTH(*t);
// 	margin = EDID_DETAILED_TIMING_VERTICAL_BLANKING(*t) -
// 			(mode.lower_margin + mode.vsync_len);
// 	if (margin <= 0)
// 		return -EINVAL;

// 	mode.upper_margin = margin;

// 	mode.sync = 0;
// 	if (EDID_DETAILED_TIMING_FLAG_HSYNC_POLARITY(*t))
// 		mode.sync |= FB_SYNC_HOR_HIGH_ACT;
// 	if (EDID_DETAILED_TIMING_FLAG_VSYNC_POLARITY(*t))
// 		mode.sync |= FB_SYNC_VERT_HIGH_ACT;

// 	if (EDID_DETAILED_TIMING_FLAG_INTERLACED(*t))
// 		mode.vmode = FB_VMODE_INTERLACED;
// 	else
// 		mode.vmode = FB_VMODE_NONINTERLACED;

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

func ps2khz(ps uint32) uint32 {
	return 1000000000 / ps
}

/*
 * ARIES M28EVK:
 * setenv videomode
 * video=ctfb:x:800,y:480,depth:18,mode:0,pclk:30066,
 *       le:0,ri:256,up:0,lo:45,hs:1,vs:1,sync:100663296,vmode:0
 *
 * Freescale mx23evk/mx28evk with a Seiko 4.3'' WVGA panel:
 * setenv videomode
 * video=ctfb:x:800,y:480,depth:24,mode:0,pclk:29851,
 * 	 le:89,ri:164,up:23,lo:10,hs:10,vs:10,sync:0,vmode:0
 */

func mxs_set_lcdclk(rw rw, ba uint32, freq uint32) error {
	var (
		reg uint32 = 0
		hck uint32 = MXC_HCLK / 1000
		/* DIV_SELECT ranges from 27 to 54 */
		min                         uint32 = hck * 27
		max                         uint32 = hck * 54
		temp                        uint32
		best                        uint32 = 0
		pll_div, pll_num, pll_denom uint32
		post_div                    uint32 = 1

		max_pred  uint32 = 8
		max_postd uint32 = 8
		pred      uint32 = 1
		postd     uint32 = 1
	)

	// always true now			if (base_addr == LCDIF1_BASE_ADDR) {
	reg, err := readl(rw, cscdr2)
	if err != nil {
		return err
	}
	/* Can't change clocks when clock not from pre-mux */
	if reg&MXC_CCM_CSCDR2_LCDIF1_CLK_SEL_MASK != 0 {
		return fmt.Errorf("Can't change clocks when clock not from pre-mux")
	}

	temp = freq * max_pred * max_postd
	if temp < min {
		/*
		 * Register: PLL_VIDEO
		 * Bit Field: POST_DIV_SELECT
		 * 00 — Divide by 4.
		 * 01 — Divide by 2.
		 * 10 — Divide by 1.
		 * 11 — Reserved
		 * No need to check post_div(1)
		 */
		for post_div := uint32(2); post_div <= 4; post_div <<= 1 {
			if (temp * post_div) > min {
				freq *= post_div
				break
			}
		}

		if post_div > 4 {
			return fmt.Errorf("Fail to set rate to %dkhz", freq)
		}
	}

	/* Choose the best pred and postd to match freq for lcd */
	for i := uint32(1); i <= max_pred; i++ {
		for j := uint32(1); j <= max_postd; j++ {
			temp = freq * i * j
			if temp > max || temp < min {
				continue
			}
			if best == 0 || temp < best {
				best = temp
				pred = i
				postd = j
			}
		}
	}

	if best == 0 {
		return fmt.Errorf("Fail to set rate to %dKHz", freq)
	}

	log.Printf("best %d, pred = %d, postd = %d\n", best, pred, postd)

	pll_div = best / hck
	pll_denom = 1000000
	pll_num = (best - hck*pll_div) * pll_denom / hck

	/*
	 *                                  pll_num
	 *             (24MHz * (pll_div + --------- ))
	 *                                 pll_denom
	 *freq KHz =  --------------------------------
	 *             post_div * pred * postd * 1000
	 */

	if err := enable_pll_video(pll_div, pll_num, pll_denom, post_div); err != nil {
		return err
	}

	enable_lcdif_clock(ba, 0)

	/* Select pre-lcd clock to PLL5 and set pre divider */
	if err := bitsetclr(rw, cscdr2, MXC_CCM_CSCDR2_LCDIF1_PRED_SEL_MASK|MXC_CCM_CSCDR2_LCDIF1_PRE_DIV_MASK, (0x2<<MXC_CCM_CSCDR2_LCDIF1_PRED_SEL_OFFSET)|((pred-1)<<MXC_CCM_CSCDR2_LCDIF1_PRE_DIV_OFFSET)); err != nil {
		return err
	}

	/* Set the post divider */
	if err := bitsetclr(rw, cbcmr, MXC_CCM_CBCMR_LCDIF1_PODF_MASK, ((postd - 1) << MXC_CCM_CBCMR_LCDIF1_PODF_OFFSET)); err != nil {
		return err
	}

	return nil

}
func mxs_lcd_init(rw rw, panel *panel, mode *ctfb_res_modes, bpp int) error {
	var word_len, bus_width uint32
	var valid_data uint32

	/* Kick in the LCDIF clock */
	mxs_set_lcdclk(rw, hw_lcdif_base, ps2khz(mode.pixclock))

	/* Restart the LCDIF block */
	//	mxs_reset_block(&regs.hw_lcdif_ctrl_reg);

	switch bpp {
	case 24:
		word_len = LCDIF_CTRL_WORD_LENGTH_24BIT
		bus_width = LCDIF_CTRL_LCD_DATABUS_WIDTH_24BIT
		valid_data = 0x7
		break
	case 18:
		word_len = LCDIF_CTRL_WORD_LENGTH_24BIT
		bus_width = LCDIF_CTRL_LCD_DATABUS_WIDTH_18BIT
		valid_data = 0x7
		break
	case 16:
		word_len = LCDIF_CTRL_WORD_LENGTH_16BIT
		bus_width = LCDIF_CTRL_LCD_DATABUS_WIDTH_16BIT
		valid_data = 0xf
		break
	case 8:
		word_len = LCDIF_CTRL_WORD_LENGTH_8BIT
		bus_width = LCDIF_CTRL_LCD_DATABUS_WIDTH_8BIT
		valid_data = 0xf
		break
	}

	writel(rw, bus_width|word_len|LCDIF_CTRL_DOTCLK_MODE|LCDIF_CTRL_BYPASS_COUNT|LCDIF_CTRL_LCDIF_MASTER, hw_lcdif_ctrl)

	writel(rw, valid_data<<LCDIF_CTRL1_BYTE_PACKING_FORMAT_OFFSET, hw_lcdif_ctrl1)

	// weeak function in original. mxsfb_system_setup()

	writel(rw, (mode.yres<<LCDIF_TRANSFER_COUNT_V_COUNT_OFFSET)|mode.xres, hw_lcdif_transfer_count)

	writel(rw, LCDIF_VDCTRL0_ENABLE_PRESENT|LCDIF_VDCTRL0_ENABLE_POL|LCDIF_VDCTRL0_VSYNC_PERIOD_UNIT|LCDIF_VDCTRL0_VSYNC_PULSE_WIDTH_UNIT|mode.vsync_len, hw_lcdif_vdctrl0)
	writel(rw, mode.upper_margin+mode.lower_margin+mode.vsync_len+mode.yres, hw_lcdif_vdctrl1)
	writel(rw, (mode.hsync_len<<LCDIF_VDCTRL2_HSYNC_PULSE_WIDTH_OFFSET)|(mode.left_margin+mode.right_margin+mode.hsync_len+mode.xres), hw_lcdif_vdctrl2)
	writel(rw, ((mode.left_margin+mode.hsync_len)<<LCDIF_VDCTRL3_HORIZONTAL_WAIT_CNT_OFFSET)|(mode.upper_margin+mode.vsync_len), hw_lcdif_vdctrl3)
	writel(rw, (0<<LCDIF_VDCTRL4_DOTCLK_DLY_SEL_OFFSET)|mode.xres, hw_lcdif_vdctrl4)

	writel(rw, panel.frameAdrs, hw_lcdif_cur_buf)
	writel(rw, panel.frameAdrs, hw_lcdif_next_buf)

	/* Flush FIFO first */
	writel(rw, LCDIF_CTRL1_FIFO_CLEAR, hw_lcdif_ctrl1+set)

	//	if CONFIG_VIDEO_MXS_MODE_SYSTEM {
	//	/* Sync signals ON */
	//	setbits_le32(hw_lcdif_vdctrl4, LCDIF_VDCTRL4_SYNC_SIGNALS_ON)
	//}

	/* FIFO cleared */
	writel(rw, LCDIF_CTRL1_FIFO_CLEAR, hw_lcdif_ctrl1+clr)

	/* RUN! */
	writel(rw, LCDIF_CTRL_RUN, hw_lcdif_ctrl+set)
	return nil
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

	// more fun.
	var panel = &panel{}
	panel.winSizeX = mode.xres
	panel.winSizeY = mode.yres
	panel.plnSizeX = mode.xres
	panel.plnSizeY = mode.yres

	switch bpp {
	case 24:
	case 18:
		panel.gdfBytesPP = 4
		panel.gdfIndex = GDF_32BIT_X888RGB

	case 16:
		panel.gdfBytesPP = 2
		panel.gdfIndex = GDF_16BIT_565RGB

	case 8:
		panel.gdfBytesPP = 1
		panel.gdfIndex = GDF__8BIT_INDEX

	default:
		return fmt.Errorf("MXSFB: Invalid BPP specified! (bpp = %)\n", bpp)
	}

	panel.memSize = mode.xres * mode.yres * panel.gdfBytesPP

	/* Allocate framebuffer */
	// fb = memalign(ARCH_DMA_MINALIGN,
	// 	      roundup(panel.memSize, ARCH_DMA_MINALIGN));
	// if (!fb) {
	// 	printf("MXSFB: Error allocating framebuffer!\n");
	// 	return NULL;
	// }

	// /* Wipe framebuffer */
	// memset(fb, 0, panel.memSize);

	panel.frameAdrs = 0x80000000 // fb;

	// printf("%s\n", panel.modeIdent);

	/* Start framebuffer */
	mxs_lcd_init(ccm, panel, &mode, bpp)

	// var VIDEO_MXS_MODE_SYSTEM bool
	// if VIDEO_MXS_MODE_SYSTEM {
	// 	/*
	// 	 * If the LCD runs in system mode, the LCD refresh has to be triggered
	// 	 * manually by setting the RUN bit in HW_LCDIF_CTRL register. To avoid
	// 	 * having to set this bit manually after every single change in the
	// 	 * framebuffer memory, we set up specially crafted circular DMA, which
	// 	 * sets the RUN bit, then waits until it gets cleared and repeats this
	// 	 * infinitelly. This way, we get smooth continuous updates of the LCD.
	// 	 */
	// 	// var MXS_LCDIF_BASE uintptr

	// 	// memset(&desc, 0, sizeof(struct mxs_dma_desc));
	// 	// desc.address = (dma_addr_t)&desc;
	// 	// desc.cmd.data = MXS_DMA_DESC_COMMAND_NO_DMAXFER | MXS_DMA_DESC_CHAIN |
	// 	// 		MXS_DMA_DESC_WAIT4END |
	// 	// 		(1 << MXS_DMA_DESC_PIO_WORDS_OFFSET);
	// 	// desc.cmd.pio_words[0] = readl(hw_lcdif_ctrl) | LCDIF_CTRL_RUN;
	// 	// desc.cmd.next = (uint32_t)&desc.cmd;

	// 	// /* Execute the DMA chain. */
	// 	// mxs_dma_circ_start(MXS_DMA_CHANNEL_AHB_APBH_LCDIF, &desc);
	// }

	return nil
}
