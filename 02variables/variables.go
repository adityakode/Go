package main

import "fmt"

func main() {
	// var varname type
	var username string = "aditya"
	//it shows error because you didnt use the variable,just declared it
	fmt.Println(username)
	//use fp (shortcut)+tab for writing fmt.Println("")
	fmt.Printf("type of the variable is  %T \n", username)
	// %T prints the type of variable in a c-style output statement (printf).

	fmt.Println("..............................................................")

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("type of the variable is  %T \n", isLoggedIn)

	fmt.Println("..............................................................")

	var smallVal uint8 = 255
	//uint8 holds nums upto 255. if you put 256 it'll be an error
	fmt.Println(smallVal)
	fmt.Printf("type of the variable is  %T \n", smallVal)

	fmt.Println("..............................................................")

	var smallFloat float32 = 255.56891
	//floats are the values after . for eg 32.45322
	//float32 prints 5 values after the decimal
	// if you want more values,you can use float64
	fmt.Println(smallFloat)
	fmt.Printf("type of the variable is  %T \n", smallFloat)

	//now lets talk about default values and aliases.

	var anotherVar int
	fmt.Println(anotherVar)
	//it is uninitialised,still it prints 0 . Go does not collect any garbage value
	fmt.Printf("type of the variable is  %T \n", anotherVar)

	var stringVar string
	fmt.Println(stringVar)

	fmt.Printf("type of the variable is  %T \n", stringVar)
	//string value prints nothing if not initialised(or you can say it prints space)

	//there is another type of declaring variables
	//Implicit type:
	// var varname = value
	fmt.Println("****************************************************************************")
	var a = "hello"
	//it automatically decides datatype based upon the value that is assigned to it.
	fmt.Println(a)

	fmt.Printf("type of the variable is  %T \n", a)
	//if you dint declare the variable type, the lexer declares it for you

	//there is also another style called no-var style
	noVar := 300
	fmt.Println(noVar)
	//the operator is :=
	//you are allowed to use Walrus operator only inside a method. if you try to declare any var using it globally, it'll throw an error
	fmt.Println(LoginToken)
}

//if you declare any var with first letter Capital, it becomes public
//same as writing "public" in front of the keyword.
const LoginToken string = "aeiou"
