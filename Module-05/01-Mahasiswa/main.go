package main
import "fmt"

type Mahasiswa struct {
	NIM string
	nama string
	nilai_tugas float64
	nilai_kuis float64
	nilai_praktikum float64
	nilai_uts float64
	nilai_uas float64
	nilai_akhir float64

}

func main() {
	const(
		kuis float64 = 0.15
		tugas float64 = 0.15
		praktikum float64 = 0.2
		uts float64 = 0.25
		uas float64 = 0.25
	)

	var m Mahasiswa
	fmt.Scan(&m.nama, &m.NIM, &m.nilai_tugas, &m.nilai_kuis, &m.nilai_praktikum, &m.nilai_uts, &m.nilai_uas)
	m.nilai_akhir = (m.nilai_tugas * tugas) + (m.nilai_kuis * kuis) + (m.nilai_praktikum * praktikum) + (m.nilai_uts * uts) + (m.nilai_uas * uas)
	fmt.Println("--- Data Mahasiswa ---")
	fmt.Println("Nim", m.NIM)
	fmt.Println("Nama", m.nama)
	fmt.Println("Tugas", m.nilai_tugas)
	fmt.Println("Kuis", m.nilai_kuis)
	fmt.Println("Praktikum", m.nilai_praktikum)
	fmt.Println("UTS", m.nilai_uts)
	fmt.Println("UAS", m.nilai_uas)
	fmt.Println("Nilai Akhir", m.nilai_akhir)
    
}