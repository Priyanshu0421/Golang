package main

import (
	"fmt"
	"math"
)

func main() {
	pi := math.Pi
	// val := math.Abs(-34)  // it will convert the negative value to the positive value
	// fmt.Println(val)
	// Acos -> Acos returns the arccosine, in radians, of x.
	// val := math.Acos(0) // speacial case will print 1.5707963267948966
	val := math.Acos(1) // it will print 0
	fmt.Println(val)
	fmt.Println("The value of the pi in golang is: ", pi)
}
