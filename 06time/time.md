## Time in Go

To use time in go , we need `Time` package.
The functions in time lib are listed here : [https://pkg.go.dev/time](https://pkg.go.dev/time)

### Present time
 To find the present time, we use the `.Now()` function
 ```go
 	presentTime := time.Now()
	fmt.Println(presentTime)
 ```
 This prints the current time along with date, time zone.
 But this is not very appealing as we jist wanted to see time and if possible,Date.
 
 For this, we use `Format` and add arguments as Date and time.
 
 In Go, there is a standard example. For eg,
 * Date = 01-02-2006
 * Time = 15:04:05
 * Day = Monday
 
 ### Examples
 
 For example, if you want to Display present Date and Day,
 ```go
 	// but it looks too long . to print only dayand date, we use format
	fmt.Println(presentTime.Format("01-02-2006 Monday"))
 ```
 This will Print todays `Day` and `Day`
 
 
 If we want to print time , date and day,
 
 ```go
 	fmt.Println(presentTime.Format(" 15:04:05 01-02-2006 Monday"))
 ```
 
 ### More functions
 There are many more functions which do various things in the time library.
 We recommend you check out the library and discover more functions yourself : [timee](https://pkg.go.dev/time)
