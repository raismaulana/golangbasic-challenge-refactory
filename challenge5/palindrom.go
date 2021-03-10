package challenge5

import (
	"fmt"
	"strings"
)

//Palindrom is function
func Palindrom(word string) {
	upTheWord := strings.ToUpper(word)
	var backword string
	for _, v := range upTheWord {
		backword = string(v) + backword
	}

	if upTheWord == backword {
		fmt.Printf("%s adalah palindrom\n", word)
	} else {
		fmt.Printf("%s adalah bukan palindrom\n", word)
	}
}
