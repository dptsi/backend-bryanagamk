# Structs in Go
A struct is a type that contains one or more variables. It’s like a collection of variables. We call them fields. And they can have different types.

Here’s an example of a struct definition:
```go
type Person struct {
	Name string
	Age int
}
```
Note that I used uppercase names for the fields, otherwise those will be private to the package. And when you pass the struct to a function provided by another package, like the ones we use to work with JSON or database, those fields cannot be accessed.

Once we define a struct we can initialize a variable with that type:

```go
flavio := Person{"Flavio", 39}
```
and we can access the individual fields using the dot syntax:

```go
flavio.Age //39
flavio.Name //"Flavio"
```

You can also initialize a new variable from a struct in this way:

```go
flavio := Person{Age: 39, Name: "Flavio"}
This lets you initialize only one field, too:

flavio := Person{Age: 39}
or even initialize it without any value:

flavio := Person{}

//or

var flavio Person
and set the values later:

flavio.Name = "Flavio"
flavio.Age = 39
```

Structs are useful because you can group unrelated data and pass it around to/from functions, store in a slice, and more.

Once defined, a struct is a type like int or string and this means you can use it inside other structs, too:

```go
type FullName struct {
	FirstName string
	LastName string
}

type Person struct {
	Name FullName
	Age int
}
```