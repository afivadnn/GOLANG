package main

import"fmt"

func main(){
	var jumlahKoin, emas, perak, tembaga int
	fmt.Scan(&jumlahKoin)
	emas = jumlahKoin / 9
	perak = (jumlahKoin % 9) / 3
	tembaga = jumlahKoin % 3
	fmt.Printf("%d %d %d",emas, perak, tembaga)

}