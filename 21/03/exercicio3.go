//exercicio swap
package main
import "fmt"
func troca(a , b *int){
	*a, *b = *b, *a
} 
func main(){
	a := 10
	b := 20
	fmt.Println("Antes da troca: a =", a, "b =", b)
	troca(&a, &b)
	fmt.Println("Depois da troca: a =", a, "b =", b)
}