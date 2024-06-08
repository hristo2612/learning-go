package main

import "fmt"

func main() {
	tutorial1()
	fmt.Println()
}

func tutorial1() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum ++
		fmt.Println(sum)
	}
}
