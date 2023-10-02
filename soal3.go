package main

import (
	"fmt"
)

func main() {
	var x int

	fmt.Print("masukkan nilai x (kurang dari sama dengan 999) = ")
	fmt.Scan(&x)

	if x < 0 || x > 999 {
		return
	}

	d1 := x / 100
	d2 := (x / 10) - ((x / 100) * 10)
	d3 := x % 10

	fmt.Printf("Hasil = %d\n", d1)
	fmt.Printf("Hasil = %d\n", d2)
	fmt.Printf("Hasil = %d\n", d3)

}
