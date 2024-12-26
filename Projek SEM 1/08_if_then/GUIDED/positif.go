package main

import "fmt"

func main() {
	var bilangan int

	fmt.Print("masukan bilangan: ")
	fmt.Scan(&bilangan)

	if bilangan < 0 && bilangan%2 == 0 {
		fmt.Printf("genap negatif")
	} else {
		fmt.Printf("bukan")
	}
}
