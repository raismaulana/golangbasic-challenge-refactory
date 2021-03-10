package challenge2

import "fmt"

//OddOrEven is function
func OddOrEven(val int) {
	mod := val % 2
	if mod == 0 {
		fmt.Println("Genap")
	} else {
		fmt.Println("Ganjil")
	}
}
