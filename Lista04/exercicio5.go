package main

import "fmt"

func main() {
	var vetor [10]int
	fmt.Println("Digite 10 números inteiros diferentes para o vetor:")
	for i := 0; i < 10; i++ {
		fmt.Printf("Elemento %d: ", i+1)
		fmt.Scan(&vetor[i])
	}

	menor := vetor[0]
	posicao := 0

	for i := 1; i < 10; i++ {
		if vetor[i] < menor {
			menor = vetor[i]
			posicao = i
		}
	}

	fmt.Printf("O menor elemento do vetor é: %d e sua posição dentro do vetor é: %d.\n", menor, posicao)
}
