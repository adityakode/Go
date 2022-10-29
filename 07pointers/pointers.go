package main

import "fmt"

func main() {
	fmt.Println("Welcome to the pointers")
	mynum := 23
	var ptr *int = &mynum
	fmt.Println("the pointer stores  address ", ptr)
	//ptr stores the address of mynum
	// in order to see the value of address the pointer is holding, we need to use "*" (dereferencing operator)
	fmt.Println("the value whose address ptr points to is: ", *ptr)

	// when we pass a value, a copy of that gets passed. This leads to problems and sometimes when is is out of scope, then the changed value does not get
	// reflected in the var. for this we use pointers.We gorectly pass pointers who store address of the var.
	// So, no doubt of duplication
	//                             *
	//                                / \
	// to  initialise the pointer(*int)  to access value whose address the ptr is holding(*ptr)
	*ptr = *ptr + 2
	fmt.Println("the updated value is: ", *ptr)

}
