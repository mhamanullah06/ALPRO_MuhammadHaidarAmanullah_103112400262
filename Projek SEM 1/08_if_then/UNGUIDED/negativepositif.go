package main

import "fmt"

func main() {
	var bilangan int

	fmt.Print("masukan bilangan: ")
	fmt.Scan(&bilangan)

	if bilangan >= 0 {
		fmt.Printf("bukan")
	} else {
		fmt.Printf("genap negatif")
	}
}
