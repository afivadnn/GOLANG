package main

import "fmt"

func main() {
	var alas, tinggi int
	fmt.Print("Masukkan panjang alas segitiga: ")
	fmt.Scan(&alas)
	fmt.Print("Masukkan tinggi segitiga: ")
	fmt.Scan(&tinggi)
	luas := 0.5 * float64(alas) * float64(tinggi)
	fmt.Printf("Luas segitiga adalah: %2f\n", luas)
}