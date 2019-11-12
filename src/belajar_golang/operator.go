package main

import (
	"fmt"
)

func main() {
	var value = ((2 + 6) / 2)
	var isTrue = (value == 4)
	var isFalse = (value == 2)

	fmt.Printf("Nilai %d (%t) \n", value, isTrue)  //%t untuk memunculkan nilai bool
	fmt.Printf("Nilai %d (%t) \n", value, isFalse) //%t untuk memunculkan nilai bool
}
