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

  //print the strings
  fmt.Println(str1,time.Now().Format(time.UnixDate))

  //count the characters in the strings
  str1Count = strings.Count(str1,"")
  timeCount = strings.Count(time.Now().Format(time.UnixDate),"")

  //print the character count
  fmt.Println("There are",str1Count+timeCount-1,"characters in the previous line.")
}
