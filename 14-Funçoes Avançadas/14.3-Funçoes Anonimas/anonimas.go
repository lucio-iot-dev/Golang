package main

import "fmt"


func main() {

    func () {

			fmt.Println("Olá Mundo")

		}()
		/*-------------------------------*/
		//SEGUNDA FORMA
		func (texto string)  {
			fmt.Println(texto)
		}("Passando Parâmetro")

		/*-----------------------------------------*/
		//OUTRA FORMA DE FAZER 

		retorno := func(texto string) string {
			return fmt.Sprintf("Recebido -> %s %d", texto, 10) //PARA CONCATENAR COM UMA STRING USA %s ,PARA CONCATENAR COM UM NUMEROS USA %d
		}("Passando Parâmetros")

		fmt.Println(retorno)
					
		}


