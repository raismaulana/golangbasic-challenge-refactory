package challenge10

import (
	"fmt"
)

//TimeConverter is func
func TimeConverter() {
	var output string
	var converter int
	var input float64
	fmt.Print("Pilih jenis konversi:\n1 : Konversi jam ke menit\n2 : Konversi menit ke jam\n3 : konversi menit ke tahun\nPilihan:")
	fmt.Scan(&converter)
	switch converter {
	case 1:
		fmt.Print("Konversi jam ke menit\nMasukkan angka jam:")
		fmt.Scan(&input)
		a := input * 60
		output = fmt.Sprint(float64(a), " Menit")
	case 2:
		fmt.Print("Konversi menit ke jam\nMasukkan angka menit:")
		fmt.Scan(&input)
		a := input / 60
		output = fmt.Sprint(float64(a), " jam")
	case 3:
		fmt.Print("Konversi menit ke tahun\nMasukkan angka menit:")
		fmt.Scan(&input)
		a := input / 525600
		output = fmt.Sprint(float64(a), " tahun")
	}
	fmt.Println(output)
}
