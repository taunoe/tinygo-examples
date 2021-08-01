package main

import (
	"machine"
	"time"
)

func main() {
	delay := func(t uint16) { // implicit function
		time.Sleep(time.Duration(1000000 * uint32(t)))
	}

	// LED pins
	leds := []machine.Pin{
		machine.D13,
		machine.D12,
		machine.D11,
		machine.D10,
		machine.D9,
		machine.D8,
	}

	// Pins as output
	for _, led := range leds {
		led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	}

	// Variables
	index, delta := 0, 1

	// Loop
	for {
		for i, led := range leds {
			led.Set(i == index) // In Go you cannot convert int to bool; you'll have to use logical operators.
		}

		index += delta

		if index == 0 || index == len(leds)-1 {
			delta *= -1
		}

		delay(75)
	}
}
