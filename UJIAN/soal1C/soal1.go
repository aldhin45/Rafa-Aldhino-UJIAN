package main

import "fmt"

func main() { //2311102023
	var angka int
	total := 0
	fmt.Println("masukan bilangan bulat positif lebih dari 100")

	for {
		fmt.Print("masukan angka :")
		fmt.Scanln(&angka)
		if angka < 0 {
			break
		}
		total += angka
	}
	fmt.Printf("total angka yang dimasukan adalah : %d\n", total)
}
