package main

import "fmt"


func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
   fmt.Println("Funçoes Recursivas")//FUNÇÃO RECURSDIVA  TEM A CARACTERISTICA DE CHAMAR ELA MESMA
   

    posicao := uint(5)

		for i := uint(1); i <= posicao; i++ {
			fmt.Println(fibonacci(i))
		}
		fmt.Println("-------------------------")
		fmt.Println(fibonacci(posicao))

	   
}