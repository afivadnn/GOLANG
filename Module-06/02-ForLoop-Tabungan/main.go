package main

import "fmt"

func main() {
    var minggu, saldo, tambahan int

    fmt.Scan(&minggu, &saldo, &tambahan)
	for i := 1; i <= minggu; i++ {
        saldo += i * tambahan
    }
    fmt.Println(saldo)
}