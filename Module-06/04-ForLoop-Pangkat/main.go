package main
import "fmt"

func main(){
	var a,b,i, hasil int
	fmt.Scan(&a,&b)
	hasil = 1
	for i = 1 ;i <= b; i++{ //a b, b pangkat i harus ke b kali 
		hasil *= a //hasil akan sell dikalikan dng a sebanyak b kali
	}
	fmt.Print(hasil)
}

