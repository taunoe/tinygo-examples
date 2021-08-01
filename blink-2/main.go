// Toggle two LEDs
package main

import (
	"machine"
	"time"
)

func delay(t uint32) {
	time.Sleep(time.Duration(1000000 * t))
}

func main() {
	// Setup
	led_1 := machine.D12
	led_2 := machine.D11

	led_1.Configure(machine.PinConfig{Mode: machine.PinOutput})
	led_2.Configure(machine.PinConfig{Mode: machine.PinOutput})

	led_switch_1 := true
	led_switch_2 := false

	// Loop
	for {
		led_1.Set(led_switch_1)
		led_2.Set(led_switch_2)

		led_switch_1 = !led_switch_1
		println("led 1", led_switch_1)
		led_switch_2 = !led_switch_2
		println("led 2", led_switch_2)

		delay(500)
	}

}
