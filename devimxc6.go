// https://github.com/usbarmory/tamago-example
//
// Copyright (c) WithSecure Corporation
// https://foundry.withsecure.com
//
// Use of this source code is governed by the license
// that can be found in the LICENSE file.

//go:build mx6ullevk || usbarmory
// +build mx6ullevk usbarmory

package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"log"
	"runtime"
	"syscall"

	"github.com/usbarmory/tamago/soc/nxp/imx6ul"
	"github.com/usbarmory/tamago/soc/nxp/usb"
)

const (
	romStart = 0x00000000
	romSize  = 0x17000
)

//var i2c []*imx6ul.I2C

func init() {
	// muxclone creates a new mux instance, unitialized.
	// if err := syscall.MkDev("/dev/muxclone", 0666, func() (syscall.DevFile, error) {
	// 	return &longMemory{
	// 		// This is how we'd do a relative address to IOMUXC_START,
	// 		// how to do this tbd.
	// 		//addr:   imx6ul.IOMUXC_START,
	// 		//length: imx6ul.IOMUXC_END - imx6ul.IOMUXC_START + 1,
	// 		adjust: 0,
	// 		base:   0,
	// 		length: imx6ul.IOMUXC_END + 1,
	// 	}, nil
	// }); err != nil {
	// 	log.Printf("Can't set up iomux: %v", err)
	// }
	// if err := syscall.MkDev("/dev/iomuxctl", 0666, func() (syscall.DevFile, error) {
	// 	return &longMemory{
	// 		adjust: 0,
	// 		base:   0,
	// 		length: imx6ul.IOMUXC_END + 1,
	// 	}, nil
	// }); err != nil {
	// 	log.Printf("Can't set up iomux: %v", err)
	// }
	// TODO: just make this nonsense relative to CCMGR
	if err := syscall.MkDev("/dev/ccm", 0666, func() (syscall.DevFile, error) {
		return &longMemory{
			adjust: 0,
			base:   0,
			length: 0x2200000, //imx6ul.CCM_CCGR6 + 4,
		}, nil
	}); err != nil {
		log.Printf("Can't set up iomux: %v", err)
	}

}

func info() string {
	var res bytes.Buffer

	rom := mem(romStart, romSize, nil)

	res.WriteString(fmt.Sprintf("Runtime ......: %s %s/%s\n", runtime.Version(), runtime.GOOS, runtime.GOARCH))
	res.WriteString(fmt.Sprintf("Board ........: %s\n", boardName))
	res.WriteString(fmt.Sprintf("SoC ..........: %s %d MHz\n", imx6ul.Model(), imx6ul.ARMFreq()/1000000))
	res.WriteString(fmt.Sprintf("SDP ..........: %v\n", usb.SDP()))
	res.WriteString(fmt.Sprintf("Secure boot ..: %v\n", imx6ul.SNVS()))
	res.WriteString(fmt.Sprintf("Boot ROM hash : %x\n", sha256.Sum256(rom)))

	return res.String()
}
