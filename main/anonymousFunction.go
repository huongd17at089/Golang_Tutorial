package main

import "fmt"

/*
summary
- func as type : assign, pass , return
*/

func main() {
	// anonymous func
	//func (){
	//	println("hello")
	//}
	hi := "hi" // main scope
	func() {
		msg := "hi" // func scope
		println("hello", msg)
	}()

	func(msg string) {
		println("hello", msg)
	}(hi)

	f := func() {
		msg := "hi" // func scope
		println("hello", msg)
	}
	f()

	var div func(float64, float64) (float64, error)
	div = func(a, b float64) (float64, error) {
		if b == 0.0 {
			return 0.0, fmt.Errorf("st wrong")
		}
		return a / b, nil

	}

	d, err := div(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(d)
}
