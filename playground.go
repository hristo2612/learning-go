package main

import (
	"fmt"
	"math"
)

func main() {
	tutorial1()
	fmt.Println()
	tutorial2()
	fmt.Println()
	tutorial3()
	fmt.Println()
	tutorial4()
	fmt.Println()
	tutorial5()
}

func tutorial1() {
	fmt.Println("1. In GO exported names always begin with a capital letter. For example math.Pi")
	fmt.Println("math.Pi =", math.Pi)
}

func add(x int, y int) int {
	return x + y
}

func tutorial2() {
	fmt.Println("2. Functions can take zero or more arguments:")
	fmt.Println(`func add(x int, y int) int { 
 return x + y
}
add(5, 10)`)
	fmt.Println("Result:", add(5, 10))
}

func add2(x int, y int) int {
	return x + y
}

func tutorial3() {
	fmt.Println("3. When two or more consecutive arguments share the same type we can add the type only to the last argument like so:")
	fmt.Println(`func add(x, y int) int { 
 return x + y
}
add(5, 10)`)
	fmt.Println("Result:", add2(5, 10))
}

func swap(x, y string) (string, string) {
	return y, x
}

func tutorial4() {
	fmt.Println("4. A function can return any number of results:")
	fmt.Println(`func swap(x, y string) (string, string) {
 return y, x
}
a, b := swap("X", "Y")`)
	a, b := swap("X", "Y")
	fmt.Println("a =", a)
	fmt.Println("b =", b)
}

func tutorial5() {

}
