package main

import "fmt"

/*
summary
- declare : *type , new(type)
- * : get value
- & : get address - hex
- call func by value, reference
- slice, map contain internal pointers
*/

func main() {

	aa := 100
	bb := 200
	fmt.Println(aa, bb) // 100 200
	aa = 300
	fmt.Println(aa, bb) // 300 200

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

	var ms *myStruct
	fmt.Println(ms) // nil
	ms = new(myStruct)
	fmt.Println(ms) //{0}
	ms.foo = 100
	fmt.Println(ms.foo) // 100

}

func one(ptr *int) {
	*ptr = 1
}

func zero(num int) {
	num = 0
}

type myStruct struct {
	foo int
}
