package main

import "fmt"

func main() {
	var a, b, c, d, e, hasil int

	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)
	fmt.Print("Masukkan nilai c: ")
	fmt.Scan(&c)
	fmt.Print("Masukkan nilai d: ")
	fmt.Scan(&d)
	fmt.Print("Masukkan nilai e: ")
	fmt.Scan(&e)

	hasil = a + b + c + d + e
	fmt.Println("Hasil penjumlahan:", hasil)
}