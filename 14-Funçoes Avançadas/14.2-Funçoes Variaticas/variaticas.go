package main

import "fmt"


//COLOCANDO ...INT QUER DIZER QUE ELA VAI RECEBER QUANTOS NUMEROS EU QUISER
func soma(numeros ...int) int {
   total := 0
	 for _, numero := range numeros {
		 total += numero
		}

	 return total
}
/*----------------------------------------------------*/
//PODEMOS PASSAR STRING COM NUMEROS DA SEGUINTE FORMA:
func escrever(texto string, numeros ...int) { //O VARIATICO SEMPRE TEM QUE FICAR NO ULTIMO PARAMETRO
	for _, numero := range numeros {
      fmt.Println(texto, numero)
		}
	}


func main() {
   totalDaSoma := soma(1, 2, 3, 4, 5, 6, 200, 102, 12, 13)
	 fmt.Println(totalDaSoma)

	 escrever("Olá Mundo", 10, 2, 3, 4, 5, 6, 7)
}