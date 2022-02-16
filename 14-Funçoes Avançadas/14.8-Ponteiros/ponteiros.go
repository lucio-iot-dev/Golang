package main

import "fmt"

func inverterSinal(numero int) int {
	return numero *-1
}

func inverterSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
   numero := 20
	 numeroInvertido := inverterSinal(numero)
	 fmt.Println(numeroInvertido)
	// fmt.Println(inverterSinal(numero))//outra forma de chamar
	//A vaiável numero não é alterada, simplesmente é passada uma cópia da variável numero para a função de cima executar .Uma prova disso vamos imprimir a variável numero e verá que estará com o sinal positivo novamente
	fmt.Println(numero)

	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)

	
}