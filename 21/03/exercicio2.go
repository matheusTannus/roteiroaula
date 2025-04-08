//exercicio 1 e 1.1
package main
import "fmt"
func main(){
	var b float64 = 50.0
	var p2 *float64 = &b
	fmt.Println("Valor de b antes da modificação:", b)
	*p2 = 100.0
	fmt.Println("Valor de b após a modificação:", b)
}



package main
import "fmt"
func incrementByPointer(val *int) {
*val++
fmt.Println("Dentro da função incrementByPointer:", *val)
}
func main() {
x := 10
incrementByPointer(&x)
fmt.Println("Fora da função incrementByPointer:", x)
}