# Arrays in Go
Arrays are a sequence of items of a single type.

We define an array in this way:

```go
var myArray [3]string //an array of 3 strings
```

and you can initialize the array with values using:

```go
var myArray = [3]string{"First", "Second", "Third"}
```

In this case you can also let Go do some work and count the items for you:

```go
var myArray = [...]string{"First", "Second", "Third"}
```
An array can only contain values of the same type.

The array cannot be resized – you have to explicitly define the length of an array in Go. That’s part of the type of an array. Also, you cannot use a variable to set the length of the array.

Due to this limitation, arrays are rarely used directly in Go. Instead we use slices (more on them later). Slices use arrays under the hood, so it’s still necessary to know how they work.

You can access an item in the array with the square brackets notation we already used in strings to access a single character:

```go
myArray[0] //indexes start at 0
myArray[1]
```
You can set a new value for a specific position in the array:

```go
myArray[2] = "Another"
```
And you can get the length of an array using the len() function:

```go
len(myArray)
```
Arrays are value types. This means copying an array:

```go
anotherArray := myArray
```
or passing an array to a function, or returning it from a function, creates a copy of the original array.

This is different from other programming languages out there.

Let’s make a simple example where we assign a new value to an array item after copying it. See, the copy doesn't change:

```go
var myArray = [3]string{"First", "Second", "Third"}
myArrayCopy := myArray
myArray[2] = "Another"

myArray[2]     //"Another"
myArrayCopy[2] //"Third"
```

Remember you can only add a single type of items in an array, so setting the 
```
myArray[2] = 2
```
for example will raise an error.

Low-level elements are stored continuously in memory.