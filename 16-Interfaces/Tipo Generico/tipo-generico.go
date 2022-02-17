package main

import "fmt"
// como não passou nenhum parametro na interface ela vai aceitar qualquer tipo de dado
func generica(interf interface{}) {
	  fmt.Println(interf)
}


func main() {
   generica("string")
	 generica(1)
	 generica(true)

	 fmt.Println(1, 2, "string", false, true, float64(12345))

}