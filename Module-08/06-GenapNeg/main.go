package main
import "fmt"

func main(){
	var bil int	
	var teks string
	fmt.Scan(&bil)
	teks = "Bukan"
	if bil % 2 == 0 && bil < 0 {
		teks = "Genap negatif"
	}
	fmt.Println(teks)
}