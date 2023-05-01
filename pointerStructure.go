package main

import "fmt"

type pointerSructure struct {
	Name    string
	Surname string
	Height  int64
}

func createStructure(name, surname string, height int64) *pointerSructure {
	if height > 300 {
		height = 0
	}
	return &pointerSructure{name, surname, height}
}

func returnStructure(name, surname string, height int64) pointerSructure {
	if height > 300 {
		height = 0
	}
	return pointerSructure{name, surname, height}
}

func main() {

	var s1 = createStructure("Alex", "Zamuraev", 45)
	s2 := returnStructure("Alex", "Zamuraev", 45)

	fmt.Println((*s1).Name)
	fmt.Println(s2.Name)

	fmt.Println(s1)
	fmt.Println(s2)

}
