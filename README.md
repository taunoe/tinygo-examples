# TinyGo examples

Simple experiments with TinyGo.

## Defining Pins

    button := machine.D2 // Pin D2
    led := machine.D9  // Pin D9

**Input mode:**

    button.Configure(machine.PinConfig{Mode: machine.PinInput})

**Output mode:**

    led.Configure(machine.PinConfig{Mode: machine.PinOutput})

**Internal pull-up resistor:**

    button.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

## Methods

**Digital read** pin with `.Get()`. Returns bool (true/false).

    button.Get()

**Digital write** with `.Set()` method or `.High()` and `.Low()`.

    led.Set(true) // High
    led.Low()
    led.High()

## Serial print

Baudrate 9600. How to change?

    print("Message\n")
    println("Message")

## Time

    time.Second
    time.Millisecond

## Init modul

On folder create Go module:

    go mod init

Ad used depencies to `go.mod` file:

    go mod tidy

## Flash to board

Uploading to Arduino Uno:

    tinygo flash -target=arduino

Specifiying the port and path if needed:

    tinygo flash -target=arduino -port=/dev/ttyACM0 path-to/main.go

### Targets

For Arduino Nano and Arduino Pro Mini use:

    -target=arduino-nano

## Build separate file

    tinygo build --target=arduino -o outputfile path-to/main.go

## Example codes

* [Blink one LED](./blink/)
* [Toggle two LEDs](./blink-2/)
* [Toggle two LEDs with goroutine](./blink-goroutines/)
* [Knight Rider effect 1, LED array](./led-array-1/)
* [Knight Rider effect 2, LED array](./led-array-2/)
* [Button/Digital input](./digital-input/) Button with 10K pullup resistor.
* [Button/Digital input](./button/) With internal pullup resistor.
* [PWM - Pulse Width Modulation](./PWM/) Analog output
* [Analog input, ADC](./analog_input/)
* [Pininterrupt](./interrupt/) Arduino not supported??
* [UART0](./uart/)

## Links

* [TinyGo on Arduino Uno: An Introduction](https://create.arduino.cc/projecthub/alankrantas/tinygo-on-arduino-uno-an-introduction-6130f6) - not updated
* [Creative DIY Microcontroller Projects with TinyGo and WebAssembly](https://github.com/PacktPublishing/Creative-DIY-Microcontroller-Projects-with-TinyGo-and-WebAssembly)
* [TinyGo.org examples](https://github.com/tinygo-org/tinygo/tree/release/src/examples)
* [Goroutines in TinyGo](https://aykevl.nl/2019/02/tinygo-goroutines)
* [TinyGo Drivers](https://github.com/tinygo-org/drivers)

 ___

Copyright Tauno Erik 2021 [taunoerik.art](https://taunoerik.art/)
