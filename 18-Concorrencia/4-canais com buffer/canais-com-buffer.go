package main

import "fmt"


func main() {
	 canal := make(chan string, 2) //Esse numero 2 é o buffer,determina o numero de envios e recebimentos para não dar deadlock!
	 canal <- "Olá Mundo!"
	 canal <- "Programando em Go"

	 mensagem := <-canal
	 mensagem2 := <-canal

	 fmt.Println(mensagem)
	 fmt.Println(mensagem2)

}