# Methods in Go
You can assign a function to a struct, and in this case we call it a method.

Example:
```go
type Person struct {
	Name string
	Age int
}

func (p Person) Speak() {
	fmt.Println("Hello from " + p.Name)
}

func main() {
	flavio := Person{Age: 39, Name: "Flavio"}
	flavio.Speak()
}
```
You can declare methods to be pointer receiver or value receiver.

The above example shows a value receiver. It receives a copy of the struct instance.

This would be a pointer receiver that receives the pointer to the struct instance:

```go
func (p *Person) Speak() {
	fmt.Println("Hello from " + p.Name)
}
```