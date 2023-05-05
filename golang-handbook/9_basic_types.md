# Basic Types in Go
Go is a typed language.

We saw how you can declare a variable, specifying its type:

```go
var age int
```
Or you can let Go infer the type from the initial value assigned:

```go
var age = 10
```
The basic types in Go are:

* Integers (`int`, `int8`, `int16`, `int32`, `rune`, `int64`, `uint`, `uintptr`, `uint8`, `uint16`, `uint64`)
* Floats (`float32`, `float64`), useful to represent decimals
Complex types (`complex64`, `complex128`), useful in math
* Byte (`byte`), represents a single ASCII character
* Strings (`string`), a set of bytes
* Booleans (`bool`), either true or false

We have a lot of different types to represent integers. You will use `int` most of the time, and you might choose a more specialized one for optimization (not something you need to think about when you are just learning).

An `int` type will default to be 64 bits when used on a 64 bit system, 32 bits on a 32 bit system, and so on.

`uint` is an `int` that’s unsigned, and you can use this to double the amount of values you can store if you know the number is not going to be negative.

All the above basic types are value types, which means they are passed by value to functions when passed as parameters, or when returned from functions.