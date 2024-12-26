package main

import "fmt"

func main() {
	var bilangan int

	fmt.Print("Masukan bilangan: ")
	fmt.Scan(&bilangan)

	if bilangan < 0 {
		bilangan = -bilangan
	}
	fmt.Print(bilangan)
}
