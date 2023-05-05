# Maps in Go
A map is a very useful data type in Go.

In other language it’s also called a dictionary or hash map or associative array.

Here’s how you create a map:

```go
agesMap := make(map[string]int)
```

You don’t need to set how many items the map will hold.

You can add a new item to the map in this way:

```go
agesMap["flavio"] = 39
```
You can also initialize the map with values directly using this syntax:

```go
agesMap := map[string]int{"flavio": 39}
```
You can get the value associated with a key using:

```go
age := agesMap["flavio"]
```
You can delete an item from the map using the delete() function in this way:

```go
delete(agesMap, "flavio")
```