package main
import"fmt"

func main(){
	var tinggi, bmi, berat float64
	fmt.Scan(&bmi, &tinggi)
	berat = bmi * (tinggi * tinggi)
	fmt.Printf("%.f", berat) 
}
