package main
import "fmt"
func main() {
	var num[6]int

	num[0] = 1
	num[1] = 2
	num[2] = 4
	num[3] = 5
	num[4] = 0
	num[5] = 8

	
	fmt.Println(num[0])
	fmt.Println(num[1])
	fmt.Println(num[2])
	fmt.Println(num[3])
	fmt.Println(num[4])
	fmt.Println(num[5])


	
	for i := 0; i<len(num); i++ {
		fmt.Println(num[i])
	}


	
	for _ , v:= range num{
		fmt.Println(v)
	}

	
	var cadastro [2][3]string
	for i := 0; i < 2; i++{
		for j := 0; j < 3; j++{
			fmt.Println("Digite o nome: ",i,j)
			fmt.Scan(&cadastro[i][j])
		}
	
	}
	fmt.Println(cadastro)

	
	
}


package main
import "fmt"
type Produtos struct{
		nome string
		valor float64
		quantidade int
}
func main(){
	p := Produtos{
		nome: "Arroz",
		valor: 5.99,
		quantidade: 10,
	}
	
	j:= Produtos{
		nome: "Feijão",
		valor: 4.99,
		quantidade: 20,
	}
	fmt.Println(p)
	fmt.Println(j)
}