package basics

import "fmt"

// This function is exported
// Because of the capital letter
func SomeFunction() {
	fmt.Println("Hello this is an exported package!")
}

// This variable is exported
// Because of the capital letter
var SomeVariable = "Exported Variable"

// Function continuation.
func ContinuatedFunc(arg1, arg2 int, arg3 string, arg4, arg5 int) {
	// arg1 to arg2 are int
	// arg3 is a string
	// arg4 to arg5 are int
	fmt.Println(arg1, arg2, arg3, arg4, arg5)
}

// Variables
var Value1, Value2, Value3 int = 1, 3, 5

// Short variable declarations
func ShortVarDeclaration() int {
	shortVar := 55
	return shortVar
}
