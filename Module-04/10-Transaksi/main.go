package main
import "fmt"
type DataTransaksi struct{
	nama_barang string
	jumlah int
	harga_barang float64
	total_harga_barang float64
}
func main(){
	var d DataTransaksi
	fmt.Println("Inputkan Data : ")
	fmt.Print("Nama Barang : ")
	fmt.Scanln(&d.nama_barang)
	fmt.Print("Jumlah Barang : ")
	fmt.Scanln(&d.jumlah)
	fmt.Print("Harga Barang : ")
	fmt.Scanln(&d.harga_barang)
	d.total_harga_barang = float64(d.jumlah) * d.harga_barang
	fmt.Printf("Total Harga Barang : Rp. %.2f", d.total_harga_barang)
}