package main

import (
	"fmt"
)

func main() {
	var n, alas, tinggi int
	fmt.Println("Masukan n:")
	fmt.Scanln(&n)

	for i := n; i <= n; i++ {
		fmt.Println("Masukan alas:")
		fmt.Scanln(&alas)
		fmt.Println("Masukan tinggi:")
		fmt.Scanln(&tinggi)
		luas := alas * tinggi / 2
		fmt.Println(luas)
	}
}
