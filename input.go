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
	//Reads the input
	scan := bufio.NewScanner(os.Stdin)
	//Scan advances the Scanner to the next token, which will then be available through the Bytes or Text method.
	scan.Scan()
	//Prints on the terminal
	fmt.Println("Hi " + scan.Text())
}