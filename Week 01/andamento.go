/*
	Author: Simone Alessandro Casciaro
	Date: 22/12/2022
	Lesson: 1.3
	Text: Data da standard input una serie di interi positivi terminata da 0, stampare ’+’ ogni volta che il
		nuovo valore è maggiore o uguale al precedente e ’-’ altrimenti. */

package main

import "fmt"

func main() {
	var prec int
	number := -1
	// Read first number
	fmt.Scan(&prec)
	// Continue reading until number is != 0 and print "+" or "-"
	for number != 0 {
		fmt.Scan(&number)
		if number < prec {
			fmt.Println("-")
		} else {
			fmt.Println("+")
		}
		prec = number
	}
}
