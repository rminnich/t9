package main

/* SPDX-License-Identifier: GPL-2.0+ */
/*
 * Freescale i.MX28/6SX/6UL/7D LCDIF Register Definitions
 *
 * Copyright (C) 2011 Marek Vasut <marek.vasut@gmail.com>
 * on behalf of DENX Software Engineering GmbH
 *
 * Based on code from LTIB:
 * Copyright 2008-2010 Freescale Semiconductor, Inc. All Rights Reserved.
 */

const (
	hw_lcdif_base           = 0x21C_8000
	set                     = 4
	clr                     = 8
	tog                     = 12
	hw_lcdif_ctrl           = hw_lcdif_base + 0x00
	hw_lcdif_ctrl1          = hw_lcdif_base + 0x10
	hw_lcdif_ctrl2          = hw_lcdif_base + 0x20
	hw_lcdif_transfer_count = hw_lcdif_base + 0x30
	hw_lcdif_cur_buf        = hw_lcdif_base + 0x40
	hw_lcdif_next_buf       = hw_lcdif_base + 0x50

	hw_lcdif_timing         = hw_lcdif_base + 0x60
	hw_lcdif_vdctrl0        = hw_lcdif_base + 0x70
	hw_lcdif_vdctrl1        = hw_lcdif_base + 0x80
	hw_lcdif_vdctrl2        = hw_lcdif_base + 0x90
	hw_lcdif_vdctrl3        = hw_lcdif_base + 0xa0
	hw_lcdif_vdctrl4        = hw_lcdif_base + 0xb0
	hw_lcdif_dvictrl0       = hw_lcdif_base + 0xc0
	hw_lcdif_dvictrl1       = hw_lcdif_base + 0xd0
	hw_lcdif_dvictrl2       = hw_lcdif_base + 0xe0
	hw_lcdif_dvictrl3       = hw_lcdif_base + 0xf0
	hw_lcdif_dvictrl4       = hw_lcdif_base + 0x100
	hw_lcdif_csc_coeffctrl0 = hw_lcdif_base + 0x110
	hw_lcdif_csc_coeffctrl1 = hw_lcdif_base + 0x120
	hw_lcdif_csc_coeffctrl2 = hw_lcdif_base + 0x130
	hw_lcdif_csc_coeffctrl3 = hw_lcdif_base + 0x140
	hw_lcdif_csc_coeffctrl4 = hw_lcdif_base + 0x150
	hw_lcdif_csc_offset     = hw_lcdif_base + 0x160
	hw_lcdif_csc_limit      = hw_lcdif_base + 0x170

	hw_lcdif_data          = hw_lcdif_base + 0x180
	hw_lcdif_bm_error_stat = hw_lcdif_base + 0x190
	hw_lcdif_crc_stat      = hw_lcdif_base + 0x1a0
	hw_lcdif_lcdif_stat    = hw_lcdif_base + 0x1b0
	hw_lcdif_version       = hw_lcdif_base + 0x1c0
	hw_lcdif_debug0        = hw_lcdif_base + 0x1d0
	hw_lcdif_debug1        = hw_lcdif_base + 0x1e0
	hw_lcdif_debug2        = hw_lcdif_base + 0x1f0
)

const (
	LCDIF_CTRL_SFTRST                    = 1 << 31
	LCDIF_CTRL_CLKGATE                   = 1 << 30
	LCDIF_CTRL_YCBCR422_INPUT            = 1 << 29
	LCDIF_CTRL_READ_WRITEB               = 1 << 28
	LCDIF_CTRL_WAIT_FOR_VSYNC_EDGE       = 1 << 27
	LCDIF_CTRL_DATA_SHIFT_DIR            = 1 << 26
	LCDIF_CTRL_SHIFT_NUM_BITS_MASK       = 0x1f << 21
	LCDIF_CTRL_SHIFT_NUM_BITS_OFFSET     = 21
	LCDIF_CTRL_DVI_MODE                  = 1 << 20
	LCDIF_CTRL_BYPASS_COUNT              = 1 << 19
	LCDIF_CTRL_VSYNC_MODE                = 1 << 18
	LCDIF_CTRL_DOTCLK_MODE               = 1 << 17
	LCDIF_CTRL_DATA_SELECT               = 1 << 16
	LCDIF_CTRL_INPUT_DATA_SWIZZLE_MASK   = 0x3 << 14
	LCDIF_CTRL_INPUT_DATA_SWIZZLE_OFFSET = 14
	LCDIF_CTRL_CSC_DATA_SWIZZLE_MASK     = 0x3 << 12
	LCDIF_CTRL_CSC_DATA_SWIZZLE_OFFSET   = 12
	LCDIF_CTRL_LCD_DATABUS_WIDTH_MASK    = 0x3 << 10
	LCDIF_CTRL_LCD_DATABUS_WIDTH_OFFSET  = 10
	LCDIF_CTRL_LCD_DATABUS_WIDTH_16BIT   = 0 << 10
	LCDIF_CTRL_LCD_DATABUS_WIDTH_8BIT    = 1 << 10
	LCDIF_CTRL_LCD_DATABUS_WIDTH_18BIT   = 2 << 10
	LCDIF_CTRL_LCD_DATABUS_WIDTH_24BIT   = 3 << 10
	LCDIF_CTRL_WORD_LENGTH_MASK          = 0x3 << 8
	LCDIF_CTRL_WORD_LENGTH_OFFSET        = 8
	LCDIF_CTRL_WORD_LENGTH_16BIT         = 0 << 8
	LCDIF_CTRL_WORD_LENGTH_8BIT          = 1 << 8
	LCDIF_CTRL_WORD_LENGTH_18BIT         = 2 << 8
	LCDIF_CTRL_WORD_LENGTH_24BIT         = 3 << 8
	LCDIF_CTRL_RGB_TO_YCBCR422_CSC       = 1 << 7
	LCDIF_CTRL_LCDIF_MASTER              = 1 << 5
	LCDIF_CTRL_DATA_FORMAT_16_BIT        = 1 << 3
	LCDIF_CTRL_DATA_FORMAT_18_BIT        = 1 << 2
	LCDIF_CTRL_DATA_FORMAT_24_BIT        = 1 << 1
	LCDIF_CTRL_RUN                       = 1 << 0

	LCDIF_CTRL1_COMBINE_MPU_WR_STRB               = 1 << 27
	LCDIF_CTRL1_BM_ERROR_IRQ_EN                   = 1 << 26
	LCDIF_CTRL1_BM_ERROR_IRQ                      = 1 << 25
	LCDIF_CTRL1_RECOVER_ON_UNDERFLOW              = 1 << 24
	LCDIF_CTRL1_INTERLACE_FIELDS                  = 1 << 23
	LCDIF_CTRL1_START_INTERLACE_FROM_SECOND_FIELD = 1 << 22
	LCDIF_CTRL1_FIFO_CLEAR                        = 1 << 21
	LCDIF_CTRL1_IRQ_ON_ALTERNATE_FIELDS           = 1 << 20
	LCDIF_CTRL1_BYTE_PACKING_FORMAT_MASK          = 0xf << 16
	LCDIF_CTRL1_BYTE_PACKING_FORMAT_OFFSET        = 16
	LCDIF_CTRL1_OVERFLOW_IRQ_EN                   = 1 << 15
	LCDIF_CTRL1_UNDERFLOW_IRQ_EN                  = 1 << 14
	LCDIF_CTRL1_CUR_FRAME_DONE_IRQ_EN             = 1 << 13
	LCDIF_CTRL1_VSYNC_EDGE_IRQ_EN                 = 1 << 12
	LCDIF_CTRL1_OVERFLOW_IRQ                      = 1 << 11
	LCDIF_CTRL1_UNDERFLOW_IRQ                     = 1 << 10
	LCDIF_CTRL1_CUR_FRAME_DONE_IRQ                = 1 << 9
	LCDIF_CTRL1_VSYNC_EDGE_IRQ                    = 1 << 8
	LCDIF_CTRL1_BUSY_ENABLE                       = 1 << 2
	LCDIF_CTRL1_MODE86                            = 1 << 1
	LCDIF_CTRL1_RESET                             = 1 << 0

	LCDIF_CTRL2_OUTSTANDING_REQS_MASK                = 0x7 << 21
	LCDIF_CTRL2_OUTSTANDING_REQS_OFFSET              = 21
	LCDIF_CTRL2_OUTSTANDING_REQS_REQ_1               = 0x0 << 21
	LCDIF_CTRL2_OUTSTANDING_REQS_REQ_2               = 0x1 << 21
	LCDIF_CTRL2_OUTSTANDING_REQS_REQ_4               = 0x2 << 21
	LCDIF_CTRL2_OUTSTANDING_REQS_REQ_8               = 0x3 << 21
	LCDIF_CTRL2_OUTSTANDING_REQS_REQ_16              = 0x4 << 21
	LCDIF_CTRL2_BURST_LEN_8                          = 1 << 20
	LCDIF_CTRL2_ODD_LINE_PATTERN_MASK                = 0x7 << 16
	LCDIF_CTRL2_ODD_LINE_PATTERN_OFFSET              = 16
	LCDIF_CTRL2_ODD_LINE_PATTERN_RGB                 = 0x0 << 16
	LCDIF_CTRL2_ODD_LINE_PATTERN_RBG                 = 0x1 << 16
	LCDIF_CTRL2_ODD_LINE_PATTERN_GBR                 = 0x2 << 16
	LCDIF_CTRL2_ODD_LINE_PATTERN_GRB                 = 0x3 << 16
	LCDIF_CTRL2_ODD_LINE_PATTERN_BRG                 = 0x4 << 16
	LCDIF_CTRL2_ODD_LINE_PATTERN_BGR                 = 0x5 << 16
	LCDIF_CTRL2_EVEN_LINE_PATTERN_MASK               = 0x7 << 12
	LCDIF_CTRL2_EVEN_LINE_PATTERN_OFFSET             = 12
	LCDIF_CTRL2_EVEN_LINE_PATTERN_RGB                = 0x0 << 12
	LCDIF_CTRL2_EVEN_LINE_PATTERN_RBG                = 0x1 << 12
	LCDIF_CTRL2_EVEN_LINE_PATTERN_GBR                = 0x2 << 12
	LCDIF_CTRL2_EVEN_LINE_PATTERN_GRB                = 0x3 << 12
	LCDIF_CTRL2_EVEN_LINE_PATTERN_BRG                = 0x4 << 12
	LCDIF_CTRL2_EVEN_LINE_PATTERN_BGR                = 0x5 << 12
	LCDIF_CTRL2_READ_PACK_DIR                        = 1 << 10
	LCDIF_CTRL2_READ_MODE_OUTPUT_IN_RGB_FORMAT       = 1 << 9
	LCDIF_CTRL2_READ_MODE_6_BIT_INPUT                = 1 << 8
	LCDIF_CTRL2_READ_MODE_NUM_PACKED_SUBWORDS_MASK   = 0x7 << 4
	LCDIF_CTRL2_READ_MODE_NUM_PACKED_SUBWORDS_OFFSET = 4
	LCDIF_CTRL2_INITIAL_DUMMY_READ_MASK              = 0x7 << 1
	LCDIF_CTRL2_INITIAL_DUMMY_READ_OFFSET            = 1

	LCDIF_TRANSFER_COUNT_V_COUNT_MASK   = 0xffff << 16
	LCDIF_TRANSFER_COUNT_V_COUNT_OFFSET = 16
	LCDIF_TRANSFER_COUNT_H_COUNT_MASK   = 0xffff << 0
	LCDIF_TRANSFER_COUNT_H_COUNT_OFFSET = 0

	LCDIF_CUR_BUF_ADDR_MASK   = 0xffffffff
	LCDIF_CUR_BUF_ADDR_OFFSET = 0

	LCDIF_NEXT_BUF_ADDR_MASK   = 0xffffffff
	LCDIF_NEXT_BUF_ADDR_OFFSET = 0

	LCDIF_TIMING_CMD_HOLD_MASK     = 0xff << 24
	LCDIF_TIMING_CMD_HOLD_OFFSET   = 24
	LCDIF_TIMING_CMD_SETUP_MASK    = 0xff << 16
	LCDIF_TIMING_CMD_SETUP_OFFSET  = 16
	LCDIF_TIMING_DATA_HOLD_MASK    = 0xff << 8
	LCDIF_TIMING_DATA_HOLD_OFFSET  = 8
	LCDIF_TIMING_DATA_SETUP_MASK   = 0xff << 0
	LCDIF_TIMING_DATA_SETUP_OFFSET = 0

	LCDIF_VDCTRL0_VSYNC_OEB                = 1 << 29
	LCDIF_VDCTRL0_ENABLE_PRESENT           = 1 << 28
	LCDIF_VDCTRL0_VSYNC_POL                = 1 << 27
	LCDIF_VDCTRL0_HSYNC_POL                = 1 << 26
	LCDIF_VDCTRL0_DOTCLK_POL               = 1 << 25
	LCDIF_VDCTRL0_ENABLE_POL               = 1 << 24
	LCDIF_VDCTRL0_VSYNC_PERIOD_UNIT        = 1 << 21
	LCDIF_VDCTRL0_VSYNC_PULSE_WIDTH_UNIT   = 1 << 20
	LCDIF_VDCTRL0_HALF_LINE                = 1 << 19
	LCDIF_VDCTRL0_HALF_LINE_MODE           = 1 << 18
	LCDIF_VDCTRL0_VSYNC_PULSE_WIDTH_MASK   = 0x3ffff
	LCDIF_VDCTRL0_VSYNC_PULSE_WIDTH_OFFSET = 0

	LCDIF_VDCTRL1_VSYNC_PERIOD_MASK   = 0xffffffff
	LCDIF_VDCTRL1_VSYNC_PERIOD_OFFSET = 0

	LCDIF_VDCTRL2_HSYNC_PULSE_WIDTH_MASK   = 0x3fff << 18
	LCDIF_VDCTRL2_HSYNC_PULSE_WIDTH_OFFSET = 18
	LCDIF_VDCTRL2_HSYNC_PERIOD_MASK        = 0x3ffff
	LCDIF_VDCTRL2_HSYNC_PERIOD_OFFSET      = 0

	LCDIF_VDCTRL3_MUX_SYNC_SIGNALS           = 1 << 29
	LCDIF_VDCTRL3_VSYNC_ONLY                 = 1 << 28
	LCDIF_VDCTRL3_HORIZONTAL_WAIT_CNT_MASK   = 0xfff << 16
	LCDIF_VDCTRL3_HORIZONTAL_WAIT_CNT_OFFSET = 16
	LCDIF_VDCTRL3_VERTICAL_WAIT_CNT_MASK     = 0xffff << 0
	LCDIF_VDCTRL3_VERTICAL_WAIT_CNT_OFFSET   = 0

	LCDIF_VDCTRL4_DOTCLK_DLY_SEL_MASK            = 0x7 << 29
	LCDIF_VDCTRL4_DOTCLK_DLY_SEL_OFFSET          = 29
	LCDIF_VDCTRL4_SYNC_SIGNALS_ON                = 1 << 18
	LCDIF_VDCTRL4_DOTCLK_H_VALID_DATA_CNT_MASK   = 0x3ffff
	LCDIF_VDCTRL4_DOTCLK_H_VALID_DATA_CNT_OFFSET = 0
)

const (
	GDF__8BIT_INDEX   = 0
	GDF_15BIT_555RGB  = 1
	GDF_16BIT_565RGB  = 2
	GDF_32BIT_X888RGB = 3
	GDF_24BIT_888RGB  = 4
	GDF__8BIT_332RGB  = 5
)
