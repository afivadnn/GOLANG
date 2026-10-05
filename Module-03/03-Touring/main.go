package main

import"fmt"

func main(){
	var kecepatan, sisaJarak, jam, menit int
	const totalJarak int = 330
	fmt.Printf("Masukkan kecepatan (km/jam): ")
	fmt.Scan(&kecepatan)
	jam = totalJarak / kecepatan
	sisaJarak = totalJarak % kecepatan
	menit = (sisaJarak * 60) / kecepatan
	fmt.Printf("%d km/jam = %d jam, %d menit\n", kecepatan, jam, menit)
}