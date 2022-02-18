package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retornar a aplicação de linha de comando pronta para ser executada
func Gerar() *cli.App {    // função Gerar() com letra maiúscula para que o pacote main importe essa função
	 app := cli.NewApp()
	  app.Name = "Aplicação de linha de Comando"
	  app.Usage = "Busca IPs e Nomes de Servidor na Internet"

		flags := []cli.Flag {
			cli.StringFlag{
				Name: "host",
				Value: "devbook.com.br",
			},
		}

		app.Commands = []cli.Command{ //commands é do tipo slice/////command é do tipo Struct
			{
				Name: "ip",
				Usage: "Busca IPs de endereços na internet",
				Flags:flags,
				Action: buscarIps,
			},
			{
				 Name: "servidores",
				 Usage: "Busca o nome dos servodores na internet",
				 Flags: flags,
				 Action: buscarServidores,
			},
		}
	 
	  return app

}

func buscarIps(c *cli.Context) {
    host := c.String("host")

		ips, erro := net.LookupIP(host)
		if erro != nil {
			log.Fatal(erro)
		}

		for _, ip := range ips {
			fmt.Println(ip)
		}
}

func buscarServidores(c *cli.Context) {
	 host := c.String("host")

	 servidores, erro := net.LookupNS(host) // name server
	 if erro != nil {
		 log.Fatal(erro)
	 }

	 for _,servidor := range servidores {
		  fmt.Println(servidor.Host)
	 }
}




