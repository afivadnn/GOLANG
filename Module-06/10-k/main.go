package main

import "fmt"

func main() {
    var qirat, dinar, dirham, fals, sisa int

    fmt.Scan(&qirat)

    dinar = qirat / 600
    sisa = qirat % 600

    dirham = sisa / 60
    sisa = sisa % 60

    fals = sisa / 6
    sisa = sisa % 6

    fmt.Println(dinar, dirham, fals, sisa)
}
