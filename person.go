package main

import(
	"fmt"
)

type person struct{
	firstName string
	lastName string	
}

func (p *person) updateFirstName(name string){
	(*p).firstName = name
}

func (p person) print(){
	fmt.Printf("First Name : %v\n", p)
	fmt.Printf("First Name : %+v\n", p)
	fmt.Printf("First Name : %#v\n", p)
	fmt.Printf("First Name : %T\n", p)
}