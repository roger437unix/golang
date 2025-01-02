package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	nome := "tux"
	idade := 25
	pPerson := newPerson(&nome, &idade)
	fmt.Println(*pPerson)
	fmt.Println(pPerson.name)
	fmt.Println(pPerson.age)
}

func newPerson(pName *string, pAge *int) *person {
	pessoa1 := person{name: *pName, age: *pAge}
	return &pessoa1
}
