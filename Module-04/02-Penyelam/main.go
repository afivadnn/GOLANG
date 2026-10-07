package main

import "fmt"

func main() {
	var kedalaman, batas, tekanan  float64
	var hasil bool

	fmt.Scan(&kedalaman, &batas)
	tekanan = 1025 * 9.8 * kedalaman
	hasil = tekanan <= batas
	fmt.Printf("%.1f\n%t", tekanan, hasil)
}