package main

import "fmt"

func main() {
	var num int

	fmt.Print("masukan bilangan")
	fmt.Scan(&num)

	if num < 0 && num%2 == 0 {
		fmt.Println(true)
		return
	}
	fmt.Println(false)
}
