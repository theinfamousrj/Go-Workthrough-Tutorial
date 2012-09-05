package main

import ("fmt"
		"strings")
		
var str1, str2, str3, str4 string
var str123 []string
var total int

func main() {
	
	//lets give these strings some data
	str1 = "This is the first string."
	str2 = "This is the second string."
	str3 = "This is the third string."
	
	//we want to change that first string
	//replacing the first instance of the 
	//word 'is' with the word 'was'
	str1 = strings.Replace(str1, " is ", " was ", 1)
	
	//lets put these strings in a string array
	str123 := []string	{str1,str2,str3}
	
	//we need to use the string array in the Join() function
	str4 = strings.Join(str123, " ")
	
	//now we can count it all up and check it
	fmt.Println(str1,"\t char count:",strings.Count(str1,"")-1)
	fmt.Println(str2,"\t char count:",strings.Count(str2,"")-1)
	fmt.Println(str3,"\t char count:",strings.Count(str3,"")-1)
	fmt.Println(str4,"char count:",strings.Count(str4,"")-1)
	
	//we must account for the two spaces we've
	//added by using the Join() function [line 20]
	total = strings.Count(str1,"")+strings.Count(str2,"")+strings.Count(str3,"")-1
	fmt.Println("Total characters:",total)
	
	//lets print the second string 3 times
	fmt.Println(strings.Repeat(str2,3))
	
}

//Instructions for running:
//(assuming you have Go installed)
//0.save this file as strings.go
//1.open terminal or cmd
//2.cd to the directory that contains this file
//3.type 'go run strings.go'