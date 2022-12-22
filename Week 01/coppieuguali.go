/*
	Author: Simone Alessandro Casciaro
	Date: 22/12/2022
	Lesson: 1.10
	Text: Scrivere un programma che legga da standard input una sequenza di numeri terminata da zero e
		conti quante sono le coppie di numeri naturali consecutivi uguali.
*/

package main

import "fmt"

func main() {
	counter := 0
	var prec int
	number := -1
	fmt.Scan(&prec)
	// Read input and in the meantime count numbers
	for number != 0 {
		fmt.Scan(&number)
		if number == prec {
			counter++
			continue
		}
		prec = number
	}
	fmt.Println(counter)
}
