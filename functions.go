package main

import ("fmt"
		"strings")
		
func whoIsAWhat(who, what string) {
	var is, a string
	
	if strings.ToLower(who) == "i" {
		is = "am"
	} else if strings.ToLower(who) == "you" {
		is = "are"
	} else {
		is = "is"
	}
	
	if strings.Index(what, "a") == 0 || strings.Index(what, "e") == 0 || strings.Index(what, "i") == 0 || strings.Index(what, "o") == 0 {
		a = "an"
	} else {
		a = "a"
	}
	
	fmt.Printf("%v %v %v %v!\n", who, is, a, what)
}

func main() {
	whoIsAWhat("John Pisano","mother lover")
}