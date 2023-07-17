// slices1
// Make me compile!

package main

import "fmt"

func main() {
	a := make([]int, 5)
	fmt.Println("length of 'a':", len(a))
	fmt.Println("capacity of 'a':", cap(a))
}
