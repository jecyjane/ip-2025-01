package main

import "fmt"

func main() {
	var vetor1 [10]int
	fmt.Println("Digite 10 números inteiros para o primeiro vetor:")
	for i := 0; i < 10; i++ {
		fmt.Printf("Elemento %d: ", i+1)
		fmt.Scan(&vetor1[i])
	}

	var vetor2 [5]int
	fmt.Println("\nDigite 5 números inteiros para o segundo vetor:")
	for i := 0; i < 5; i++ {
		fmt.Printf("Elemento %d: ", i+1)
		fmt.Scan(&vetor2[i])
	}

	somaVetor2 := 0
	for _, num := range vetor2 {
		somaVetor2 += num
	}

	var vetorResultantePar [10]int
	var vetorResultanteImpar [10]int

	for i, num := range vetor1 {
		if num%2 == 0 {
			vetorResultantePar[i] = num + somaVetor2
		} else {
			vetorResultanteImpar[i] = num + somaVetor2
		}
	}

	fmt.Println("\nPrimeiro vetor:", vetor1)
	fmt.Println("Segundo vetor:", vetor2)

	fmt.Print("Primeiro vetor resultante: [")
	for i, val := range vetorResultantePar {
		fmt.Print(val)
		if i < len(vetorResultantePar)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println("]")

	fmt.Print("Segundo vetor resultante: [")
	for i, val := range vetorResultanteImpar {
		fmt.Print(val)
		if i < len(vetorResultanteImpar)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println("]")
}
