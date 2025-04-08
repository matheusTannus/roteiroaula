package main

import "fmt"

func soma(n int) int{
	var soma int
	for i:=1; i<=n; i++{
		soma += i
	}
	return soma
}

func main() {
	var num int
	fmt.Println("insira um numero ?")
	fmt.Scanln(&num)
	fmt.Println(soma(num))
	
}