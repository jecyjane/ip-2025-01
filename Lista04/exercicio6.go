package main

import "fmt"

func main() {
	vetor := make([]int, 100)

	for i := 0; i < 100; i++ {
		vetor[i] = 100 - i
	}

	fmt.Print("Elementos do vetor: (")
	for i, num := range vetor {
		fmt.Printf("%d", num)
		if i < len(vetor)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println(")")
}
