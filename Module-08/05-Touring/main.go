package main
import "fmt"

func main() {
	var org , mtr int
	fmt.Scan(&org)
	mtr = org / 2
	if org % 2 != 0 {
		mtr++
	}
	fmt.Println(mtr)
}
