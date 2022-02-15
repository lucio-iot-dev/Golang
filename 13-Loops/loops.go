package main

import (
	"fmt"
	// "time"
)

func main() {
	// i := 0 

	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementando i")
	// 	fmt.Println(i)
	// 	time.Sleep(time.Second)

	// }
	// fmt.Println(i)

	// for j:= 0; j <= 10; j++ {
	// 	fmt.Println("Incrementando j", j)
	// 	time.Sleep(time.Second)
	// }

	// nomes := [3]string {"João", "Davi", "Lucas"}

	// for indice, nome := range nomes {
	// 	fmt.Println(indice, nome)
	// }

//retorna sempre primeiro o indice e depois o nome
//caso não queira o indice ,para ignorar o primeiro item deve-se usar o anderline como mostra abaixo

/*----------------------------------------------------------*/
	nomes := [3]string {"João", "Davi", "Lucas"}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra)) // NESSE CASO É NECESSÁRIO COLOCAR STRING ANTES DE LETRA PARA O GO NÃO RETORNAR O NUMERO REFERENTE A LETRA PELA TABELA ASC
	}

}