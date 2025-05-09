package main

import "fmt"

func main() {
	vetorImpares := make([]int, 100)

	// Armazena os 100 primeiros números ímpares no vetor
	numeroImpar := 1
	for i := 0; i < 100; i++ {
		vetorImpares[i] = numeroImpar
		numeroImpar += 2
	}

	// Imprime o vetor entre colchetes
	fmt.Print("Vetor de números ímpares: [")
	for i, num := range vetorImpares {
		fmt.Printf("%d", num)
		if i < len(vetorImpares)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println("]")
}
