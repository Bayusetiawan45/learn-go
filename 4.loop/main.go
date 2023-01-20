package main

import "fmt"

func main() {

	// use for
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	// loop use condition only
	var a = 0
	for a < 5 {
		fmt.Println("Angka a", a)
		a++
	}
}
