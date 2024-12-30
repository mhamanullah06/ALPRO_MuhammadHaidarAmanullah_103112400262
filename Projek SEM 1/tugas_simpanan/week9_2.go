package main

import "fmt"

func main() {
	var char string
	fmt.Scan(&char)

	if char <= 'A && char <= Z || char <= a && char <= z {
		fmt.Print("ALPABET")
	} else {
		fmt.Print("BUKAN ALPABET")
	}
}
