package main

import (
	"Project/challenge1"
	"Project/challenge10"
	"Project/challenge2"
	"Project/challenge3"
	"Project/challenge4"
	"Project/challenge5"
	"Project/challenge6"
	"Project/challenge7"
	"Project/challenge8"
	"Project/challenge9"
)

func main() {
	challenge1.Grader(10)
	challenge2.OddOrEven(10)
	challenge3.Stat([]int{10, 2, 30})
	challenge4.Leap()
	challenge5.Palindrom("LMAO")
	challenge6.Age()
	challenge7.ReverseSentence()
	challenge8.BubbleSort([]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0})
	challenge9.GroupArray([]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0})
	challenge10.TimeConverter()
}
