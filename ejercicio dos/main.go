package main

import "fmt"

func main() {
	votos := map[string]int{"deportes": 0, "videojuegos": 0, "cine": 0, "musica": 0}

	fmt.Println("Ingresa 5 votos (deportes, videojuegos, cine, musica):")
	for i := 0; i < 5; i++ {
		var act string
		fmt.Scan(&act)
		
		if _, ok := votos[act]; ok {
			votos[act]++
		} else {
			fmt.Println("Opción no válida. Intenta de nuevo:")
			i--
		}
	}

	fmt.Println("\nResultados:")
	for act, cant := range votos {
		fmt.Printf("- %s: %d\n", act, cant)
	}

	fmt.Printf("🏆 Actividad ganadora: %s\n", obtenerGanadora(votos))
}

func obtenerGanadora(m map[string]int) string {
	ganadora, max := "", -1
	for act, cant := range m {
		if cant > max {
			max, ganadora = cant, act
		}
	}
	return ganadora
}