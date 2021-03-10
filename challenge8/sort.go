package challenge8

import "fmt"

//BubbleSort is func
func BubbleSort(list []int) []int {

	fmt.Println("input: ", list)

	unsorted := true
	for unsorted {
		unsorted = false
		for i := 1; i < len(list); i++ {

			if list[i-1] > list[i] {
				list[i], list[i-1] = list[i-1], list[i]
				unsorted = true
			}

		}
	}

	fmt.Println("output", list)
	return list
}
