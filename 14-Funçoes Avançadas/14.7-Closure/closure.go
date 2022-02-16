package main

import "fmt"

//CLOSURE SÃO FUNÇÕES QUE REFERENCIAM VARIÁVEIS QUE ESTÃO FORA DO SEU CORPO

func closure() func () {
	texto := "Dentro da função closure"

	funcao := func() {
		fmt.Println(texto)
	}
	return funcao
}


func main() {
	texto := "Dentro da função main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()
}