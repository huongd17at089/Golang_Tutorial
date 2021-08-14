package main

import "fmt"

/* summary
if
- if-else if-if
- initalizer
- > < ==
- || &&
switch
- match tag
- initalizer
- case multipe test
- no tag ~ if-else if-else
- break >< fallthrough keyword
*/

func main() {
	// if
	if true {
		fmt.Println("true")
	}

	m := map[string]int{
		"hanoi": 1,
		"vn":    2,
	}

	if pop, ok := m["vn"]; ok {
		fmt.Println(pop)
	}

	//fmt.Println(pop) // undefine error
	// number guessing game

	number := 10
	guess := 25

	if guess < 1 || guess > 100 {
		fmt.Println("1<guess<100")
	}

	// nested
	if guess >= 1 && guess <= 100 {
		if guess < number {
			fmt.Println("too low")
		}
		if guess > number {
			fmt.Println("too high")
		}
		if guess == number {
			fmt.Println("got it")
		}
	}

	fmt.Println(!true) // false
	// if-else if-else
	ok := true
	if ok {
		fmt.Println("ok")
	} else {
		fmt.Println("not oke")
	}
	// switch
	switch 2 {
	case 1, 3, 5:
		fmt.Println(1)
	//case 1,3,5: ..overlapp
	//	fmt.Println(1)
	case 2, 4, 6:
		fmt.Println(2)
	default:
		fmt.Println(3)
	}

	switch i := 1 + 1; i {
	case 1, 3, 5:
		fmt.Println(1)
	//case 1,3,5: ..overlapp
	//	fmt.Println(1)
	case 2, 4, 6:
		fmt.Println(2)
	default:
		fmt.Println(3)
	}

	x := 10
	switch { // if-else if-else
	case x < 10:
		fmt.Println(1)
		fmt.Println(10) //
		fallthrough     // >< break
	//case 1,3,5: ..overlapp
	//	fmt.Println(1)
	case x > 10:
		fmt.Println(2)
		break
		fmt.Println(20) // not execute
	default:
		fmt.Println(3)
	}

	//var i interface{} = 1
	//switch i.(type) {
	//case int:
	//	fmt.Println("int")
	//	//.....
	//case [2]int:
	//	fmt.Println("ok")
	//default:
	//	fmt.Println("string")
	//}

}
