package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum uint = 32767
	intNum = intNum + 1
	fmt.Println(intNum)

	var floatNum float64 = 12121.22222
	floatNum = floatNum + 1
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.2
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 int = 4
	fmt.Println(intNum1 / intNum2)
	fmt.Println(intNum1 % intNum2)

	var myString string = "Hello" + " " + "world!"
	fmt.Println(myString)

	fmt.Println(utf8.RuneCountInString(myString))

	var myRune rune = 'a'
	fmt.Println(myRune)

	var myBoolen bool = false
	fmt.Println(myBoolen)

	var intNum3 int
	fmt.Println(intNum3)

	var myVar string = "text"
	fmt.Println(myVar)

	var var1, var2, var3 = 1, 2, 3
	fmt.Println(var1, var2, var3)

	const myConst string = "const calee"
	fmt.Println(myConst)

	const pi float32 = 3.1453
}
