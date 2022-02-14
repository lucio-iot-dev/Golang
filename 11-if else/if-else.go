package main

import "fmt"

func main() {
	fmt.Println("Estrutura de controle")

	numero := 10

	if numero >15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}


	//CRIAR UMA VARIAVEL DENTRO DE UM IF
	// IMPORTANTE!! SE CRIAR UMA VARIAVEL DENTRO DO IF/ELSE ELA FICA LIMITADA A FUNCIONAR SOMENTE DENTRO DESTE LAÇO.
	//SE FOR UTILIZAR ESSA VARIAVEL FORA DO LAÇO ELA APARECERA COM NÃO DEFINIDA
	if outroNumero := numero; outroNumero >0 {
		fmt.Println("Numero é maior que zero")
	}

}