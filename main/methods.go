package main

import "fmt"

/*
- func execute in context
- func ( type) <name>(){
}
- pass value, pass pointer
*/

func main() {
	g := greeter{
		greating: "hello",
		name:     "go",
	}

	g.greet()
}

type greeter struct {
	greating string
	name     string
}

func (g greeter) greet() { // copy object value
	fmt.Println(g.greating, g.name)
}

func (g *greeter) greet2() { // pass pointer
	fmt.Println(g.greating, g.name)
}
