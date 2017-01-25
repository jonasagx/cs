package main

// Libraries used for this example
import (
   "bufio"	//<-- implements buffered I/O.
   "fmt"	//<-- implements formatted I/O with functions analogous to C's printf and scanf.
   "os"		//<-- provides a platform-independent interface to operating system functionality.
)

func main() {
	//Prints the message
	fmt.Print("Write your name then press ENTER, please: ")
	
	//Reads the text input
	var scan = bufio.NewScanner(os.Stdin)
	
	//Check if anything was passed as input and extract it
	scan.Scan()
	//Put the input in a variable of type string
	var name string = scan.Text()
	
	//Prints the input on the terminal
	fmt.Println("Hi " + name)
}