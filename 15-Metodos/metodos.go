package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do Usuário %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
   usuario1 := usuario{"Lucio", 44}
	 fmt.Println(usuario1)
	 usuario1.salvar()

	 usuario2 := usuario{"Davi", 30}
	 usuario2.salvar()

	 maiorDeIdade := usuario2.maiorDeIdade()
	 fmt.Println(maiorDeIdade)

	 usuario2.fazerAniversario()
	 fmt.Println(usuario2.idade)

	 usuario1.fazerAniversario()
	 fmt.Println(usuario1.nome, usuario1.idade)
}