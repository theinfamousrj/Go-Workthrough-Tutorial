package main

import "fmt"

//create one variable that stores an int
var one int

//create two more variables that store ints
var two, three int

func main() {
	//set the variables
	one = 1
	two = 2
	three = 3

	//print everything
	fmt.Println(one,two,three)
}

//Instructions for running:
//(assuming you have Go installed)
//0.save this file as integers.go
//1.open terminal or cmd
//2.cd to the directory that contains this file
//3.type 'go run integers.go'