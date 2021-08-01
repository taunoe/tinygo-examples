// 01.08.2021 Tauno erik
package main

import (
	"machine"
	"time"
)

func main() {
	// Define pin
	led := machine.LED // On board LED

	// Init pin
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	/* // longer version:
	   config := machine.PinConfig{
	       Mode: machine.PinOutput
	   }
	   led.Configure(config)
	*/

	// Loop
	for {
		led.Low()
		time.Sleep(time.Millisecond * 500)

		led.High()
		time.Sleep(time.Millisecond * 500)
	}
}
