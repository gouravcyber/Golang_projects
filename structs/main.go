package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct{
	firstName string
	lastName string
	contact contactInfo
}

func main(){
	jim := person{
		firstName: "Jim",
		lastName: "Mach",
		contact: contactInfo{
			email:"jim_ach@gmail.com",
			zipCode:784824,
		},
	}
	jimPointer :=&jim
	jimPointer.updateName("Jimmmy")
	jim.print()
}

func (p person) print(){
	fmt.Printf("%+v",p)
}

func (pointerToPerson *person) updateName(newFirstName string){
	(*pointerToPerson).firstName = newFirstName
}
