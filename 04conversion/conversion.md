## Conversion in Golang
In the last code, we got the rating of our pizza.If we had to add 1 to the rating, it would throw an error because it parsed it as string.
If new, the last code is [here]()
```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	
)

func main() {
	fmt.Println("Please enter rating")
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for giving rating, ", input)
  }
```

### Solution
To add 1, we must convert the input to number.
We can do this by using `strconv` library, which has a function `parseFloat`
```go
//func strconv.ParseFloat(s string, bitSize int) 
addone = strconv.parseFloat(input,64)
```
parseFloat has two arguments, one input,other is a bit size int.
Still the error persists cause the program comes accross \n.
We need to figure out a way to remove `\n` from the input.

### using "strings" library
`strings` library coms with a `TrimSpace` function, which removes/trims \n.
```go
//func strings.TrimSpace(s string) string
strings.TrimSpace(input)
```
Instead of input, we write the above code.
```go
addone, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	// strconv.ParseFloat(value,64 bit sized) converts string to float
	//this also throws an error because the input is 4\n
	//we need to trin the \n part.For this,we use ths 'strings' library and its

```

### Errors
We need to catch and print error if present, otherwise we just add 1 to input and print it.
To print error, we need to put an `if-else` statement
```go
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("1 added to your rating ", addone+1)
	}
	//if error occurs, print error . or else print rating +1
```

### Done!
The code should look somewhat like [this](https://github.com/adityakode/Go/blob/main/04conversion/conversion.go)
```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Please enter rating")
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for giving rating, ", input)

	addone, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	// strconv.ParseFloat(value,64 bit sized) converts string to float
	//this also throws an error because the input is 4\n
	//we need to trin the \n part.For this,we use ths 'strings' library and its Trimpsace Function
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("1 added to your rating ", addone+1)
	}
	//if error occurs, print error . or else print rating +1
}

```
