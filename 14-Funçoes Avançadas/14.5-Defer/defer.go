package main

import "fmt"
   


func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Media calculada. Resultado será retornado!")
	fmt.Println("Entrando na função para verificar se o aluno está aprovado")

	media := (n1 + n2) / 2
  
	if media >= 6 {
		return true
	}
	return false
}

func main() {
  
	defer funcao1() //DESTA FORMA ELE EXECUTA A FUNÇÃO DE BAIXO PRIMEIRO E DEPOIS VOLTA E EXECUTA A FUNÇÃO DEFER  DEFER = Adiar 
	funcao2()

	fmt.Println(alunoEstaAprovado(7, 8))

}