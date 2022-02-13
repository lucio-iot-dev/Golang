package main

import "fmt"

func main() {
	//ARITMÉTICOS
  soma := 1 + 2
	subtracao := 1-2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)
 // PARA FAZER QUALQUER OPERAÇÃO EM GO AS VARIÁVEIS DEVERÃO SER DO MESMO TIPO,MESMO SE FOR AMBAS DO TIPO INT ,DEVERÃO TER O MESMO VALOR BINÁRIO EX: INT16 COM INT16------INT32 COM INT32
  var numero1 int16 = 10
	var numero2 int16 = 25
	soma2 := numero1 + numero2
	fmt.Println(soma2)
	//FIM DOA ARITMÉTICOS

	//ATRIBUIÇÃO
	var variavel1 string = "string"
	variavel2 := "string2"
	fmt.Println(variavel1, variavel2)
	//FIM DOS OPERADORES DE ATRIBUIÇÃOI
  
  //OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)//DIFERENTE
	//FIM DOS RELACIONAIS

	//OPERADORES LÓGICOS
	fmt.Println("----------------------")
	verdaderio, falso := true, false
	fmt.Println(verdaderio && falso) //E
	fmt.Println(verdaderio || falso) //OU
	fmt.Println(!verdaderio) //negação
	fmt.Println(!falso)
	//FIM DOS OPERADORES LÓGICOS

	//OPERADORES UNÁRIOS
	numero := 10 //retorna 10
	// no go so aceita o pós fixado numero++
	// em algumas linguagens aceita o pre fixado ++numero
	numero++ //numero = numero + 1  retorna 11
	numero+= 15 //numero = numero + 15 retorna 26
	//decrementar
	numero-- //tira 1 no numero  retorna 25
	numero-= 15 //tira 15 no numero  retorna 10
	numero*= 3 //numero = numero * 3 retorna 30
	numero/= 3 //nunero = numero / 3  retorna 10
	numero%= 3 // numero = numero % 3 retorna 1
	//----------------------------------//
	numero = 10 // retorna 10
	numero%= 2 //numero = numero % 2  retorna 0
	fmt.Println(numero)
	//FIM DOS OPERADORES UNÁRIOS

	// OBSERVAÇÕES
	// OPERADORES TERNÁRIOS
	// texto := numero > 5 ? "Maior que 5" : "Menor que 5"--->> le-se:  se numero menor que 5 então escreve: Maior que 5 senão escreve: Menor que 5
	// nas outras linguagens essa sintaxe é aceita no go não aceita esse tipo de condição
	// COMO RESOLVER NO GO A RESOLUÇÃO DO PROBLEMA ACIMA?
	// UTILIZANDO IF /ELSE
	// numero = 6 //para fazer teste de mesa

	var texto string
	if numero >5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}
	fmt.Println(texto)







  
 









}