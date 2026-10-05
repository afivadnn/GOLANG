package main

import"fmt"

func main(){
	var n, jam, menit, detik int
	fmt.Scan(&n)
	jam = n / 3600
	menit = (n % 3600) / 60
	detik = n % 60
	fmt.Printf("%d %d %d", jam, menit, detik)
}