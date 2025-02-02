// TESTE DE UNIDADE

package enderecos

import "testing"




type cenarioDeTeste struct {
	  enderecoInserido string
		retornoEsperado string
}


func TestTipoDeEndereco(t *testing.T)  {
    
	 cenariosDeTeste := []cenarioDeTeste {
		   {"Rua ABC", "Rua"},
		   {"Avenida Paulista", "Avenida"},
		   {"Rodovia dos Imigrantes", "Rodovia"},
		   {"Praça das Rosas", "Tipo Inválido"},
		   {"Estrada qualquer", "Estrada"},
		   {"Rua dos Bobos", "Rua"},
		   {"Avenida Rebouças", "Avenida"},
		   {"", "Tipo Inválido"},
		 }

		 for _,cenario := range cenariosDeTeste {
       
			   retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)

				 if retornoRecebido != cenario.retornoEsperado {
					t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s",
				 retornoRecebido,
				 cenario.retornoEsperado,
		 )
		}
	}
}

		 

	
