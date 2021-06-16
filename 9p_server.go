// UFS is a userspace server which exports a filesystem over 9p2000.
//
// By default, it will export / over a TCP on port 5640 under the username
// of "harvey".
package main

import (
	"log"

	ufs "github.com/Harvey-OS/ninep/filesystem"
	"github.com/Harvey-OS/ninep/protocol"
	"gvisor.dev/gvisor/pkg/tcpip"
	"gvisor.dev/gvisor/pkg/tcpip/adapters/gonet"
	"gvisor.dev/gvisor/pkg/tcpip/network/ipv4"
	"gvisor.dev/gvisor/pkg/tcpip/stack"
)

func start9pServer(s *stack.Stack, addr tcpip.Address, port uint16, nic tcpip.NICID) {
	var err error

	fullAddr := tcpip.FullAddress{Addr: addr, Port: port, NIC: nic}
	ln, err := gonet.ListenTCP(s, fullAddr, ipv4.ProtocolNumber)

	if err != nil {
		log.Fatal("9p listener error: ", err)
	}

	ufslistener, err := ufs.NewUFS(func(l *protocol.Listener) error {
		l.Trace = log.Printf
		return nil
	})

	if err := ufslistener.Serve(ln); err != nil {
		log.Fatal(err)
	}
}
