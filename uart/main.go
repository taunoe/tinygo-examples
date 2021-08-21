package main

import (
	"machine"
	"time"
)

// Delay function
func delay(delay uint16) {
	time.Sleep(time.Duration(1000000 * uint32(delay)))
}

func main() {
	for {
		// Read USB Serial
		in_byte, err := machine.UART0.ReadByte()

		if err != nil {
			// Prints UART bufer empty
			// println(err)
		}

		if in_byte == 'a' {
			println("in_byte == a")
		} else if in_byte == 'b' {
			// Write USB Serial
			machine.UART0.WriteByte('a')
			machine.UART0.WriteByte('\n')

			data := []byte("See on byte[] array\n")
			machine.UART0.Write(data)
		}

		//delay(100)
	}
}
