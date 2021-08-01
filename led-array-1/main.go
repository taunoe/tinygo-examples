package main

import (
	"machine"
	"time"
)

func delay(t uint16) {
	time.Sleep(time.Duration(1000000 * uint32(t)))
}

func main() {

	// LED pins
	leds := []machine.Pin{
		machine.D13,
		machine.D12,
		machine.D11,
		machine.D10,
		machine.D9,
		machine.D8,
	}

	// Set LED pins output mode
	for i := 0; i < len(leds); i++ {
		leds[i].Configure(machine.PinConfig{Mode: machine.PinOutput})
	}

	// Loop
	for {
		// Up
		for i := 0; i < len(leds); i++ {
			leds[i].High()
			delay(75)
			leds[i].Low()
		}
		// Down
		for i := len(leds) - 1; i >= 0; i-- {
			leds[i].High()
			delay(75)
			leds[i].Low()
		}

	}
}
