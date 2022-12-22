/*
	Author: Simone Alessandro Casciaro
	Date: 22/12/2022
	Lesson: 1.11
	Text: Si vuole scrivere una funzione che, presa come argomento una slice di stringhe, restituisca il numero massimo di consonanti contenute in una stringa (assumiamo per semplicità che le stringhe
		contengano solo caratteri minuscoli). */

package main

import "fmt"

func main() {
	slice := []string{"miky", "simo"}
	fmt.Println(countConsonants(slice))
}

// Returns the max number of consonants from the strings
func countConsonants(s []string) int {
	max, counter := 0, 0
	for i := 0; i < len(s); i++ {
		counter = 0
		for _, v := range s[i] {
			if v != 'a' && v != 'e' && v != 'i' && v != 'o' && v != 'u' {
				counter++
			}
		}
		if counter > max {
			max = counter
		}
	}
	return max
}
