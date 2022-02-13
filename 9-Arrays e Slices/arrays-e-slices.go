package main

import (
	"fmt"
	"reflect"
)


func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5] string
	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [5] string {"Posição 1","Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)
	//outra forma de fazer um array
	array3 := [...] int {1, 2, 3, 4, 5}
	fmt.Println(array3)
	//não se pode acrescentar um novo item no array após ele ja declarado.ex: fez um array de 5 elementos, depois tenta colocar um sexto elemento.Ele não aceita
	//para isso o go criou o slice que é mais fexivel que o array
	/*--------------Slice-------------------*/

	slice := [] int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 18)
	fmt.Println(slice)

	slice2 := array2[1:3] //indice 1 inclusivo , indice 3 exclusivo
	fmt.Println(slice2)

  array2[1] = "Posição aAlterada"
	fmt.Println(slice2)

	//---------------------ARRAYS INTERNOS---------------------------------//

	  fmt.Println("-----------------------------------------------")

		slice3 := make([]float32, 10, 15) // A FUNÇÃO MAKE RECEBE 3 PARAMENTROS: O 1° REFERE-SE AO TIPO, A 2° REFERE-SE AO NUMERO DE POSIÇÕES QUE TERÁ, E A 3° REFERE-SE AO LIMITE MÁXIMO DE ITENS QUE O SLICE PODERÁ TER
		fmt.Println(slice3)
		fmt.Println(len(slice3)) //length 10
		fmt.Println(cap(slice3)) //capacidade máxima  15
		// A FUNÇÃO MAKE VAI CRIAR UM ARRAY DE 15 POSIÇÕES E VAI RETORNAR PARA NÓS UM SLICE QUE VAI PEGAR ESSA 10 PRIMEIRAS POSIÇÕES DESSE ARRAY
		//O SLICE SEMPRE REFERENCIA UM ARRAY...ELE É COMO SE FOSSE UMA FATIA DE UM ARRAY
		fmt.Println("---------------------------------------------------------")

		slice4 := make([]int32, 10, 11)
		fmt.Println(slice4)

		slice4 = append(slice4, 5)

		fmt.Println(slice4)


		fmt.Println(len(slice4)) //length 11
		fmt.Println(cap(slice4)) //capacidade 11

		fmt.Println("-----------------------------------------------------------")

		slice4 = append(slice4, 8)

		fmt.Println(slice4)

    //automaticamente o slice vai criar um array interno dobrando o valor que ja tinha para se adequar a nova quantidade
		fmt.Println(len(slice4)) //length 12
		fmt.Println(cap(slice4)) //capacidade 24

		/*-------------------------------------------------------*/
		fmt.Println("-------------------------------------------------------")

		//O ULTIMO PARAMETRO DA FUNÇÃO MAKE NÃO É OBRIGATÓRIA DE PASSAR EX:

		slice5 := make([]float32, 5)
		fmt.Println(slice5)
 
    slice5 = append(slice5, 5)//acrescentando mais um item o slice vai ficar com 6 itens e o array interno do go vai para 12 itens

		fmt.Println(slice5)
    

		fmt.Println(len(slice5)) //length 
		fmt.Println(cap(slice5)) //capacidade 





}