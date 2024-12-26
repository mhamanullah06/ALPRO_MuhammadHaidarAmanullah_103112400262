package main

import "fmt"

func main() {
	var umur int
	fmt.Print("Masukan Umur: ")
	fmt.Scan(&umur)

	var kk bool
	fmt.Print("Apakah memiliki Kartu Keluarga: ")
	fmt.Scan(&kk)

	if umur >= 17 && kk {
		fmt.Println("anda bisa memuat ktp.")
	} else {
		fmt.Println("belum bisa membuat ktp.")
	}
}
