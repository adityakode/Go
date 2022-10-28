## Building files in Mac, Linux and Windows

Suppose we build an application. But how do we build files that we run it?
For example we get a `.exe` file in windows. How can you build it in other operating systems?
Go gets you here! You can build files from any OS that run in the specified OS.

### env command
in the integrated terminal, type the following command:
```go
go env
```
this will sjow you all the commands. But we are intrested in a command that will look like this:
```
GOOS = "(Your os name)"
// Windows,Linux and Darwin for Mac

```
This command is uded to build files.

### Own OS
If you wanted to build for OS that you are using, just type
```
GOOS= "" go build
```
### Windows executable `.exe`
If you want to build windows `.exe` file, just type
```
GOOS="Windows" go build
```
This creates a .exe file. This can be built from even linux and MAc!

### Linux
```
GOOS="Linux" go build

```

### Mac 
```
GOOS="Darwin" go build
```
