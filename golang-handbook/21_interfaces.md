# Interfaces in Go
An interface is a type that defines one or more method signatures.

Methods are not implemented, just their signature: the name, parameter types and return value type.

Something like this:
```go
type Speaker interface {
	Speak()
}
```
Now you could have a function accept any type that implements all the methods defined by the interface:

```go
func SaySomething(s Speaker) {
	s.Speak()
}
```
And we can pass it any struct that implements those methods:

```go
type Speaker interface {
	Speak()
}

type Person struct {
	Name string
	Age int
}

func (p Person) Speak() {
	fmt.Println("Hello from " + p.Name)
}

func SaySomething(s Speaker) {
	s.Speak()
}

func main() {
	flavio := Person{Age: 39, Name: "Flavio"}
	SaySomething(flavio)
}
```