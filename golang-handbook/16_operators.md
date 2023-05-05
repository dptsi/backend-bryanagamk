# Operators in Go
We've used some operators so far in our code examples, like `=`, `:=` and `<`.

Let’s talk a bit more about them.

We have assignment operators `=` and `:=` we use to declare and initialize variables:

```go
var a = 1
b := 1
```

We have comparison operators `==` and `!=` that take 2 arguments and return a boolean:

```go
var num = 1
num == 1 //true
num != 1 //false
and <, <=, >, >=:

var num = 1
num > 1 //false
num >= 1 //true
num < 1 //false
num <= 1 //true
```
We have binary (require two arguments) arithmetic operators, like `+, -, *, /, %`.

```go
1 + 1 //2
1 - 1 //0
1 * 2 //2
2 / 2 //1
2 % 2 //0
```
+ can also join strings:

```go
"a" + "b" //"ab"
```
We have unary operators ++ and -- to increment or decrement a number:

```go
var num = 1
num++ // num == 2
num-- // num == 1
```
Note that unlike C or JavaScript we can’t prepend them to a number like `++num`. Also, the operation does not return any value.

We have boolean operators that help us with making decisions based on true and false values: `&&`, `||` and `!`:

```go
true && true  //true
true && false //false
true || false //true
false || false //false
!true  //false
!false //true
```
Those are the main ones.