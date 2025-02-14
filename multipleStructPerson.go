package main

import "fmt"

type name struct {
	fName string
	sName string
}

type person struct {
	Name  name
	empid string
	age   int
}

type node struct {
	Person person
	next   *node
}

func main() {
	p := new(node)
	p.Person.Name.fName = "abc"
	p.Person.Name.sName = "def"

	q := new(node)
	q.Person.Name.fName = "cba"
	q.Person.Name.sName = "fed"
	p.next = q
	fmt.Println(p) // &{{{abc def}  0} 0xc000024140}
}
