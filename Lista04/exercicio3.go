package main

import "fmt"

func main() {
	var numeros [10]int
	fmt.Println("Digite 10 números inteiros:")
	for i := 0; i < 10; i++ {
		fmt.Printf("Número %d: ", i+1)
		fmt.Scan(&numeros[i])
	}

	pares := []int{}
	somaPares := 0
	impares := []int{}
	quantidadeImpares := 0

	for _, num := range numeros {
		if num%2 == 0 {
			pares = append(pares, num)
			somaPares += num
		} else {
			impares = append(impares, num)
			quantidadeImpares++
		}
	}

	fmt.Println("\nNúmeros pares digitados:", pares)
	fmt.Println("Soma dos números pares digitados:", somaPares)
	fmt.Println("Números ímpares digitados:", impares)
	fmt.Println("Quantidade de números ímpares digitados:", quantidadeImpares)
}
