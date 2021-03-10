package challenge1

import (
	"fmt"
)

//Grader is an function for grader a value (int)
func Grader(value int) {
	if value >= 90 {
		fmt.Println("A")
	} else if value >= 80 && value <= 89 {
		fmt.Println("B")
	} else if value >= 70 && value <= 79 {
		fmt.Println("C")
	} else if value >= 60 && value <= 69 {
		fmt.Println("D")
	} else {
		fmt.Println("E")
	}
}
