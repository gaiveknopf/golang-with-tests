package main

import (
	"testing"
)

func TestBuscaEndereco(t *testing.T) {
	endereco, err := BuscaEndereco("04571010")
	if err != nil {
		t.Errorf("Erro ao buscar endereço: %s", err)
	}
	if endereco == nil || endereco.Localidade != "São Paulo" {
		t.Errorf("Endereço incorreto: %v", endereco)
	}

	endereco, err = BuscaEndereco("12345678")
	if err == nil || endereco != nil {
		t.Errorf("Busca por CEP inválido deveria retornar erro e nil, mas retornou: %v, %v", endereco, err)
	}
}
