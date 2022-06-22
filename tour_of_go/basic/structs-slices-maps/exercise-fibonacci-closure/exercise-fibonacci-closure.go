package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a1 := 0
	a2 := 1
	return func() (result int) {
		result = a1
		temp := a2
		a2 = a1 + a2
		a1 = temp
		return
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
