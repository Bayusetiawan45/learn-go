package main

import "fmt"

func main() {
	/*---------------------------------------------------------------------------------------------------------------------------
	If Else kayak biasa, tapi ga perlu kurung ()
	--------------------------------------------------------------------------------------------------------------------------- */
	var point = 10
	if point == 10 {
		fmt.Println("lulus dengan sempurna")
	} else if point >= 8 {
		fmt.Println("lulus")
	} else if point < 8 {
		fmt.Println("tidak lulus")
	}

	/*---------------------------------------------------------------------------------------------------------------------------
	Temporary Variable adalah variabel yang hanya bisa digunakan pada blok seleksi kondisi di mana ia ditempatkan saja.
	Ex: percent
	--------------------------------------------------------------------------------------------------------------------------- */
	var score = 8654
	if percent := score / 100; percent >= 100 {
		fmt.Println("Perfect win", percent, "%")
	} else if percent >= 70 {
		fmt.Println("Good", percent, "%")
	} else {
		fmt.Println("Not bad", percent, "%")
	}

	/*---------------------------------------------------------------------------------------------------------------------------
	Switch Case juga mirip JS
	--------------------------------------------------------------------------------------------------------------------------- */
	var gender = "MALE"
	switch gender {
	case "MALE":
		fmt.Println("Laki-laki")
	case "FEMALE":
		fmt.Println("Perempuan")
		// bisa dikasih default juga setelahnya
	}
}
