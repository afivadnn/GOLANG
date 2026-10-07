package main
import "fmt"

func main() {
	var n, d1,d2,d3 int
	fmt.Scan(&n)
	d1 = n / 100
	d2 = (n % 100) / 10
	d3 = n % 10
	fmt.Println(d1 <= d2 && d2 <= d3)
}