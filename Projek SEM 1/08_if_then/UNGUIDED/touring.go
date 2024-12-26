package main

import "fmt"

func main() {
	var peserta int

	fmt.Print("masukan jumlah peserta: ")
	fmt.Scan(&peserta)

	motor := peserta / 2
	if peserta%2 != 0 {
		motor++
	}
	fmt.Println("jumlah motor yang diberikan: ", motor)
}
