package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	//declaring what our reader is
	fmt.Println("Please enter rating for our Pizza: ")
	//comma ok || comma error

	input, _ := reader.ReadString('\n')
	//save whatever the reader reads in input
	// reader.ReadString('read till what?')
	fmt.Println("Thanks for rating, ", input)
}
