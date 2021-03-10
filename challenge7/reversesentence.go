package challenge7

import (
	"fmt"
	"strings"
)

//ReverseSentence is func
func ReverseSentence() {
	var newSentence string
	sentence := "HEllo World iM RobOt"
	words := strings.Fields(sentence)
	fmt.Println(sentence, words)
	for _, v := range words {
		newSentence = v + " " + newSentence
	}

	fmt.Println(newSentence)
}
