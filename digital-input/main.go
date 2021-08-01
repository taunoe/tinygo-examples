package main

import (
	"machine"
	"time"
)

func main() {

	// Define button pin
	button := machine.D2
	// Init button pin
	button.Configure(machine.PinConfig{Mode: machine.PinInput})

	// Define LED pin
	led := machine.D9
	// Init LED pin
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// Loop
	for {
		//.Get() method returns bool value
		led.Set(!button.Get())
		time.Sleep(time.Millisecond * 100)
	}
}
