package main
import "fmt"

func main(){
	const (
		diskon1 = 0.10
		diskon2 = 0.20
	)
	var harga_awal, harga_diskon1, harga_akhir float64
	fmt.Scan(&harga_awal)
	harga_diskon1 = harga_awal * (1 - diskon1)
	harga_akhir = harga_diskon1 * (1 - diskon2)
	fmt.Printf("%.2f", harga_akhir)
}