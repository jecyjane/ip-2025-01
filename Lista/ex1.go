package main
import (
 "fmt"
)
func main() {
 notas := make([]float64, 3)
 for i := 0; i < 3; i++ {
 fmt.Printf("Digite a nota %d: ", i+1)
 fmt.Scanln(&notas[i])
 }
 media := (notas[0] + notas[1] + notas[2]) / 3
 fmt.Printf("MEDIA = %.2f\n", media)
 if media >= 6 {
 fmt.Println("APROVADO")
 } else {
 fmt.Println("REPROVADO")
 }
}
