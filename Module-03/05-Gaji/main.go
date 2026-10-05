package main

import"fmt"

func main(){
	var gajiPokok, jamLembur int
	fmt.Scan(&gajiPokok, &jamLembur)
	bonus := jamLembur * 45000
	potongan := (gajiPokok * 55) / 1000
	gajiBersih := gajiPokok + bonus - potongan
	fmt.Printf("%d\n", gajiBersih)

}
