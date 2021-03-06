package main

import (
	"machine"
	"time"
)

const (
	led    = machine.D9 //machine.LED
	button = machine.D2 //machine.BUTTON
)

func main() {
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	button.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	for {
		if button.Get() {
			led.Low()
		} else {
			led.High()
		}

		time.Sleep(time.Millisecond * 10)
	}
}
