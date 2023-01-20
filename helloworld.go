package main

import "fmt"

/*---------------------------------------------------------------------------------------------------------------------------
For me, just keep learning!
--------------------------------------------------------------------------------------------------------------------------- */

func main() {
	// var with data type
	var name string = "bayu"
	// no var no data type
	job := "frontend"
	// can't use := for reasign variable :=, use = instead
	job = "backend"
	// multiple declaration
	var (
		firstname = "bayu"
		lastname = "setiawan"
	)
	// like console.log()
	fmt.Println("hello world")
	// fmt.Printf() not created new line, so use \n instead for spacing
	fmt.Printf("satu %s %s!\n", firstname, lastname)
	fmt.Println(name)
	fmt.Println(job)
	fmt.Println(firstname, lastname)
}