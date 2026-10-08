package main

import "fmt"

func main() {

	notas := [6][4]float64{
		{8, 9, 7, 8}, {6, 7, 6, 7}, {9, 10, 9, 9},
		{5, 6, 5, 6}, {8, 8, 9, 8}, {7, 7, 8, 7},
	}

	var sumaTotal float64

	for i, est := range notas {
		s := est[:] // Convertir array a slice
		suma, max, min := 0.0, s[0], s[0]

		for _, n := range s {
			suma += n
			if n > max { max = n }
			if n < min { min = n }
		}

		sumaTotal += suma
		fmt.Printf("Est. %d -> Prom: %.1f | Max: %.1f | Min: %.1f\n", i+1, suma/4, max, min)
	}

	fmt.Printf("Promedio general de la clase: %.1f\n", sumaTotal/24)
}