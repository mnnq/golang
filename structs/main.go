package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	/*
			This is another way of creating a new struct
			mat := person{"Matias", "Novoa"}
			c := person{firstName: "asd", lastName: "ff"}
			fmt.Printf("%+v", matias)
			var matias person
			matias.firstName = "Matias"
			matias.lastName = "N"

			jimPointer := &jim
		jimPointer.updateName("asd")
	*/

	//Embedded structs
	jim := person{

		firstName: "Jum",
		lastName:  "sarasa",
		contactInfo: contactInfo{
			email:   "asd@gmail.com",
			zipCode: 2344,
		}}
	jim.updateName("sadasdasdsda")
	jim.print()

}

/*
Slices behave different with pointers, if i have a func that recieves a slice it will directly edit the slice, you dont need a pointer to the slice

behind the scenes the slice its composed of two things (uses two memory spaces) the slice data structure and the array of the slice, so when you use a function
that needs a slice GO will create a copy of the slice data structure and will access the memory where the slice array is stored and modify it

Reference Types (dont need pointers): slices, maps, channels, pointers, functions

value types (you need pointers to change this in a function): int, float, string, bool, structs

*/

//*person is a type of pointer that points to a person
func (p *person) updateName(newFirstName string) {
	//*p is an operator, we want to manipulate the value the pointer is refernecing
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v \n", p)
}
