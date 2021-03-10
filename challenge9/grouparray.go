package challenge9

import (
	"Project/challenge8"
	"fmt"
)

//GroupArray is func
func GroupArray(list []int) {
	list = challenge8.BubbleSort(list)
	var newList [][]int
	for i := 0; i < len(list); i += 2 {
		newList = append(newList, list[i:i+2])
	}
	fmt.Println(newList)
}
