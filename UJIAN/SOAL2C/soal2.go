package main

import "fmt"

func main() { //2311102023
	var n int
	fmt.Scan(&n)

	utama, sembako, konsol := 0, 0, 0

	for i := 0; i < n; i++ {
		var tiket int
		fmt.Scan(&tiket)

		if sama && digitAwal%2 == 0 {
			utama++
		} else if ganjil {
			sembako++
		} else {
			konsol++
		}
	}

	fmt.Println("Jumlah Hadiah:")
	fmt.Println("Utama:", utama)
	fmt.Println("Sembako:", sembako)
	fmt.Println("Konsol:", konsol)
}
