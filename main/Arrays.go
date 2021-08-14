package main

import "fmt"

/* sumary
Array
- item same type
- fixed size
- 3  declare style
- 0-base index
- copy
- len
Slice
- slice existing array  or slice
- dynamic size
- make()
- len, cap
- apply
*/

func main() {
	//array fixed size
	//create
	arr := [3]int{1, 2, 3} //size = 3
	//arr1 := [...] int{1,2,3}
	fmt.Println(arr) //[1,2,3]
	var students [3]string
	students[0] = "lisa"
	fmt.Println(students) // [lisa  ] // zero string
	//array of array
	var matrix [3][3]int = [3][3]int{[3]int{1, 2, 3}, [3]int{1, 2, 3}, [3]int{1, 2, 3}}
	fmt.Println(matrix)
	//build in
	fmt.Println(len(students)) //3
	// work with
	a := [...]int{1, 2, 3}
	b := a // only copy value // array
	b[1] = 100
	fmt.Println(a) // 1 2 3
	fmt.Println(b) // 1 100 3
	c := &a        // point to a data // pointer
	fmt.Println(c) // &[1,2,3]
	c[1] = 100
	fmt.Println(a) // 1 100 3
	// slices not fixed size
	// create
	x := []int{1, 2, 3}
	fmt.Println(x)
	// build in
	fmt.Println(len(x))
	fmt.Println(cap(x))
	hi := make([]int, 3, 100) //object , len, cap
	fmt.Println(hi)           // [0,0,0]
	hi = append(hi, 1)        // create new array with double size
	hi = append(hi, 2, 3, 4)
	hi = append(hi, []int{4, 5, 6}...) // 1 2 3 4 4 5 6
	hihi := append(hi[:2], hi[3:]...)
	fmt.Println(hihi) // 1 2 4 4
	fmt.Println(hi)   // 1 2 4 4 4 5 6

	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2) // slice1 : 1,2,3  slice2 : 1,2
	// work with
	a1 := []int{1, 2, 3, 4, 5}
	b1 := a1
	b1[1] = 100
	fmt.Println(a1) // 1 100 3 4 5
	fmt.Println(b1) // 1 100 3 4 5

	a2 := [...]int{1, 2, 3, 4, 5, 6, 7}
	suba2 := a2[1:4]
	a2[2] = 42
	fmt.Println(suba2) // 2 42 4

}
