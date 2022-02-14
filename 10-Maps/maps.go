package main

import "fmt"

func main() {
	fmt.Println("Maps")
	//pimeiro string é o tipo das chaves e o segundo valores
	//parece um pouco com a estrutura structs
 usuario := map[string]string{  //so aceita se os dois campos for string ou float ou int ,tem que ser ambos do mesmo tipo
	 "nome":      "Pedro",
	 "sobrenome": "Silva",
 }
 fmt.Println(usuario)
 fmt.Println(usuario["nome"])// não funciona colocando usuario.nome //retorna Pedro

 //ANINHANDO MAPS--->> UM MAP DENTRO DE OUTRO

 usuario2 := map[string]map[string]string {
	 "nome": {
		 "primeiro":"Joaõ",
		 "ultimo": "Pedro",
	 },
	 "curso": {
		 "nome": "Engenharia",
		 "campus": "Campus 1",
	 },
 }
 fmt.Println(usuario2)
 // para apagar uma chave
 delete(usuario2, "nome")
 fmt.Println(usuario2)

 //PARA ADICIONAR UMA CHAVE

 usuario2["Signo"] = map[string]string {
	 "nome": "Gêmeos",
 }

  fmt.Println(usuario2)

}