# PWM

**Arduino Uno:**

* `Timer0` - 8 bit timer for PD5 and PD6
* `Timer1` - 16 bit timer for PB1 and PB2
* `Timer2` - 8 bit timer for PB3 and PD3

For the two 8 bit timers, there is only a limited number of periods available, namely the CPU frequency divided by 256 and again divided by 1, 8, 64, 256, or 1024. For a MCU running at 16MHz, this would be a period of 16µs, 128µs, 1024µs, 4096µs, or 16384µs

## Links

* [tinygo.org: machine/arduino](https://tinygo.org/docs/reference/microcontrollers/machine/arduino/)
* [tinygo.org: examples/pwm](https://github.com/tinygo-org/tinygo/tree/release/src/examples/pwm)
* [tinygo.org: pwm.md](https://github.com/tinygo-org/tinygo-site/blob/pwm/content/microcontrollers/machine/pwm.md#using-pwm-in-tinygo)

 ___

Copyright Tauno Erik 2021 [taunoerik.art](https://taunoerik.art/)
