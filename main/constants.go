package main

import (
	"fmt"
)

const (
	a = iota // 0
	b // 1
	c // 2
)

const (
	a2 = iota // 0
)

const (
	_ = iota + 5 // 0 +5
	// 1 << ( 10 * iota)
	h // 6 = 1+ 5
	k // 7

)

func main(){
	/* maming convention
		PascalCase for exported constants
		camel -> internal
	 */
// type : interoparate with same type
	const MYCONST int = 123
	//const MyConst float64 = math.Abs(5)
	//MyConst = 100 error
//untype : interoparate with similar type
	const (
		PI = 3.14
		HELLOWORLD= "hello world!"
	)
	const x = 12
	var y int16 = 13
	fmt.Println(x+y)
	fmt.Println(PI)
//enumerated
/*
- start at 0
 */
	fmt.Println(a,b,c, a2) // 0 1 2 0

//Enumeration expressions : arithmetic, bitswise bitshifting

/* summary
   - all primitive type : int float string boolen
   	- can shadow
   	- imutable
   	- operation +  %- * /

*/
}