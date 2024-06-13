package basics

import "fmt"

// A simple for loop...
func SimpleForLoop() {
	for i := 0; i <= 10; i++ {
		fmt.Print(i)
	}
	fmt.Println()
}

// For is also a while loop..
func SimpleWhileLoop() {
	sum := 0
	for sum < 10 {
		sum++
	}
}
