# Loops in Go
One of Go’s best features is to give you fewer choices.

We have one loop statement: for.

You can use it like this:

```go
for i := 0; i < 10; i++ {
	fmt.Println(i)
}
```

We first initialize a loop variable, then we set the condition we check for with each iteration to decide if the loop should end. Finally we have the post statement, executed at the end of each iteration, which in this case increments i.

`i++` increments the `i` variable.

The `<` operator is used to compare i to the number 10 and returns true or false, determining if the loop body should be executed or not.

We don’t need parentheses around this block, unlike other languages like C or JavaScript.

Other languages offer different kind of loop structures, but Go only has this one. We can simulate a while loop, if you’re familiar with a language that has it, like this:

```go
i := 0

for i < 10 {
	fmt.Println(i)
  i++
}
```

We can also completely omit the condition and use break to end the loop when we want:

```go
i := 0

for {
	fmt.Println(i)

	if i < 10 {
		break
	}

  i++
}
```
I used a if statement inside the loop body, but we haven’t seen conditionals yet! We’ll do that next.

One thing I want to introduce now is range.

We can use for to iterate through an array using this syntax:

```go
numbers := []int{1, 2, 3}

for i, num := range numbers {
	fmt.Printf("%d: %d\n", i, num)
}

//0: 1
//1: 2
//2: 3
```
Note: I used `fmt.Printf()` which allows us to print any value to the terminal using the verbs `%d` which mean decimal integer and `\n` means add a line terminator.

It’s common to use this syntax when you don’t need to use the index:

```go
for _, num := range numbers {
  //...
}
```

We're using the special `_` character that means “ignore this” to avoid the Go compiler raising an error saying “you’re not using the i variable!”.