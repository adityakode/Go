## Pointers in Go

Pointers are very necessary n any programming language

### What is need of pointers?
In some cases we need to pass a variable to a function. In these cases, the variable is not passed,instead, a copy of variable is 
passed.This leads to problems such as the var is not updated due to scope limit.
To tackle these, we use pointers which ensure that the values are passed directly and changes are made in the orignal value itself.

## Using pointers

### the "*" operator
This operator is used in two ways:
* To initialise a pointer
* to dereference the value

For example we want to initialise a pointer `ptr` that holds address of an *int* variable

```go
var ptr *int
//this initialised a pointer that can store address
//of a integer
```

### "&"(address of) operator

How do we store address in a pointer? we store them using `&` operator.
```go
mynum := 32

var ptr *int = &mynum
//address of mynum for eg 0xc0000a6058 is stored in ptr
```
### Dereferencing
Suppose we want to print the value that is stored in the address the pointer is holding

```go
fmt.Println("the value is, ",*ptr)
//this prints 32

```


### You can even perform operations on pointers
```go
*ptr = *ptr +2
fmt.Println(*ptr)
//output = 34
```
