package main

import "fmt"

//PRIMEIRA FORMA DE FAZER
// func diaDaSemana(numero int) string {
 
// 	switch numero {
// 	case 1:
// 		   return "Domingo"
// 	case 2:
// 		return "Segunda-Feira"		
// 	case 3: 
// 	   return "Terça-Feira"
// 	case 4:	 
// 	  return "Quarta-Feira"
// 	case 5:
// 		return "Quinta-feira"		 
// 	case 6:
// 		return "Sexta-Feira"
// 	case 7:
// 		return "Domingo"
// 	default:
// 		return "Número Inválido"	
// 	}
// }
/*-------------------------------*/
//SEGUNDA  FORMA DE FAZER
// func diaDaSemana2(numero int) string {
  
// 	switch {
// 	case numero == 1:
// 		   return "Domingo"
// 	case numero == 2:
// 		return "Segunda-Feira"		
// 	case numero == 3: 
// 	   return "Terça-Feira"
// 	case numero == 4:	 
// 	  return "Quarta-Feira"
// 	case numero == 5:
// 		return "Quinta-feira"		 
// 	case numero == 6:
// 		return "Sexta-Feira"
// 	case numero == 7:
// 		return "Domingo"
// 	default:
// 		return "Número Inválido"	
// 	}
// }

//TERCEIRA FORMA DE FAZER

func diaDaSemana3(numero int) string {
	 var diaDaSemana string
  
	switch {
	case numero == 1:
		   diaDaSemana = "Domingo"
			 //fallthrough este comando é pouco utilizado ,mas funciona da seguinte forma: 
			 //caso coloquemos 1 na função main ele teria que retornar domingo ,mas ele pula direto para o valor de baixo retornando segunda sem nem mesmo avaliar a condição
			 //é um metodo pouco utilizado
	case numero == 2:
		diaDaSemana = "Segunda-Feira"		
	case numero == 3: 
	   diaDaSemana = "Terça-Feira"
	case numero == 4:	 
	  diaDaSemana = "Quarta-Feira"
	case numero == 5:
		diaDaSemana = "Quinta-feira"		 
	case numero == 6:
		diaDaSemana = "Sexta-Feira"
	case numero == 7:
		diaDaSemana = "Domingo"
	default:
		diaDaSemana = "Número Inválido"	
	}
	return diaDaSemana
}


func main() {
	fmt.Println("Switch")

	dia := diaDaSemana3(5)
	fmt.Println(dia)
}