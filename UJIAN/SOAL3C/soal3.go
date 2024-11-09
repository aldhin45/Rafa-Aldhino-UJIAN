package main

import "fmt"

func main() { //2311102023
	var n int
	fmt.Println("masukan bilangan penyebut n :")
	fmt.Scan(&n)
	fmt.Println(sum(n))

}

func sum(n int) int {
	var i, sum int
	fmt.Println("masukan bilangan penyebut m :")
	fmt.Scan(&i)
	for n >= i {
		sum += i
		i++
	}
	return sum
}
