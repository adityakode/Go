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
