package main

import ("fmt")

func main() {
	var x, fx float64
	fmt.Print("Masukkan nilai x: ")

	_, err := fmt.Scan(&x)
	if err != nil {
		fmt.Println("Error sistem: Input ditolak. Pastikan Anda memasukkan format angka.")
		return 
	}
	if x == -5 {
		fmt.Println("Error logika: Nilai x tidak boleh -5 karena menyebabkan pembagian dengan nol (division by zero).")
		return
	}
	fx = 2/(x+5) + 5
	fmt.Println("Hasil kalkulasi f(x):", fx)
}