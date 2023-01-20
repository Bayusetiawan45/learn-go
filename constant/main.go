package main

import "fmt"

/*---------------------------------------------------------------------------------------------------------------------------
Konstanta adalah jenis variabel yang nilainya tidak bisa diubah. Inisialisasi nilai hanya dilakukan sekali di awal, setelah itu variabel tidak bisa diubah nilainya
--------------------------------------------------------------------------------------------------------------------------- */

func main() {
	const firstname = "bayu"
	// multiple declaration
	const (
		name         = "bayu setiawan"
		isMale bool  = true
		age    uint8 = 22
	)
	/*---------------------------------------------------------------------------------------------------------------------------
	Ketika tipe data dan nilai tidak dituliskan dalam deklarasi konstanta, maka tipe data dan nilai yang dipergunakan adalah sama seperti konstanta yang dideklarasikan diatasnya
	--------------------------------------------------------------------------------------------------------------------------- */
	const (
		today = "senin"
		now
	)

	fmt.Println(firstname, name, isMale, age, today, now)
}
