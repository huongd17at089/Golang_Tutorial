package main

import "fmt"

/*
summary
func <name> (<parameter>) (<return value>){
}
- pass value( int, float, struct, array)  , pass pointer (map, pointer, slice)
- variadic parameters : ~ pass slice, last parameter
- multi return value
- stack memory , heap memory ???
*/

func main() {
	saySt("hello")
	name := "haha"
	greeting := "hello"
	saySt4(&greeting, &name)
	fmt.Println(name)
	// hello haha
	// hi
	//hi
	sum("hihi", 2, 3, 4, 5)
	s := sum2(1, 2, 3, 4)
	fmt.Println(s)

	s2 := sum3(1, 2, 3) //pointer
	fmt.Println(*s2)

	d, err := divide2(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)

}

// parameter
func saySt(msg string) {
	println(msg)
}

func saySt2(msg string, name string) {
	println(msg + name)
}

func saySt3(msg, name string) {
	println(msg + name)
}

func saySt4(msg, name *string) {
	println(*msg + *name)
	*name = "hi"
	fmt.Println(*name)
}

func sum(msg string, values ...int) {
	println(values)
	res := 0
	for _, v := range values {
		res += v
	}
	fmt.Println(msg, res)
}

// return func
func sum2(values ...int) int {
	println(values)
	res := 0
	for _, v := range values {
		res += v
	}
	return res
}

func sum3(values ...int) *int {
	println(values)
	res := 0
	for _, v := range values {
		res += v
	}
	return &res
}

func sum4(values ...int) (res *int) {
	println(values)
	*res = 0
	for _, v := range values {
		*res += v
	}
	return
}

func divide(a, b float64) (d float64) {
	if b == 0 {
		panic("st wrong!") // break execution
	}
	d = a / b
	return
}

func divide2(a, b float64) (float64, error) {
	if b == 0 {
		return 0.0, fmt.Errorf("st wrong!")
	}
	return a / b, nil
}
