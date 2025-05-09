package main

import (
	"fmt"
	"math"
)

func main() {
	numeros := make([]int, 15)
	raizes := make([]float64, 15)

	fmt.Println("Digite 15 números inteiros:")

	for i := 0; i < 15; i++ {
		fmt.Printf("Número %d: ", i+1)
		fmt.Scan(&numeros[i])

		if numeros[i] < 0 {
			raizes[i] = -1
		} else {
			raizes[i] = math.Sqrt(float64(numeros[i]))
		}
	}

	fmt.Println("\nVetor de raízes:")
	fmt.Print("(")
	for i := 0; i < 15; i++ {
		fmt.Printf("%.2f", raizes[i])
		if i < 14 {
			fmt.Print(", ")
		}
	}
	fmt.Println(")")
}
