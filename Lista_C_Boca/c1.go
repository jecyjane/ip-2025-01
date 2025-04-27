package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	soma := 0
	for i := 1; i <= n; i++ {
		fmt.Print(i)
		if i < n {
			fmt.Print(" ")
		}
		soma += i
	}
	fmt.Println()
	fmt.Println(soma)
}
