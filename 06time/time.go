package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time in golang")

	//for this, we use time package/library
	presentTime := time.Now()
	fmt.Println(presentTime)

	// this shows the present time
	// but it looks too long . to print only Dayand date, we use format
	fmt.Println(presentTime.Format("01-02-2006 Monday"))
	// to use time , you have to use a specific no in format that is 15:04:05
	// specific date is 01-02-2006
}
