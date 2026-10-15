package main

import (
    "fmt"
    "unicode"
)

func main() {
    var c1, c2 rune
	var b1, b2 bool
    fmt.Scanf("%c%c", &c1, &c2)

    b1 = !(unicode.IsLetter(c1) || unicode.IsDigit(c1))
    b2 = !(unicode.IsLetter(c2) || unicode.IsDigit(c2))

    fmt.Printf("%t %t\n", b1, b2)
}
