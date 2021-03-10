package challenge4

import "fmt"

// Leap is function to search leap year
func Leap() {
	var fYear, lYear int
	fmt.Println("Masukkan tahun awal: ")
	fmt.Scan(&fYear)
	fmt.Println("Masukkan tahun akhir: ")
	fmt.Scan(&lYear)
	fmt.Println("Output: ")

	if fYear <= lYear {
		for i := fYear; i <= lYear; i++ {
			if i%4 == 0 {
				fmt.Println(i)
			}
		}
	} else {
		fmt.Println("Masukan salah")
	}
}
