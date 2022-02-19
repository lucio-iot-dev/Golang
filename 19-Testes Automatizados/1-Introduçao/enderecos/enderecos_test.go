// TESTE DE UNIDADE

package enderecos

import (
	"testing"
	"fmt"
)


func TestTipoDeEndereco(t *testing.T)  {
	  enderecoParaTeste := "Avenida Paulista"
		tipoDeEnderecoEsperado := "Avenida"
		tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

		if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
			   t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s",
				tipoDeEnderecoEsperado,
			tipoDeEnderecoRecebido,
		)
		} else {
			fmt.Println("---------------------------------------------------------------------------------------------------------")

			fmt.Printf("Funcionou ok! Teste realizado com sucesso! O valor esperado: %s, foi igual o valor recebido: %s",
		tipoDeEnderecoEsperado,
	tipoDeEnderecoRecebido,
	)
	  fmt.Println("   ")
	  fmt.Println("---------------------------------------------------------------------------------------------------------")

		}
}