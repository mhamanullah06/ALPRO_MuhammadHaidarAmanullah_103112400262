package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Print("Masukkan satu karakter: ")
	fmt.Scanln(&input)

	if len(input) != 1 {
		fmt.Println("Harap masukkan hanya satu karakter.")
		return
	}
	char := input[0]
	var jenis string

	if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') {
		if char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U' ||
			char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
			jenis = "vokal"
		} else {
			jenis = "konsonan"
		}
	} else {
		jenis = "bukan huruf"
	}

	fmt.Printf("%s: %s\n", input, jenis)
}
