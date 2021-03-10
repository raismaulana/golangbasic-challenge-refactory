package challenge3

import (
	"fmt"
	"sort"
)

//Stat is function
func Stat(arr []int) {
	fmt.Printf("Max:\t %d \nMin:\t %d \nMedian:\t %d \nAverage: %f\n", max(arr), min(arr), median(arr), average(arr))
}

func average(arr []int) float64 {
	var sum int
	for _, v := range arr {
		sum = sum + v
	}
	return float64(sum / len(arr))
}

func max(arr []int) int {
	sort.Ints(arr)
	return arr[len(arr)-1]
}

func min(arr []int) int {
	sort.Ints(arr)
	return arr[1]
}

func median(arr []int) int {
	sort.Ints(arr)
	half := len(arr) / 2
	if len(arr)%2 == 0 {
		return arr[half] + arr[half+1]/2
	}
	return arr[half+1]
}
