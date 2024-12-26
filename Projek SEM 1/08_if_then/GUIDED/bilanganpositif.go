package main

import "fmt"

func main() {
	var bilangan int

	fmt.Print("Masukan bilangan: ")
	fmt.Scan(&bilangan)

	if bilangan > 1 {
		fmt.Println("positif")
	} else {
		fmt.Println("negatif")
	}
}
