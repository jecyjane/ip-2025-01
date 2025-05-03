package main

import "fmt"

func main() {
	var a [10]int
	fmt.Println("Digite 10 números inteiros para o vetor A:")
	for i := 0; i < 10; i++ {
		fmt.Printf("Elemento %d: ", i+1)
		fmt.Scan(&a[i])
	}

	ocorrencias := make(map[int]int)
	for _, num := range a {
		ocorrencias[num]++
	}

	fmt.Println("\nElementos repetidos e suas ocorrências:")
	for num, count := range ocorrencias {
		if count > 1 {
			fmt.Printf("O número %d se repete %d vezes.\n", num, count)
		}
	}
}
