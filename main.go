package main

import (
	"fmt"

	"github.com/hristo2612/learning-go/basics"
)

func main() {
	fmt.Println("Hello from GO!")
	basics.SomeFunction()
	fmt.Println(basics.SomeVariable)
	basics.ContinuatedFunc(11, 12, "hello", 5, 6)
	fmt.Println(basics.Value1, basics.Value2, basics.Value3)
	fmt.Println(basics.ShortVarDeclaration())
}
