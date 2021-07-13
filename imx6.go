// https://github.com/f-secure-foundry/tamago-example
//
// Copyright (c) F-Secure Corporation
// https://foundry.f-secure.com
//
// Use of this source code is governed by the license
// that can be found in the LICENSE file.

// build mx6ullevk usbarmory

package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"log"
	"runtime"
	"syscall"

	"github.com/f-secure-foundry/tamago/soc/imx6"
)

const (
	romStart = 0x00000000
	romSize  = 0x17000
)

var i2c []*imx6.I2C

func init() {
	// The iomux device presents two files: iomuxctl and iomuxdata.
	// The differentiation is maintained in case we need semantics later.
	// For convenience, they are zero-relative, but this might change.
	if err := syscall.MkDev("/dev/iomuxdata", 0666, func() (syscall.DevFile, error) {
		return &longMemory{
			// This is how we'd do a relative address to IOMUXC_START,
			// how to do this tbd.
			//addr:   imx6.IOMUXC_START,
			//length: imx6.IOMUXC_END - imx6.IOMUXC_START + 1,
			adjust: 0,
			base:   0,
			length: imx6.IOMUXC_END + 1,
		}, nil
	}); err != nil {
		log.Printf("Can't set up iomux: %v", err)
	}
	if err := syscall.MkDev("/dev/iomuxctl", 0666, func() (syscall.DevFile, error) {
		return &longMemory{
			adjust: 0,
			base:   0,
			length: imx6.IOMUXC_END + 1,
		}, nil
	}); err != nil {
		log.Printf("Can't set up iomux: %v", err)
	}
	// TODO: just make this nonsense relative to CCMGR
	if err := syscall.MkDev("/dev/ccm", 0666, func() (syscall.DevFile, error) {
		return &longMemory{
			adjust: 0,
			base:   0,
			length: 0x2100000, //imx6.CCM_CCGR6 + 4,
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
	res.WriteString(fmt.Sprintf("SoC ..........: %s %d MHz\n", imx6.Model(), imx6.ARMFreq()/1000000))
	res.WriteString(fmt.Sprintf("Secure boot ..: %v\n", imx6.SNVS()))
	res.WriteString(fmt.Sprintf("Boot ROM hash : %#x\n", sha256.Sum256(rom)))
	//	res.WriteString(fmt.Sprintf("IOMUX spans from : %#x - %#x\n", imx6.IOMUXC_START, imx6.IOMUXC_END-imx6.IOMUXC_START+1))
	res.WriteString(fmt.Sprintf("IOMUX spans from : %#x - %#x\n", 0, imx6.IOMUXC_END+1))

	return res.String()
}
