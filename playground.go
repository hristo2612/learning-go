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
	fmt.Println()
	tutorial6()
	fmt.Println()
	tutorial7()
	fmt.Println()
	tutorial8()
	fmt.Println()
	tutorial9()
}

func tutorial1() {
	fmt.Println("1. In GO exported names always begin with a capital letter. For example math.Pi")
	fmt.Println("math.Pi =", math.Pi)
}

func add(x int, y int) int {
	return x + y
}

func tutorial2() {
	fmt.Println(`2. Functions can take zero or more arguments:
func add(x int, y int) int { 
 return x + y
}
add(5, 10)`)
	fmt.Println("Result:", add(5, 10))
}

func add2(x int, y int) int {
	return x + y
}

func tutorial3() {
	fmt.Println(`3. When two or more consecutive arguments share the same type we can add the type only to the last argument like so:
func add(x, y int) int { 
 return x + y
}
add(5, 10)`)
	fmt.Println("Result:", add2(5, 10))
}

func swap(x, y string) (string, string) {
	return y, x
}

func tutorial4() {
	fmt.Println(`4. A function can return any number of results:
func swap(x, y string) (string, string) {
 return y, x
}
a, b := swap("X", "Y")`)
	a, b := swap("X", "Y")
	fmt.Println("a =", a)
	fmt.Println("b =", b)
}

func tutorial5() {
	var python, java, csharp bool
	fmt.Println(`5. The 'var' statement declares a list of variables. The type is last as in function arguments:
var python, java, csharp bool
fmt.Println(python, java, csharp)
Result:`, python, java, csharp)
}

func tutorial6() {
	var python, java, csharp = true, 4, "csharp"
	fmt.Println(`6. The 'var' declaration can also include initialisers like so. The type is inferred from the value:
var python, java, csharp = true, 4, "csharp"
fmt.Println(python, java, csharp)
Result:`, python, java, csharp)
}

func tutorial7() {
	fmt.Println(`7. Inside a function instead of declaring vars with the 'var' keyword we can use := directly like so:
var i = 3
j := 3
Those are the same thing.`)
}

func tutorial8() {
	fmt.Println(`8. The expression T(v) converts value 'v' to type 'T':
someIntNumber := 5
someFloatNumber := float64(someIntNumber)`)
	someIntNumber := 5
	someFloatNumber := float64(someIntNumber)
	fmt.Println(someFloatNumber)
}

func tutorial9() {
	fmt.Println(`9. Constant variables are defined with 'const' keyword:
const penguin = "penguin"`)
}
