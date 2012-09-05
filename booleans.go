package main

import "fmt"

//instantiate the variables
var one, two int
var b1, b2, b3 bool

func main() {
  
	//set the variables
	one = 1
	two = 2
	b1 = true
	b2 = false
	b3 = true

	//print the variables
	fmt.Println(one, two, b1, b2, b3)
	
}

//Instructions for running:
//(assuming you have Go installed)
//0.save this file as booleans.go
//1.open terminal or cmd
//2.cd to the directory that contains this file
//3.type 'go run booleans.go'