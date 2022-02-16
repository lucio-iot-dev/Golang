package main

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("execução foi recuperada com sucesso!")
	}
}


func alunosEstaAprovado(n1, n2 float64) bool {
	 defer recuperarExecucao()
   media := (n1 + n2) / 2

	 if media > 6 {
		 return true
	 } else if media < 6 {
		   return false
	 }
   panic("A media é exatamente 6 !")

	}

func main() {
   fmt.Println(alunosEstaAprovado(6, 6))
   fmt.Println("Pós execução!")
   

}