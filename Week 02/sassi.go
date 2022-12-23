/*
	Author: Simone Alessandro Casciaro
	Date: 22/12/2022
	Lesson: 2.3
	Text: Dei sassi sferici sono ammucchiati a formare una piramide, con un
		sasso in cima, posto al centro di un quadrato formato da 4 sassi (2
		per lato), posti a loro volta sopra un quadrato formato da 9 sassi (3
		per lato), e così via.
		Scrivete una funzione ricorsiva sassi(height int) int che, data l’altezza height della piramide, restituisca il numero di sassi che la compongono.
		Ad esempio: f(1) deve restituire 1 e f(2) deve restituire 5. */

package main

import "fmt"

func main() {
	fmt.Println(sassi(3))
}

// recursive function
func sassi(heigh int) int {
	if heigh == 1 {
		return 1
	}
	return heigh*heigh + sassi(heigh-1)
}
