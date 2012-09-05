package main

//import multiple items
import ("fmt"
		"time"
		"strings")

//instantiate string variable
var str1 string
var str1Count, timeCount int

func main() {

	//set the variable
	str1 = "The time is:"

	//print the string and time now
	fmt.Println(str1,time.Now().Format(time.UnixDate))

	//count the characters in the strings
	str1Count = strings.Count(str1,"")
	timeCount = strings.Count(time.Now().Format(time.UnixDate),"")

	//print the character count
	//remember, we have to subtract
	//one from the total count
	fmt.Println("There are",str1Count+timeCount-1,"characters in the previous line.")
	
}

//Instructions for running:
//(assuming you have Go installed)
//0.save this file as imports.go
//1.open terminal or cmd
//2.cd to the directory that contains this file
//3.type 'go run imports.go'