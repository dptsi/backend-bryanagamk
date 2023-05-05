# Variables in Go
One of the first things you do in a programming language is defining a variable.

In Go we define variables using `var`:

```go 
var age = 20
```

You can define variables at the package level:

```go
package main

import "fmt"

var age = 20

func main() {
	fmt.Println("Hello, World!")
}
```

or inside a function:

```go
package main

import "fmt"

func main() {
	var age = 20

	fmt.Println("Hello, World!")
}
```

Defined at the package level, a variable is visible across all the files that compose the package. A package can be composed of multiple files, you just need to create another file and use the same package name at the top.

Defined at the function level, a variable is visible only within that function. It’s initialized when the function is called, and destroyed when the function ends the execution.

In the example we used:

```go 
var age = 20
```
we assign the value 20 to `age`.

This makes Go determine that the type of the variable `age` is `int`.

We’ll see more about types later, but you should know there are many different ones, starting with `int`, `string`, and `bool`.

We can also declare a variable without an existing value, but in this case we must set the type like this:

```go
var age int
var name string
var done bool
```

When you know the value, you typically use the short variable declaration with the `:=` operator:

```go
age := 10
name := "Roger"
```

For the name of the variable you can use letters, digits, and the underscore `_` as long as the name starts with a character or `_`.

Names are **case sensitive**.

If the name is long, it’s common to use camelCase. So to indicate the name of the car we use `carName`.

You can assign a new value to a variable with the assignment operator `=`

```go
var age int
age = 10
age = 11
```

If you have a variable that never changes during the program you can declare it as a constant using `const`:

```go
const age = 10
```

You can declare multiple variables on a single line:

```go
var age, name
```

and initialize them too:

```go

var age, name = 10, "Roger"

//or

age, name := 10, "Roger"
```

Declared variables that are not used in the program raise an error and the program does not compile.

If you declare a variable without initializing it to a value, it is assigned a value automatically that depends on the type – for example an integer is 0 and a string is an empty string.