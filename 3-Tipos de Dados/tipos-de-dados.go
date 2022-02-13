package main

import (
	"errors"
	"fmt"
)

func main() {
	  //int8, int16, int32, int64  o numero refere-se ao numero de bits
		var numero int16 = 100
		fmt.Println(numero)

		var numero2 int64 = -100000000 //aceita numeros negativos
		fmt.Println(numero2)

		var numero3 uint32 = 10000 //não aceita numeros negativos
		fmt.Println(numero3)

		// alias (significa apelido)
		var numero4 rune = 1234 // rune é o mesmo que colocar int32 -->> rune = unit32
		fmt.Println(numero4)

		var numero5 byte = 123 // byte é o mesmo que digitar uint8 -->> byte = unit8
		fmt.Println(numero5)

		// numeros quebrados so existe : float32 e flat64

		var numeroReal1 float32 = 123.45
		fmt.Println(numeroReal1)

		var numeroReal3 float64 = 1230000000.45
		fmt.Println(numeroReal3)

		numeroReal4 := 12345.67  //inferencia de tipo
		fmt.Println(numeroReal4)

		var str string = "Texto"
		fmt.Println(str)

		str2 := "Texto2"
		fmt.Println(str2)

		char := 'A'
		fmt.Println(char) //a resposta vai ser um numero 65 referente a tabela asc

		var texto string
		fmt.Println(texto) //devolve um resultado em branco

		var texto2 int
		fmt.Println(texto2) //devolve zero

		var booleano1 bool = true
		fmt.Println(booleano1)
		
		var booleano2 bool 
		fmt.Println(booleano2) // se não declarar o valor vem como padão o zero que é representado por false

		var erro error
		fmt.Println(erro) // retorna o valor <nil> significa zero

		var erro2 error = errors.New("Erro interno") // Gerar uma mensagem de erro interno na aplicação. 
		fmt.Println(erro2) // erro2 : variável
		                   // error : Tipo da variável
											 // errors : Nome do pacote (pacote existente dentro do go .Detalhe: Tem que ser declarado la em cima)
		                    









}