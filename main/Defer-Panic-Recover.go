package main

import (
	"fmt"
	"log"
)

/*summary
defer
	- defer func
	- func execute after final statement, before return
	- LIFO
	-
// panic
- return error value instead of throw exception
- panic ~ new exception
- defer execute before panic
// recover
- nill = not panicking
- nil ~ none
- ~ throw panic

*/

func main() {
	//defer
	fmt.Println("start")
	fmt.Println("middle")
	fmt.Println("end")
	// start middle end

	fmt.Println("start")
	defer fmt.Println("middle")
	fmt.Println("end")
	// start  end middle

	defer fmt.Println("start")
	defer fmt.Println("middle")
	defer fmt.Println("end")
	//   end middle start

	a := "start"
	defer fmt.Println(a)
	a = "end"
	// start

	//panic
	x, b := 1, 0
	ans := x / b
	fmt.Println(ans)
	// panic : runtione error ....

	m := map[string]int{
		"hanoi": 1,
		"vn":    2,
	}

	_, ok := m["hcm"]
	if !ok {
		panic("key doesnt exist")
	}
	// panic : key doesnt exist

	fmt.Println("start")
	panic("st wrong!")
	fmt.Println("end")
	// start
	// panic : st wrong!

	// recover
	fmt.Println("start")
	defer func() { // anonymous func
		if err := recover(); err != nil { // nill = not panicking
			log.Println("Error :", err)
		}
	}() // throw panic
	panic("st wrong!")
	fmt.Println("end")
	// error : "st wrong!"

}
