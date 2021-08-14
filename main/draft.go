package main

import (
	"fmt"
	"strconv"
	"strings"
)

func countZeroDigit(a int, b int) int {
	count := 0
	for i := a; i <= b; i++ {
		temp := strconv.Itoa(i)
		if strings.Contains(temp, "0") {
			count = count + 1
		}
	}

	return count
}

func main() {
	//a := map[string]int{
	//	"hanoi" : 1,
	//	"vn" : 2,
	//}
	//b := a
	//makemap := make(map[string]int)
	//makemap = a
	//delete(b, "vn")
	//fmt.Println(b) // hanoi :1
	//fmt.Println(a) // hanoi : 1
	//fmt.Println(makemap)

	//a2 := [...]int{1, 2, 3, 4, 5, 6, 7}
	//suba2 := a2[1:4]
	//a2[2] = 42
	//fmt.Println(suba2)
	//suba2[0] = 100
	//fmt.Println(a2)
	//fmt.Println(suba2)

	//var x *int
	//var y int
	//y = 0
	//x = &y
	//
	//fmt.Println(x)
	//fmt.Println(&y)
	//fmt.Println(*x)
	//fmt.Println(y)
	//fmt.Println("___________________________")
	//
	//*x = 1
	//
	//fmt.Println(*x)
	//fmt.Println(y)
	//fmt.Println("___________________________")
	//
	//y = 10
	//fmt.Println(*x)
	//fmt.Println(y)
	fmt.Print(countZeroDigit(100, 110))
}
