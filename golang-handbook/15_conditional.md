# Conditionals in Go
We use the if statement to execute different instructions depending on a condition:

```go
if age < 18 {
	//underage
}
```
The else part is optional:

```go
if age < 18 {
	//underage
} else {
  //adult
}
```
and can be combined with other ifs:

```go
if age < 12 {
	//child
} else if age < 18  {
  //teen
} else {
	//adult
}
```
If you define any variable inside the if, that’s only visible inside the if (same applies to else and anywhere you open a new block with `{}`).

If you’re going to have many different if statements to check a single condition, it’s probably better to use switch:

```go
switch age {
case 0: fmt.Println("Zero years old")
case 1: fmt.Println("One year old")
case 2: fmt.Println("Two years old")
case 3: fmt.Println("Three years old")
case 4: fmt.Println("Four years old")
default: fmt.Println(i + " years old")
}
```
Compared to C, JavaScript, and other languages, you don’t need to have a break after each case.