package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	Alit "github.com/Alit54/General/gofunc"
)

func main() {
	// Compilare gli archi in base all'input
	graph := createGraph()
	// Effettuare ricerca (BFS?) utilizzando la EQL
	// Restituire lunghezza di EQL
	fmt.Println("Parte A:", len(BFS(graph, "0")))
	fmt.Println("Parte B:", groups(graph))
}

// Crea il grafo leggendo da input i collegamenti
func createGraph() map[string][]string {
	graph := make(map[string][]string)
	scanner := bufio.NewScanner(Alit.File(os.Args[1]))
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), " <-> ")
		graph[row[0]] = strings.Split(row[1], ", ")
	}
	return graph
}

// Effettua ricerca in ampiezza del grafo e restituisce la lunghezza della EQL
func BFS(graph map[string][]string, start string) map[string]bool {
	EQL := make(map[string]bool)
	F := make([]string, 0) // Pila FIFO: BFS
	// Inizializza EQL e F col nodo 0
	EQL[start] = true
	F = append(F, start)
	for len(F) > 0 {
		// Scegli nodo da espandere
		node := F[0]
		F = Alit.RemoveFirstElement(F)
		// Espandi nodo
		for i := 0; i < len(graph[node]); i++ {
			// Aggiungo a F
			if !EQL[graph[node][i]] {
				F = append(F, graph[node][i])
				// Aggiungo a EQL i nodi visitati
				EQL[graph[node][i]] = true
			}
		}
	}
	return EQL
}

// Restituisce quanti gruppi esistono (Quante volte viene chiamata BFS)
func groups(graph map[string][]string) int {
	counter := 0
	// Chiama il primo BFS
	EQL := BFS(graph, "0")
	counter++
	// Ripeto per ogni nodo non ancora visitato
	for i := 0; i < len(graph); i++ {
		tmp := strconv.Itoa(i)
		if !EQL[tmp] {
			tmpEQL := BFS(graph, tmp)
			for k := range tmpEQL {
				EQL[k] = true
			}
			counter++
		}
	}
	return counter
}
