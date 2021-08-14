package main

import "fmt"

/*
summary
- declare : *type , new(type)
- * : get value
- & : get address - hex
- call func by value, reference
*/

func main() {
	var x *int //declare
	var y int
	y = 0
	x = &y

	fmt.Println(x)  // 0x..........
	fmt.Println(&y) // 0x..........
	fmt.Println(*x) // 0
	fmt.Println(y)  // 0

	*x = 1 //get value

	fmt.Println(*x) // 1
	fmt.Println(y)  // 1

	y = 10
	fmt.Println(*x) // 10
	fmt.Println(y)  // 10

	//pass by value, reference
	var a *int
	*a = 10
	fmt.Println(a) // 10
	one(a)
	fmt.Println(a) // 1
	zero(*a)
	fmt.Println(a) //1

	b := 100
	fmt.Println(b) // 100
	zero(b)
	fmt.Println(b) // 100
	one(&b)
	fmt.Println(b) // 1

}

func one(ptr *int) {
	*ptr = 1
}

func zero(num int) {
	num = 0
}
