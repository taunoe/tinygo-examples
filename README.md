# TinyGo examples

Simple experiments with TinyGo.

## Defining Pins

    button := machine.D2 // Pin D2
    led := machine.D9  // Pin D9

**Input mode:**

    button.Configure(machine.PinConfig{Mode: machine.PinInput})

**Output mode:**

    led.Configure(machine.PinConfig{Mode: machine.PinOutput})

## Methods

**Digital read** pin with `.Get()`. Returns bool (true/false).

    button.Get()

**Digital write** with `.Set()` method or `.High()` and `.Low`.

    led.Set(true) // High
    led.Low()
    led.High()

## Example code

* [Blink one LED](./blink/)
* [Toggle two LEDs](./blink-2/)
* [Toggle two LEDs with goroutine](./blink-goroutines/main.go)
* [Knight Rider effect 1, LED array](./led-array-1/main.go)
* [Knight Rider effect 2, LED array](./led-array-2/main.go)
* [Digital input](./digital-input/main.go) Buttons with 10K pullup resistor.

## Links

* [TinyGo on Arduino Uno: An Introduction](https://create.arduino.cc/projecthub/alankrantas/tinygo-on-arduino-uno-an-introduction-6130f6)

 ___

Copyright Tauno Erik 2021 [taunoerik.art](https://taunoerik.art/)
