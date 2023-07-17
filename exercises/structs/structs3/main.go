// structs3
// Make me compile!
package main

import "fmt"

type FullName func(string, string) string
type Person struct {
	firstName string
	lastName  string
	fullName  FullName
}

func main() {
	person := Person{firstName: "Maurício",
		lastName: "Antunes",
		fullName: func(firstName string, lastName string) string {
			return firstName + " " + lastName
		},
	}
	fmt.Printf("Person full name is: %s\n", person.fullName(person.firstName, person.lastName)) // here it must output Person full name is: Maurício Antunes
}
