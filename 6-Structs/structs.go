package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
	endereco endereco
}
//PODEMOS TER UM STRUCT DENTRO DE OUTRO STRUCT.OBRSERVA O STRUCT CRIADO ABAIXO E ACRESCENTADO ACIMA 

type endereco struct {
  logradouro string
	numero    uint8
}

func main() {
	fmt.Println("Arquivo structs")
 //primeira forma de popular uma variavel struct

	var u usuario //criou uma variável do tipo usuário
	u.nome = "Davi" //esta populando a variavel com dados
	u.idade = 21
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua dos Lobos", 49}
/*---------------------------------------------------*/
	//segunda forma de popular um tipo de struct

	usuario2 := usuario{"Davi", 21, enderecoExemplo} //outra forma de popular uma variavel struct com a inferencia de tipos
	fmt.Println(usuario2)
/*---------------------------------------------------*/
  //terceira forma de popular uma variavel struct

	usuario3 := usuario{nome: "Davi"}
	fmt.Println(usuario3)

}