/*
	Author: Simone Alessandro Casciaro
	Date: 02/01/2023
	Lesson: 6.1
	Text: Scrivete un programma che:
		1. generi a caso una sequenza di interi (di lunghezza massima fissata con una opportuna
			macro) e la memorizzi in una slice;
		2. costruisca un albero binario a partire dalla slice (come descritto in seguito);
		3. stampi l’albero nella reppresentazione a sommario (come descritta in seguito);
		4. stampi i nodi dell’albero negli ordini determinati rispettivamente dalle visite in preordine,
			inordine e postordine*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type bitreeNode struct {
	left  *bitreeNode
	right *bitreeNode
	val   int
}
type bitree struct {
	root *bitreeNode
}

func main() {
	var a = generateNumbers(5)
	fmt.Println(a)
	stampaAlberoASommario(arr2tree(a, 0), 0)
	stampaAlbero(arr2tree(a, 0))
}

// Returns a slice of random numbers as defined in request 1.
func generateNumbers(many int) []int {
	output := make([]int, many)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < many; i++ {
		output[i] = rand.Intn(100)
	}
	return output
}

// Returns a tree starting from a slice as defined in request 2.
func arr2tree(a []int, i int) (root *bitreeNode) {
	if i >= len(a) {
		return nil
	}
	root = &bitreeNode{nil, nil, a[i]}
	root.left = arr2tree(a, 2*i+1)
	root.right = arr2tree(a, 2*i+2)
	return root
}

// Prints the tree "a sommario" as defined in request 3.
func stampaAlberoASommario(node *bitreeNode, spaces int) {
	for i := 0; i < spaces; i++ {
		fmt.Print(" ")
	}
	fmt.Print("*")
	fmt.Println(node.val)
	if node.left != nil || node.right != nil {
		stampaAlberoASommario(node.left, spaces+1)
		stampaAlberoASommario(node.right, spaces+1)
	}
}
func stampaAlbero(node *bitreeNode) {
	if node == nil {
		return
	}
	fmt.Print(node.val, "")
	if node.left == nil && node.right == nil {
		return
	}
	fmt.Print(" [")
	if node.left != nil {
		stampaAlbero(node.left)
	} else {
		fmt.Print("-")
	}
	fmt.Print(", ")
	if node.right != nil {
		stampaAlbero(node.right)
	} else {
		fmt.Print("-")
	}
	fmt.Print("]")
}
