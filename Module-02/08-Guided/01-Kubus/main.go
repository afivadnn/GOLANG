package main

import"fmt"

func main(){
	var nilai int
	fmt.Print("Masukkan panjang sisi kubus: ")
	fmt.Scan(&nilai)
	var sisi float64 = float64(nilai)
	volume := sisi * sisi * sisi
	fmt.Printf("Volume kubus adalah: %2f\n", volume)
}