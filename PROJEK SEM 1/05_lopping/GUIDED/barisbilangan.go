package main

import "fmt"

func main() {
	var n, m int
	fmt.Println("mMsukan bilangan pertama")
	fmt.Scanln(&n)
	fmt.Println("Masukan bilangan kedua")
	fmt.Scanln(&m)
	for i := n; i <= m; i++ {
		fmt.Println(i)
	}
}
