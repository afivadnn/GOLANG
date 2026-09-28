package main

import "fmt"

func main() {

	var nama, nim, kelas string
	fmt.Print("Masukkan nama: ")
	fmt.Scanln(&nama)
	fmt.Print("Masukkan NIM: ")
	fmt.Scanln(&nim)
	fmt.Print("Masukkan kelas: ")
	fmt.Scanln(&kelas)
	fmt.Printf("Perkenalkan saya adalah %s, salah satu mahasiswa Prodi S1-IF dari kelas %s, dengan NIM %s\n", nama, kelas, nim)
}