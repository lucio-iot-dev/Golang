package main

import (
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
    cachorroEmJSON := `{"nome":"Rex","raca":"Dalmata","idade":3}`

		var c cachorro
                              //slice de byte         ponteirto de c
		if erro := json.Unmarshal([]byte(cachorroEmJSON), &c); erro != nil {
			   log.Fatal(erro)
		}

		fmt.Println(c) // vai retornar um struct

		cachorro2EmJSON := `{"nome":"Toby", "raca":"Poodle"}`

		c2 := make(map[string]string)
		if erro := json.Unmarshal([]byte(cachorro2EmJSON), &c2); erro != nil {
			log.Fatal(erro)
		}
		fmt.Println(c2)
}