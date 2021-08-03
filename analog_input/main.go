package main

import (
	"machine"
	"time"
)

func main() {
	machine.InitADC()
	ldr := machine.ADC{machine.ADC0}
	ldr.Configure(machine.ADCConfig{})

	led := machine.Pin(machine.D9)
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		// Serial print baudrate 9600
		print(ldr.Get()) // Analogread
		print("\n")

		if ldr.Get() > 40000 {
			led.Set(true)
		} else {
			led.Set(false)
		}

		time.Sleep(time.Millisecond * 100)
	}
}
