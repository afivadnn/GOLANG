package main
import "fmt"

func main(){
	var belanja, diskon, total int
	fmt.Scan(&belanja, &diskon)
	total = belanja - (belanja * diskon / 100)
	fmt.Println(total)
}
