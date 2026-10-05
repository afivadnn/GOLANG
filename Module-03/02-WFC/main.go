package main
import "fmt"

func main() {
    var n, hari int
    _, err := fmt.Scan(&n)
    if err != nil || n < 1 || n > 1000 {
        return 
    }

    hari = ((4 + n - 1) % 7)+1
    fmt.Printf("%d", hari)
}
