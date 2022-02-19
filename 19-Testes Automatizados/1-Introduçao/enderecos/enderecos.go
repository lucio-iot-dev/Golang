package enderecos

import "strings"

//TipoDeEndereco verifica se um endereço tem um tipo válido e o retorna
func TipoDeEndereco(endereco string) string {
    tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

    enderecoEmLetraMinuscula := strings.ToLower(endereco)
		primeiraPalavraDoEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0]
     // O COMANDO : strings.Split 
		//"RUA DOS BOBOS"  Pega uma string e tranforma em um slice
		//["RUA", "DOS", "BOBOS"] Slice

    enderecoTemUmTipoValido := false
		for _, tipo := range tiposValidos {
			  if tipo == primeiraPalavraDoEndereco {
					enderecoTemUmTipoValido = true
				}
		}

		if enderecoTemUmTipoValido {
			return strings.Title(primeiraPalavraDoEndereco) //vai colocar a primeira letra em maiusculo
		}
		return "Tipo Inválido"
}