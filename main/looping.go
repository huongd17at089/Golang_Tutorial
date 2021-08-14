package main

import "fmt"

/* summary
for
- 3 simple type
- break, continue, label
- looping over collection : for k, v := range collection ( array, slide, string , channel)

*/

func main() {
	//simple loop
	for i := 0; i < 5; i++ { // i scope for loop
		fmt.Println(i)
	}

	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i)
	}

	i := 0 // i scope func
	for ; i < 5; i++ {
		fmt.Println(i)
	}

	for i < 5 { // while loop

	}

	//exitting early

	for { //inf loop
		if i == 10 {
			break // end loop
		}
		if i == 2 {
			continue // exit interation
		}
	}

	// nested loop
loop: //label
	for i := 0; i < 10; i++ {
		for j := 11; j < 20; j++ {
			fmt.Println(i * j)
			break      // break j loop
			break loop // break all loop

		}
	}

	//lopping through collection
	//slide-array-map-string-channel
	s := []int{1, 2, 3}
	fmt.Println(s)
	for k, v := range s {
		fmt.Println(k, v) // index - value
	}

	for _, v := range s {
		fmt.Println(v) //  value
	}

}
