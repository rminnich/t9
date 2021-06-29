package main

import "log"

func main() {
	if err := NewLCD(true); err != nil {
		log.Printf("NewCLD: %v", err)
	}
}
