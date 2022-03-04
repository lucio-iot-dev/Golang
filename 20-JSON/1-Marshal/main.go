package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome string `json:"nome"`
	Raca string `json:"raca"`
	Idade uint `json:"idade"`
}

func main() {
	 c := cachorro{"Rex", "Dalmata", 3}
	 fmt.Println(c)

	 cachorroEmJSON, erro := json.Marshal(c)
	 if erro != nil {
		 log.Fatal(erro)
	 }
	 fmt.Println(cachorroEmJSON) //DEVOLVE UM PACOTE DE BYTES (SLICES DE BYTES)
	 fmt.Println(bytes.NewBuffer(cachorroEmJSON))//VAI CRIAR UM NOVO BUFFER BASEADO EM UM SLICE DE BYTES

	 /*---------------UTILIZANDO MAP-------------------------------------*/

	 c2 := map[string]string{
		 "nome": "Toby",
		 "raca": "Poodle",
	 }

	 cachorro2EmJSON, erro := json.Marshal(c2)
	 if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(cachorro2EmJSON)
	fmt.Println(bytes.NewBuffer(cachorro2EmJSON))
}