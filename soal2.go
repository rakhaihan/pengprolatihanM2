package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64

	fmt.Print("nilai x: ")
	fmt.Scan(&x)

	hasil := ((1 * math.Pow(x, 3)) + (3*x)/(1*math.Pow(x, 4)) - (3 * math.Pow(x, 2)) + 4)

	fmt.Printf("hasil = %f\n", hasil)
}
