package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	grandparent, parent := 0, 1
	return func() int {
		grandparent, parent = parent, parent+grandparent
		return grandparent
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
