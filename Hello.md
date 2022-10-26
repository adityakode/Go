# Hello World 
## Writing your first program in go.

## Installation
To install go , you need to download from [Golang](go.dev) website
after downloading , ***install*** the package.

### Text Editor
Text editor is required to work with go.
Any text editor is fine, but I recommend [**Visual Studio Code**](https://code.visualstudio.com/).
After installing VS Code , you need to add **Go** extension.

### Writing your first program
Create a folder named  `01hello`
int the folder create a file `main.go`
Now we create a package (we will learn more in the subsequent sections). In the Integrated terminal, type the following command.
```
go mod init hello
```
this creates a package hello.(We learn more about packages later or you can checkout **Go Docs** [here](https://go.dev/doc/).

Now we write our hello world code in `main.go`
or you can copy from [here](https://github.com/adityakode/Golang/blob/main/hello.go).



```

package main()
import "fmt"

func main(){
fmt.Println("Hello World")
}
```
### Output
In order to run the code, in the integrated terminal,type the following command:
```
go run main.go
```
and you get the output as follows:

![image](https://user-images.githubusercontent.com/105551807/197955993-94d310ca-d6b6-449e-a0d6-b0174ba6a60c.png)

### Help
if you face some issue,you can refer to go docs or can type the following command in terminal:
```
go help
```

