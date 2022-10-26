# Variables,Types and constants
We will learn about variables,types and constants in this module
Find the code to practice and learn more in  [variables.go](https://github.com/adityakode/Go/blob/main/variables/variables.go)

### Setup
Follow the standard procedure to create a module and a `main.go` file.
If new, you can see the standard procedure [here]().

### About variables
Variables in Golang are  mostly same as any other language.
There are different types of data types like

*  int - uint8, uint32, uint64, int 
* float - float32, float64
* strings
* bool
* complex

### How to declare variables?

### There are three ways to declare variables:
* Default - _var -varname -data type_
* Implicit type
* Using the **walrus** operator


### Default type
We use the variable-variable name-data type syntax to declare a variable.
for example,
```go
var stringVar string = "Golang"
```
here, we can declare the datatype of the variable and assign its value.
Same can be done with other variables.
```go
	// var varname type
	var username string = "Go"
	//it shows error because you didnt use the variable,just declared it
	fmt.Println(username)
	//use fp (shortcut)+tab for writing fmt.Println("")
	fmt.Printf("type of the variable is  %T \n", username)
	// %T prints the type of variable in a c-style output statement (printf).
```

### Implicit type
In the implicit style declaration, we declare name and assign value.
the lexer automatically assigns data type to the variable.

```go
var stringVar = "Go"
//we omit writing the datatype
```
However, we must take note that the variable cant be converted in any other value since the lexer has detected the datatype based on value assignedto it.
```go
	//Implicit type:
	// var varname = value
	var a = "hello"
	//it automatically decides datatype based upon the value that is assigned to it.
	fmt.Println(a)
```

### Walrus operator
Walrus operator is **:=**. It can be used to assign value directly to variable name.
Writing datatype and _var_ keyword can be ommited.
```go
stringVar := "Golang"
```
It is to be noted that you can use the walrus operator to assingn value to  any variables inside any module.
However you cannot use this globally,but other two styles can be used.

### Constants
"constant" keyword is used like any in other programming language.
addding const keyword means it cant change any value.
```go
const var smallInt uint8 = 255
````
### Capital Character
if the first character of the variable is capital, then the variable can be used in any module, or like in other languages, publically.
```go
 const var LwToken string = "aeiou"
```
is same as putting 'public' keyword in statically typed anguages.

### Size table
As we learnt the types of datatypes, each of them have a specific capacity of holding data.


![image](https://user-images.githubusercontent.com/105551807/198051267-112fda0e-2d90-4fba-9ca2-d8f989dfa79b.png)
