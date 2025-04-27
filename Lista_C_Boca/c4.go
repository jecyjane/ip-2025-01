package main

import "fmt"

func main() {

	var num int

	for {
		fmt.Scan(&num)
		if num <= 0 {
			break
		}

		ehQuadradoPerfeito := false
		for i := 0; i*i <= num; i++ {
			if i*i == num {
				ehQuadradoPerfeito = true
				break
			}
		}

		if ehQuadradoPerfeito {
			fmt.Printf("%d eh quadrado perfeito\n", num)
		} else {
			fmt.Printf("%d nao eh quadrado perfeito\n", num)
		}
	}
}
