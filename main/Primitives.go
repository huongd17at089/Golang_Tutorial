package main

import "fmt"

func main(){
	/*
	*/

	// boolean
	var bu bool = true
	fmt.Println(bu)
	 n:= 1==1
	 m := 1 == 2
	 fmt.Println(n, m) // true false
	 var temp bool
	 fmt.Println(temp) // 0 -> false
	 // int int8 int16 int32 int64
	 // uint8 unit16 uint32
	 a := 10 // int64
	 b := 3 // int64
	 fmt.Println(a + b) // - * / %
	 var c int8 = 11
	 //d := a+c // error
	 fmt.Println(a + int(c))
	// bit operator & | ^ &^ << >>
	// float + - * /
	x := 3.14
	x = 3.14e3
	x = 3.1E3
	fmt.Println(x)
	// complex + - * /
	//var comp complex64 = 1 + 2i
	var comp complex64 = complex(1,2)
	fmt.Println(comp)
	fmt.Println(real(comp), imag(comp)) // type float
	// string concat(+)
	s := "string"
	fmt.Println(s[0]) // type int8
	//s[2] = "u" // error
	bi := []byte(s) // utf8
	fmt.Println(bi) // collection of byte
	char := 'a' // rune utf32
	// var char rune
	fmt.Println(char) // type int32

	/* summary
	// cant convert bool to int bit 0 is false
	// string immutable
	// string concat with +
	// string can be convert to bytes[]
	*/
}
