package main

import "fmt"

func main() {
	// Aritmatika Operator
	var value = ((2 + 2) * 4)

	/*---------------------------------------------------------------------------------------------------------------------------
	Comparison Operator
	== | apakah nilai kiri sama dengan nilai kanan
	!= | apakah nilai kiri tidak sama dengan nilai kanan
	<  | apakah nilai kiri lebih kecil daripada nilai kanan
	>  | apakah nilai kiri lebih besar daripada nilai kanan
	>= | apakah nilai kiri lebih besar atau sama dengan nilai kanan
	--------------------------------------------------------------------------------------------------------------------------- */
	var isEqual bool = (value == 16)

	/*---------------------------------------------------------------------------------------------------------------------------
	Logic Operator
	&& | kiri dan kanan
	|| | kiri atau kanan
	!  | negasi / nilai kebalikan
	--------------------------------------------------------------------------------------------------------------------------- */
	var left = false
	var right = true

	var leftAndRight = left && right
	var leftOrRight = left || right
	var leftReverse = !left

	fmt.Println(value, isEqual, leftAndRight, leftOrRight, leftReverse)
}
