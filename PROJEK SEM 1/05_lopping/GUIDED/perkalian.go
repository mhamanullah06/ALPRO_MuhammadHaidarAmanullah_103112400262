package main

import (
	"fmt"
)

func main() {
	var i, a, b, hasil int
	fmt.Scan(&a, &b)
	hasil = 0
	for i = 1; i <= b; i++ {
		hasil = hasil + a
	}
	fmt.Println("Hasil", hasil)
}
