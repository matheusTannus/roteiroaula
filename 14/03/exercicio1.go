package main

import "fmt"

func ParOiImpar(num int) string{
	if num%2==0{
		return "par"
	}else{
		return "impar"
	}
}

func main() {
	var n int
	fmt.Println("insira um numero ?")
	fmt.Scanln(&n)
	fmt.Println(ParOiImpar(n))
	
}