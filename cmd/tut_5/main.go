package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = "résumé"
	var indexed = myString[0]
	fmt.Printf("%v, %T", indexed, indexed)
	for i, v := range myString {
		fmt.Println(i, v)
	}
	fmt.Println("the len of string is: ", len(myString))

	var myRune = 'a'
	fmt.Println("\nmyRune =", myRune)

	var strSlice = []string{"s", "u", "b", "s"}
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var catStr = strBuilder.String()
	fmt.Printf("\n%v", catStr)
}
