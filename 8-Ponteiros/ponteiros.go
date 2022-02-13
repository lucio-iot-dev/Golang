package main

import "fmt"


func main() {
	var variavel1 int = 10
	var variavel2 int = variavel1
	fmt.Println(variavel1, variavel2) //10  10

  variavel1++
	fmt.Println(variavel1, variavel2) //11  10
	//ISTO ACONTECE QUANDO SE PASSA VALORES POR CÓPIA
/*--------------------------------------------------------------*/

//DIFERENÇA ATRIBUINDO VALORES USANDO PONTEIROS
// PONTEIRO É UMA REFERENCIA DE MEMÓRIA

  var variavel3 int //GUARDA UM VALOR INTEIRO NA VARIAVEL
	var ponteiro *int //GUARDA UM ENDEREÇO DE MEMORIA DE UM VALOR INTEIRO

	fmt.Println(variavel3, ponteiro) // retorna 0 e <nil> (quando não se atribui nada a variavel)

	variavel3 = 100
	ponteiro = &variavel3 //COLOCA-SE O E COMERCIAL (&) E O NOME DA VARIAVEL ONDE SE QUER GUARDAR O ENDEREÇO

	fmt.Println(variavel3, ponteiro)

	//PARA FAZER A DESREFERENCIAÇÃO E MOSTRAR O VALOR DA VARIAVEL QUE ESTA SALVA NO ENDEREÇO DE MEMÓRIA E SÓ COLOCAR UM * ANTES DE PONTEIRO EX:
	fmt.Println(variavel3, *ponteiro) //desreferenciação

	





 

  
  

}