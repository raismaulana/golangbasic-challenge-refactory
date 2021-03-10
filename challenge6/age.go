package challenge6

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

//Age is func
func Age() {
	var bday string
	var days, years int

	fmt.Print("Tanggal lahir : ")
	fmt.Scan(&bday)
	day, month, year := splitBday(bday)

	a := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	b := time.Now()

	leapDays := calcLeap(year, b.Year())

	for a.Before(b) {
		// fmt.Println(a, days)
		a = a.Add(time.Hour * 24)
		days++
	}
	// fmt.Println(a, days)

	years = days / 364
	days = (days % 364) - leapDays
	fmt.Println("Anda berumur : ", years, "tahun, ", days, "hari")

}

func splitBday(bday string) (day, month, year int) {
	q := strings.Split(bday, "-")
	day, _ = strconv.Atoi(q[0])
	month, _ = strconv.Atoi(q[1])
	year, _ = strconv.Atoi(q[2])
	return
}

func calcLeap(fYear, lYear int) (leap int) {
	for i := fYear; i <= lYear; i++ {
		if i%4 == 0 {
			leap++
		}
	}
	return
}
