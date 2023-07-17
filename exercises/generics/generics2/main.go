// generics2
// Make me compile!

package main

import "fmt"

type Number interface {
	Add(n Number) Number
}

type Int int
type Float float64

func (i Int) Add(n Number) Number {
	return i + n.(Int)
}

func (f Float) Add(n Number) Number {
	return f + n.(Float)
}

func addNumbers(n1, n2 Number) Number {
	return n1.Add(n2)
}

func main() {
	fmt.Println(addNumbers(Int(1), Int(2)))
	fmt.Println(addNumbers(Float(1.0), Float(2.3)))
}
