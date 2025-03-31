package main //pointers

import (
	"fmt"
)

func main() {
	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\n memory address of thing1 is %p", &thing1)
	var result [5]float64 = square(thing1)
	fmt.Printf("\n The result is :%v", result)
	fmt.Printf("\n The original is :%v", thing1)

}

func square(thing2 [5]float64) [5]float64 {
	fmt.Printf("\n memory address of thing2 is %p", &thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]

	}
	return thing2
}
