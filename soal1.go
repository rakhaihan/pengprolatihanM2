package main

import (
	"fmt"
	"math"
)

func main() {

	var x, y float64

	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&y)

	hasil := (1/(3*math.Pow(x, 2)) + 10) + (10*y + 7)

	fmt.Printf("Hasil = %.8f\n", hasil)
}
