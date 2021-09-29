package main

import (
	"flag"
	"log"
)

var (
	// These are used when run standalone not on the board.
	ccm = flag.String("ccm", "/dev/ccm", "Device to be used for the CCM")
)

func main() {
	flag.Parse()
	if err := NewLCD(true); err != nil {
		log.Printf("NewCLD: %v", err)
	}
	mii, err := NewMII()
	if err != nil {
		log.Printf("NewCLD: %v", err)
	}
	log.Printf("mii %v", mii)
}
