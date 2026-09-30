package main

import "fmt"

func main() {
	var rupiah int
	fmt.Print("Masukkan jumlah uang dalam rupiah: ")
	fmt.Scan(&rupiah)
	dolar := float64(rupiah) / 15000
	fmt.Printf("Jumlah uang dalam dolar adalah: %.2f\n", dolar)
}