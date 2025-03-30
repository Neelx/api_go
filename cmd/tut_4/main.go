package main //arrays, slices, and maps

import "fmt"

func main() {
	intArr := [...]int32{1, 2, 3}
	fmt.Println(intArr)

	var intSlice []int32 = []int32{4, 5, 7}
	fmt.Println(intSlice)
	intSlice = append(intSlice, 7)
	fmt.Printf("len:%v, cap:%v\n", len(intSlice), cap(intSlice))

	var intsLice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intsLice2...)
	fmt.Printf("len:%v, cap:%v\n", len(intSlice), cap(intSlice))

	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Println(intSlice3)

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{
		"Adam": 21,
		"Eve":  43,
	}
	fmt.Println(myMap2["Adam"])
	var age, ok = myMap2["Jason"]
	if ok {
		fmt.Println("age:%v", age)
	} else {
		fmt.Println("not found")
	}

	for name := range myMap2 {
		fmt.Printf("name:%v\n, age:%v\n", name, age)
	}

	for i, v := range intArr {
		fmt.Printf("Index:%v, Value:%v\n", i, v)
	}
	var i int = 0
	for {
		if i >= 10 {
			break
		}
		fmt.Println(i)
		i = i + 1
	}
}
