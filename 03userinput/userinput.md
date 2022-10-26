# User input in Go

### Libraries
In order to use input and output in Go, we need to use a few libraries
* [fmt](https://pkg.go.dev/fmt)  (we already used this)
* [bufio](https://pkg.go.dev/bufio)
* [os](https://pkg.go.dev/os)

You can use [pkg.go.dev](https://pkg.go.dev/) to find libraries.

### Importing
You dont need to impot libraries, Go will do it automatically when you save your work.
```go
import (
	"bufio"
	"fmt"
	"os"
) 
```

### Input
**Lets take an example of inputting the rating of pizza and to output the rating.**
```go
fmt.Println("Please enter rating of the pizza")
```

To input you create a variable. Here, we can create a variable `reader`(which reads the input).Now we need to create a new reader using bufio library 
`bufio.NewReader(os.Stdin)`. Note that Stdin and NewReader have the first characters uppercase,denoting it can be used publically.
to input we need to use **Os** library, and use `Stdin`(standard input) function. 
```go
reader := bufio.NewReader(os.Stdin)
```

Now we need to store the input. so we create a `input` variable to store the value which we have entered.
for reading the input, we need to use `ReadString` function. This function returns two values, `value` and `error`.
Since we do not need the error, we place a `-` (underscore) to ignore it.
we also have to provide till what do we want to read in the argument.
Here, we want to read until we come accross `\n`, which is the end of string.

```go
input,_ := reader.ReadString('\n') 
```
Now we can output it to user.
```go
fmt.Println("the rating you have given us is: ",input)

```
the entire code will look somewhat like this:
```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {


	reader := bufio.NewReader(os.Stdin)
	//declaring what our reader is
	fmt.Println("Please enter rating for our Pizza: ")
	//comma ok || comma error

	input, _ := reader.ReadString('\n')
	//save whatever the reader reads in input
	// reader.ReadString('read till what?')
	fmt.Println("Thanks for rating, ", input)
}

```

### Coffee ordering program.
Now lets createa program which promts the user to enter coffee name and thanks them for ordering coffee with name that they ordered.

#### pseudocode
                prompt to enter coffee
                        |
                input what the user enters
                        |
                store it in a variable
                        |
                output thanks for ordering
                
### Code
* prompt user to enter coffee which they want
```go
fmt.Println("Please enter the coffee which you want to order")
```
* read the input using bufio and os library. let us create a variable `cin` 
```go
cin := bufio.NewReader(os.Stdin)
```
* Now store the input in variable `coffee` and ignore the error using _. 
Use the ReadString function and read till \n.
```go
coffee,_ := cin.ReadString('\n)
```
* Now output the value
```go
fmt.Println("thank you for ordering : ",coffee)
```
Done!

Your entire code should look like this:
```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Please enter coffee which you want to order")
	cin := bufio.NewReader(os.Stdin)

	coffee, _ := cin.ReadString('\n')
	fmt.Println("thank you for ordering: ", coffee)

}

```



