package main

import (
	"machine"
	"time"
)

const delayBetweenPeriods = time.Second * 5

var (
	// Arduino Uno analog output pins: 3,5,6,9,10,11
	pwm  = machine.Timer2
	pinA = machine.D3
	pinB = machine.D11
)

func main() {

	// Delay a bit on startup to easily catch the first messages.
	time.Sleep(time.Second * 2)

	// Configure the PWM with the given period.
	err := pwm.Configure(machine.PWMConfig{
		Period: 16384e3, // 16.384ms
	})

	if err != nil {
		println("failed to configure PWM")
		return
	}
	// The top value is the highest value that can be passed to PWMChannel.Set.
	// It is usually an even number.
	println("top:", pwm.Top())

	// Configure the two channels we'll use as outputs.
	channelA, err := pwm.Channel(pinA)
	if err != nil {
		println("failed to configure channel A")
		println(err)
		return
	}
	channelB, err := pwm.Channel(pinB)
	if err != nil {
		println("failed to configure channel B")
		println(err)
		return
	}

	// Invert one of the channels to demonstrate output polarity.
	pwm.SetInverting(channelB, true)

	// Test out various frequencies below, including some edge cases.

	println("running at 0% duty cycle")
	pwm.Set(channelA, 0)
	pwm.Set(channelB, 0)
	time.Sleep(delayBetweenPeriods)

	println("running at 1")
	pwm.Set(channelA, 1)
	pwm.Set(channelB, 1)
	time.Sleep(delayBetweenPeriods)

	println("running at 25% duty cycle")
	pwm.Set(channelA, pwm.Top()/4)
	pwm.Set(channelB, pwm.Top()/4)
	time.Sleep(delayBetweenPeriods)

	println("running at top-1")
	pwm.Set(channelA, pwm.Top()-1)
	pwm.Set(channelB, pwm.Top()-1)
	time.Sleep(delayBetweenPeriods)

	println("running at 100% duty cycle")
	pwm.Set(channelA, pwm.Top())
	pwm.Set(channelB, pwm.Top())
	time.Sleep(delayBetweenPeriods)

	for {
		for i := 0; i < 20; i++ {
			pwm.Set(channelA, pwm.Top()/uint32(i))
			pwm.Set(channelB, pwm.Top()/uint32(i))
			time.Sleep(50 * time.Millisecond)
		}

		for i := 20; i > 0; i-- {
			pwm.Set(channelA, pwm.Top()/uint32(i))
			pwm.Set(channelB, pwm.Top()/uint32(i))
			time.Sleep(50 * time.Millisecond)
		}
	}
}
