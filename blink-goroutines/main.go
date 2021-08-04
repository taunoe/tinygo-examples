package main

import (
	"machine"
	"time"
)

func led1() {
	led := machine.D11
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		led.High()
		delay(500)
		led.Low()
		delay(500)
	}
}

func led2() {
	led := machine.D10
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		led.High()
		delay(400)
		led.Low()
		delay(400)
	}
}

func delay(t int64) {
	time.Sleep(time.Duration(1000000 * t))
	//time.Sleep(time.Millisecond * t)
}

func main() {
	go led1() // goroutine
	led2()
}
