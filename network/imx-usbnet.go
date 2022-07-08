// https://github.com/usbarmory/tamago-example
//
// Copyright (c) WithSecure Corporation
// https://foundry.withsecure.com
//
// Use of this source code is governed by the license
// that can be found in the LICENSE file.

//go:build mx6ullevk || usbarmory
// +build mx6ullevk usbarmory

package network

import (
	"log"
	"os"

	usbnet "github.com/usbarmory/imx-usbnet"
	"github.com/usbarmory/tamago/soc/imx6/usb"

	"github.com/miekg/dns"
)

const (
	deviceIP  = "10.0.0.1"
	deviceMAC = "1a:55:89:a2:69:41"
	hostMAC   = "1a:55:89:a2:69:42"
	resolver  = "8.8.8.8:53"
)

var iface *usbnet.Interface

func Start(journalFile *os.File) {
	var err error

	iface, err = usbnet.Init(deviceIP, deviceMAC, hostMAC, 1)

	if err != nil {
		log.Fatalf("could not initialize USB networking, %v", err)
	}

	iface.EnableICMP()

	listenerSSH, err := iface.ListenerTCP4(22)

	if err != nil {
		log.Fatalf("could not initialize SSH listener, %v", err)
	}

	listenerHTTP, err := iface.ListenerTCP4(80)

	if err != nil {
		log.Fatalf("could not initialize HTTP listener, %v", err)
	}

	listenerHTTPS, err := iface.ListenerTCP4(443)

	if err != nil {
		log.Fatalf("could not initialize HTTP listener, %v", err)
	}

	listener9P, err := iface.ListenerTCP4(564)

	if err != nil {
		log.Printf("could not initialize 9P listener, %v", err)
	}

	// create index.html
	setupStaticWebAssets()

	journal = journalFile

	go func() {
		startSSHServer(listenerSSH, deviceIP, 22)
	}()

	go func() {
		startWebServer(listenerHTTP, deviceIP, 80, false)
	}()

	go func() {
		startWebServer(listenerHTTPS, deviceIP, 443, true)
	}()

	// 9p server (see 9p_server.go)
	go func() {
		start9pServer(listener9P, deviceIP, 564, 1)
	}()

	usb.USB1.Init()
	usb.USB1.DeviceMode()
	usb.USB1.Reset()

	// never returns
	imx6ul.USB1.Start(iface.Device())
}
