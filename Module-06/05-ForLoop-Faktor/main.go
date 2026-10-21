package main

import (
	"fmt"
)

func main() {
    var a, i, hasil int
    fmt.Scan(&a)
	hasil = 1
    for i = 1; i <= a; i++ {
        hasil *= i
    }
	fmt.Print(hasil)
}
