package main

import "fmt"

/*
summary
map
- key : value
- make()
- check exist by ok
- key same type, val same type
- copy same underlying data

struct
- colect any type
- field
- normal type and anonymous type
- variable type
- no inheritance
- can embeding
- tags
*/

type Persion struct {
	Name string //'required max: "10"'
	Age  int
	Abc  []int
}

// embeding

type Doctor struct {
	Persion  //composition relationship = has a
	Hospital string
}

func main() {
	//map
	//create
	m := map[string]int{
		"hanoi": 1,
		"vn":    2,
	}
	// m := map[[]int]string{} error
	mep := map[[3]int]string{}
	fmt.Println(m, mep)
	fmt.Println(m["hanoi"]) // 1
	m["HCM"] = 3            // add key
	pop, ok := m["hanoi"]
	fmt.Println(pop, ok) // 1, true = key in map
	//buil in
	delete(m, "HCM") // del key
	fmt.Println(len(m))
	makemap := make(map[string]int)
	makemap = m
	fmt.Println(makemap)
	//manipulation
	a := map[string]int{
		"hanoi": 1,
		"vn":    2,
	}
	b := a
	delete(b, "vn")
	fmt.Println(b) // hanoi :1
	fmt.Println(a) // hanoi : 1

	//struct
	//create
	persion := Persion{
		Name: "hihi",
		Age:  22,
		Abc:  []int{1, 2, 3},
	}
	fmt.Print(persion)
	fmt.Println(persion.Age)

	doctor := struct {
		name string
	}{name: "hu"}
	fmt.Println(doctor)
	other := doctor
	other.name = "hi"
	fmt.Println(doctor.name) // hu
	fmt.Println(other.name)  //hi

	otherx := &doctor
	fmt.Println(otherx)
	//naming

	//embeding
	d := Doctor{}
	d.Name = "huhi"
	d.Hospital = "lala"
	fmt.Println(d.Name)
	//tags
	/*

	 */

}
