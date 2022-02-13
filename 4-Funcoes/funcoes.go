package main

import "fmt"

/*-----------FUNÇÃO COM 2 PARAMETROS E 1 RETORTNO-----------------------------*/
func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}
/*----------ESTA SENDO CHAMADA DA LINHA 23 A 29----------------------------------*/

/*-----------------------------------------------------------------------------------------------------*/

	/*------------AS FUNÇÕES PODEM TER MAIS DE 1 RETORNO--------------------------------*/

	func calculosMatematicos(n1, n2 int8) (int8, int8) {
		soma := n1 + n2
		subtracao := n1 - n2
		return soma, subtracao
	}
	/*---------------NA LINHA 49 VAMOS CHAMAR A FUNÇÃO calculosMatematicos-----------------*/

func main() {
	/*-----------1° forma-------------------*/
	soma := somar(10, 20)
	fmt.Println(soma)  // retornará o valor 30
  
	/*-----------2° forma-------------------*/
  soma1 := somar(25, 30)
	fmt.Println(soma1)

/*---------1° forma------------------------*/
	var f = func() {
		fmt.Println("Função f")
	}
	f()
/*----------2° forma----------------------*/
	var f1 = func(txt string) {
		fmt.Println(txt)
	}
	f1("Texto da função 1")
/*------------3° forma---------------------*/
	var f2 = func(txt string) string{
		fmt.Println(txt)
		return txt
	}
	resultado := f2("Texto da função 2")
	fmt.Println(resultado)

 resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
 fmt.Println(resultadoSoma, resultadoSubtracao)

 /*-CASO EU QUEIRA CHAMAR A FUNÇÃO QUE TEM 2 RETORNOS ,MAIS EU SO QUERO O PRIMEIRO.UTILIZA ANDERLINE NO SEGUNDO-*/

 resultadoSoma1, _ := calculosMatematicos(10, 15)
 fmt.Println(resultadoSoma1)

 _ , resultadoSubtracao1 := calculosMatematicos(10, 15)
 fmt.Println(resultadoSubtracao1)




}