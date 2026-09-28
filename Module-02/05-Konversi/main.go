package main

import"fmt"

func main(){
	var fahrenheit, celcius int

	fmt.Print("Masukkan suhu dalam Fahrenheit: ")
	fmt.Scanln(&fahrenheit)
	celcius = (fahrenheit - 32) * 5 / 9
	fmt.Print("Suhu dalam Celcius: ", celcius)
}