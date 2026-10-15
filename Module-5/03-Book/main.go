package main
import "fmt"
type DataBuku struct {
		judul string
		genre string
		tahun int
		rating float64
	}
func main() {
	var buku1, buku2 DataBuku
	fmt.Scan(&buku1.judul, &buku1.genre, &buku1.tahun, &buku1.rating)
	fmt.Scan(&buku2.judul, &buku2.genre, &buku2.tahun, &buku2.rating)
	fmt.Println("--- Data Buku 1 ---")
	fmt.Println(buku1.judul)
	fmt.Println(buku1.genre)
	fmt.Println(buku1.tahun)
	fmt.Println( buku1.rating)
	fmt.Println("--- Data Buku 2 ---")
	fmt.Println( buku2.judul)
	fmt.Println( buku2.genre)
	fmt.Println( buku2.tahun)
	fmt.Println( buku2.rating)
}

// Pseudocode
// program Buku
// kamus 
// 	type DataBuku <
// 	judul : string
// 	genre : string	
// 	tahun : integer
// 	rating : real
// >
// buku1, buku2 : DataBuku
// algoritma
// 	input (buku1.judul, buku1.genre, buku1.tahun, buku1.rating)
// 	input (buku2.judul, buku2.genre, buku2.tahun, buku2.rating)
// 	output (buku1.judul, buku1.genre, buku1.tahun, buku1.rating)
// 	output (buku2.judul, buku2.genre, buku2.tahun, buku2.rating)
// endprogram
