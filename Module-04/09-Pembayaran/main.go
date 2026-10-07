package main
import "fmt"

func main(){
	var pesanan, total int
	const harga = 20000 
	const pajak = 0.11
	fmt.Scan(&pesanan)
	total = pesanan * harga + int(float64(pesanan * harga) * pajak)
	fmt.Println(total)

}