package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e gasEngine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("can make it")
	} else {
		fmt.Println("cannot make it")
	}
}

func main() {
	var myEngine gasEngine = gasEngine{25, 15}
	canMakeIt(myEngine, 100)
}
