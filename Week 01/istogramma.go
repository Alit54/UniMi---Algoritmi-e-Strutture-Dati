/*
	Author: Simone Alessandro Casciaro
	Date: 22/12/2022
	Lesson: 1.6
	Text:
		1. Scrivete un programma che legga una riga di caratteri terminata da a-capo, poi legga
			un’altra lettera e stampi quante volte la lettera compare nella prima riga.
		2. Scrivete un programma che legga una riga di caratteri terminata da a-capo e che visualizzi
			un istogramma con una barra per ogni carattere dell’alfabeto, lunga quanto il numero
			delle sue occorrenze. Il programma non deve visualizzare le barre delle lettere che non
			compaiono e non deve fare distinzione fra maiuscole e minuscole.
		3. Due parole costituiscono un anagramma se l’una si ottiene dall’altra permutando le lettere (es: attore, teatro). Scrivete un programma che legga due parole e verifichi se sono
			anagrammi.
		Suggerimento: potete sfruttare il programma scritto per l’esercizio precedente. */

package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	var letter string
	// Read until you receive '\n'
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	// Read the letter to search
	fmt.Scan(&letter)

	fmt.Println(howMany(scanner.Text(), rune(letter[0])))
	printHistogram(scanner.Text())
	fmt.Println(anagrams(scanner.Text(), "Teatro"))
}

// Returns how many times word contains char as defined in request 1
func howMany(word string, char rune) int {
	counter := 0
	for _, v := range word {
		if unicode.ToUpper(v) == unicode.ToUpper(char) {
			counter++
		}
	}
	return counter
}

// Prints the histogram as defined in request 2
func printHistogram(word string) {
	for i := 'A'; i <= 'Z'; i++ {
		tmp := howMany(word, rune(i))
		if tmp != 0 {
			fmt.Printf("%c ", i)
			for j := 0; j < tmp; j++ {
				fmt.Print("*")
			}
			fmt.Println()
		}
	}
}

// Returns if 2 words are anagrams as defined in request 3
func anagrams(word, word2 string) bool {
	for i := 'A'; i < 'Z'; i++ {
		if howMany(word, rune(i)) != howMany(word2, rune(i)) {
			return false
		}
	}
	return true
}
