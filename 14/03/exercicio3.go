package main

import "fmt"

func classificarnota(nota int) string{
	switch {
		case nota >= 9:
			return "Excelente"
		case nota >= 7:
			return "Bom"
		case nota >= 5:
			return "Regular"
		case nota >= 3:
			return "Ruim"
		default:
			return "Muito ruim"
	}
	
}

func main() {
	var num int
	fmt.Println("insira a nota de 0 até 10 ?")
	fmt.Scanln(&num)
	fmt.Println(classificarnota(num))
	
}