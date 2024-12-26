package main

import "fmt"

func main() {

	var bb, tb, bmi float32
	fmt.Println("masukan berat anda: ")
	fmt.Scanln(&bb)

	fmt.Println("masukan tinggi anda: ")
	fmt.Scanln(&tb)

	bmi = bb / (tb * tb)
	fmt.Printf("%.2f", bmi)
}
