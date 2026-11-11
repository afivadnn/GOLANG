package main

import "fmt"

func main() {
	var pilihan int

	fmt.Println("1. Mobil")
	fmt.Println("2. Motor")
	fmt.Scan(&pilihan)

	if pilihan == 1 {
		fmt.Println("kamu akan telat")
		fmt.Println("kamu dapat pengurangan nilai")
	} 
	
	if pilihan == 2 {
		fmt.Println("kamu akan tepat waktu")
		fmt.Println("kamu bawa baju ganti")
	}
}

// Program Kendaraan

// Kamus
//     pilihan : integer

// Algoritma
//     output("1. Mobil")
//     output("2. Motor")
//     input(pilihan)

//     if (pilihan = 1) then
//         output("kamu akan telat")
//         output("kamu dapat pengurangan nilai")
//     endif

//     if (pilihan = 2) then
//         output("kamu akan tepat waktu")
//         output("kamu bawa baju ganti")
//     endif
// EndProgram