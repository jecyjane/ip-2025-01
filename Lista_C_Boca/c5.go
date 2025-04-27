package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	ehTriangular := false
	i := 1
	for {
		produto := i * (i + 1) * (i + 2)
		if produto == num {
			ehTriangular = true
			break
		}
		if produto > num {
			break
		}
		i++
	}

	if ehTriangular {
		fmt.Printf("%d eh triangular\n", num)
	} else {
		fmt.Printf("%d nao eh triangular\n", num)
	}
}
