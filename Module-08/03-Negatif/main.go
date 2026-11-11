package main
import "fmt"

func main(){
	var bil int
	var teks string
	fmt.Scan(&bil)
	teks = "Bilangan negatif"
	if bil > 0 {
		bil = -bil
		teks = "Bilangan positif"
	} 
	fmt.Println(teks)
}