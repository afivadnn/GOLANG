package main

import "fmt"

func main(){
	var phi float64 = 3.14
	var r, luas float64
	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scanln(&r)
	luas = phi * r * r
	fmt.Printf("Luas lingkaran adalah: %.1f\n", luas)
}