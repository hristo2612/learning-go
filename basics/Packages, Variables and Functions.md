# Table of Contents

- [Modules](#modules)
- [Packages](#packages)
- [Functions](#functions)
- [Variables](#variables)

# Modules

- Modules compromise of either a project or library.
- They are initialised using `go mod init [nameofmodule]` in the root directory of the project / library
- If you are planning to share those modules with other go projects the name of the module can also be the **URL of your github repo**. Example:
  ```
  go mod init github.com/hristo2612/learning-go
  ```
- Once the module is initialised it creates a `go.mod` file in the root.

# Packages

- `package main` is the main starting point of the app.
- **Packages** are imported like so:

  ```go
  import("fmt", "math/rand")
  ```

- So for example if in a file you specify:

  ```go
  package mypizza

  import ("fmt")

  func Time() {
      fmt.Println("It's pizza time!")
  }
  ```

- Then we can import our new **pizza package** like so:

  ```go
  import ("mypizza")

  mypizza.Time()
  ```

# Functions

- Functions are defined like so:

  ```go
  func myBeautifulFunction() {}
  ```

- Packages can export certain functions by **naming the functions with a capital letter**.

  ```go
  package basics

  // This one is exported.
  // It can be used like so:
  // basics.SomeFunc()
  func SomeFunc() {}

  // This one is private.
  func someFunc() {}
  ```

  > This also works for exporting variables.

- **Function arguments** are separated by a _comma_ with the types coming **after** the name.

  ```go
  func someFunc(arg1 int, arg2 int) {}
  ```

- When two or more arguments share the same type, you can put the type only on the last argument:

  ```go
  func ContinuatedFunc(arg1, arg2 int, arg3 string, arg4, arg5 int) {
      // arg1 to arg2 are int
      // arg3 is a string
      // arg4 to arg5 are int
      fmt.Println(arg1, arg2, arg3, arg4, arg5)
  }
  ```

- Functions can return multiple results:

  ```go
  func ReturnMultipleResults(arg1 int, arg2 string) {
      return arg1, arg2
  }
  ```

  Usage:

  ```go
  arg1, arg2 := ReturnMultiplerResults(4, "hello")
  ```

# Variables

- Variables can be created with the `var` keyword.
- We can assign multiple vars with one `var` keyword:

  ```go
  var value1, value2, value3 int = 1, 3, 5
  ```

- **Short variable declarations** can be used inside functions only:

  ```go
  func someFunc() {
      a := 1
      // Is the same as..
      var b = 1
  }
  ```

- IF we don't give an explicit value to a `var` it receives a Zero value:

  ```go
  var someString string // = ""
  var someNum int // = 0
  var someBool bool // = false
  ```

- Just like in JavaScript we can declare a constant variable with `const`

  ```go
  const someValue = "I cannot be changed! MUHAHAHA"
  ```
