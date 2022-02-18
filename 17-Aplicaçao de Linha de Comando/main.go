package main


import (
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
		aplicacao := app.Gerar()
		if erro := aplicacao.Run(os.Args); erro != nil {  //Para rodar a aplicação vamos invocar um Metodo chamado Run e ele recebe um argumento padrão (os.Args) ,para que os comandos do sistema operacional sejam reconhecidos pela linha de comando
			log.Fatal(erro) // O .Fatal fará parar a execução do programa caso ocorra erro
		}
}