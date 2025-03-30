package main

import (
	"errors"
	"fmt"
)

func main() {
	var printValue string = "hello"
	printMe(printValue)

	var num int = 11
	var deno int = 1
	var result, remain, err = intDivision(num, deno)
	switch {
	case err != nil:
		fmt.Printf(err.Error())
		break
	case remain == 0:
		fmt.Print("result:%v", result)
		break
	default:
		fmt.Print("result:%v, with remainder:%v", result, remain)
		break
	}
	fmt.Printf("result:%v, with remainder:%v", result, remain)
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(num int, deno int) (int, int, error) {
	var err error
	if deno == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	var result int = num / deno
	var remain int = num % deno
	return result, remain, nil
}
