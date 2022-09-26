package main

import (
	"fmt"
)

type person struct {
	FirstName string
	LastName  string
	contactInfo
}
type contactInfo struct {
	email string
	Phone int
}

//Without Pointers
// func main() {
// 	me := person{
// 		FirstName: "Eratus",
// 		LastName:  "Thambo",
// 		contactInfo: contactInfo{
// 			email: "kj5574@yahoo",
// 			Phone: 2604147822,
// 		},
// 	}

// 	x := me.updateName("Tony")
// 	x.print()
// 	me.print()

// }

// func (p person) print() {
// 	fmt.Printf("%+v", p)
// }

// func (p person) updateName(newFName string) person {
// 	p.FirstName = newFName
// 	return p

// }
//With Pointers
// Referencing to a variable
// & -> memory address
// * -> value IN memory address &
// Referencing a type
// *(type) here is merely a description of type, pointer of type ie *person -> pointer of type person to be received.
// Slices can be directly updated without using pointers

func main() {
	me := person{
		FirstName: "Eratus",
		LastName:  "Thambo",
		contactInfo: contactInfo{
			email: "kj5574@yahoo",
			Phone: 2604147822,
		},
	}
	me.print()
	mePointer := &me
	mePointer.updateName("Tony")
	me.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointer *person) updateName(newFName string) {
	//() around *pointer turns it to the actual value of the pointer.
	(*pointer).FirstName = newFName
}
