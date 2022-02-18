package main

import (
	"fmt"
	"time"
)

func main() {
   canal := make(chan string)
	 go escrever("Olá Mundo!", canal)

	 fmt.Println("Depois da função escrever começar a ser executada")

	 //Primeira forma
	//  for {
	// 	 mensagem, aberto := <- canal //Esperar o canal receber um valor e guardar esse valor na variável mensagem
	// 	 if !aberto{
	// 		 break
	// 	 }
	// 	 fmt.Println(mensagem)
	//  }
   //-------------------------------------

	 //Segunda forma
	 for  mensagem := range canal {
		fmt.Println(mensagem)
	 }


	 fmt.Println("Fim do programa!")

}



func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		 canal <- texto
		 time.Sleep(time.Second)
	}
  close(canal)
}