package main

import "fmt"

func main() {
	var numeros [10]int

	fmt.Println("Digite 10 números inteiros:")
	for i := 0; i < 10; i++ {
		fmt.Printf("Número %d: ", i+1)
		fmt.Scan(&numeros[i])
	}

	encontrouSuperior := false

	fmt.Println("\nNúmeros superiores a 50 e suas posições:")
	for i, numero := range numeros {
		if numero > 50 {
			fmt.Printf("Número: %d, Posição: %d\n", numero, i)
			encontrouSuperior = true
		}
	}

	if !encontrouSuperior {
		fmt.Println("\nNão existe nenhum número superior a 50 no vetor.")
	}
}
