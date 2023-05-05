# Functions in Go
A function is a block of code that’s assigned a name, and contains some instructions.

In the “Hello, World!” example we created a main function, which is the entry point of the program.

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```
That’s a special function.

Usually we define functions with a custom name:
```go
func doSomething() {

}
```
and then you can call them, like this:
```go
doSomething()
```
A function can accept parameters, and we have to set the type of the parameters like this:

```go
func doSomething(a int, b int) {

}

doSomething(1, 2)
```
a and b are the names we associate to the parameters internally to the function.

A function can return a value, like this:
```go
func sumTwoNumbers(a int, b int) int {
	return a + b
}

result := sumTwoNumbers(1, 2)
```
Note that we specified the return value type.

A function in Go can return more than one value:

```go
func performOperations(a int, b int) (int, int) {
	return a + b, a - b
}

sum, diff := performOperations(1, 2)
```

It’s interesting because many languages only allow one return value.

Any variable defined inside the function is local to the function.

A function can also accept an unlimited number of parameters, and in this case we call it a variadic function:

```go
func sumNumbers(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

total := sumNumbers(1, 2, 3, 4)
```