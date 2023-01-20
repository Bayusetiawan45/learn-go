package main

import "fmt"

/*---------------------------------------------------------------------------------------------------------------------------
Konstanta adalah jenis variabel yang nilainya tidak bisa diubah. Inisialisasi nilai hanya dilakukan sekali di awal, setelah itu variabel tidak bisa diubah nilainya
--------------------------------------------------------------------------------------------------------------------------- */

func constant() {
	const firstname string = "bayu"

	fmt.Println(firstname)
}