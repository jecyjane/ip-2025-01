package main

import "fmt"

func main() {
	var l, r int
	fmt.Scan(&l, &r)

	soma := 0
	contador := 0

	for i := l; i <= r; i++ {
		if i%2 == 0 {
			soma += i
			contador++
		}
	}

	fmt.Println(soma)

	if contador > 0 {
		media := float64(soma) / float64(contador)
		fmt.Println(media)
	} else {
		fmt.Println(0)
	}
}
