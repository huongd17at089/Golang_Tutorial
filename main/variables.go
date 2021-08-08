package main

import (
	"fmt"
	"strconv"
)

/*summary
- 3 way declare variable
- cant redeclare but can shadow
- all var must be used
- 3 level scope
- camelCase : internal
- PascalCase : export
*/

//package lavel
var g int = 123

var (
	name string = "name"
	age  int    = 10
)

var shadow_var string = "1st"

func main() {
	fmt.Println(123)
	//declaration :

	//function level
	var a int
	a = 1
	var b int = 2
	c := 123
	fmt.Println(a, b, c)
	// variable type
	var x float64 = 10
	y := 10
	fmt.Println("%v %T", x, x) // float 64
	fmt.Println("%v %T", y, y) // int
	z := 10.
	fmt.Println("%v %T", z, z) // float64
	// var t int = "ok" ->  string type error
	/*redeclaration and shadowing
	- cant declare the variable twice in the same scope
	- declare and not used -> compile error
	*/
	//shadowing
	fmt.Println(shadow_var) // 1st
	shadow_var := "2nd"
	fmt.Println(shadow_var) // 2nd
	/*visibility and naming convention
	3 level
	- package -> lowercase first letter
	- globally -> uppercase first letter
	- block scope

	*/
	// type conversions
	j := float32(123)
	fmt.Println(j)
	// convert num to string
	var num int = 123
	var s string = string(123)
	fmt.Println(num, s)
	s = strconv.Itoa(num)

}
