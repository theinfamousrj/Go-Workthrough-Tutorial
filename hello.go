//Go likes to work in packages
package main

//This imports fmt which contains the Println() function
import "fmt"

//Typical main function
func main() {

	//This is the line that actually prints the message
	fmt.Println("Hello world!")

}

//Instructions for running:
//(assuming you have Go installed)
//0.save this file as hello.go
//1.open terminal or cmd
//2.cd to the directory that contains this file
//3.type 'go run hello.go'
