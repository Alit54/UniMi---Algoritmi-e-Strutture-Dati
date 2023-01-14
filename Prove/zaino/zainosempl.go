package main

type item struct {
	peso, valore int
}

func main() {

}

func a(maxPeso int, items ...item) {
	slice := make([]int, maxPeso+1)
	slice[0] = 0
	max := 0
	// Programmazione Dinamica: uso soluzione n per creare n+1
	for i := 1; i <= maxPeso; i++ {
		// Cerco il max usando PD
		for t := 0; t < len(items); t++ {
			if items[t].peso > max {
				max = items[t].peso
			}
		}
	}
}
