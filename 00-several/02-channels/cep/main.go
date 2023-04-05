package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Endereco representa um endereço de CEP
type Endereco struct {
	CEP         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
}

// BuscaEndereco busca o endereço de um CEP em várias APIs e retorna o endereço mais rápido
func BuscaEndereco(cep string) (*Endereco, error) {
	// Cria um canal para receber os endereços das APIs
	enderecos := make(chan *Endereco, 3)

	// Cria uma função para buscar o endereço em uma API e enviar o resultado para o canal
	buscarAPI := func(api string) {
		url := fmt.Sprintf(api, cep)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Erro ao buscar endereço na API %s: %s\n", api, err)
			enderecos <- nil
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Erro ao buscar endereço na API %s: %s\n", api, resp.Status)
			enderecos <- nil
			return
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Erro ao ler resposta da API %s: %s\n", api, err)
			enderecos <- nil
			return
		}

		var endereco Endereco
		err = json.Unmarshal(body, &endereco)
		if err != nil {
			fmt.Printf("Erro ao decodificar resposta da API %s: %s\n", api, err)
			enderecos <- nil
			return
		}

		enderecos <- &endereco
	}

	// Chama as APIs em goroutines
	go buscarAPI("http://viacep.com.br/ws/%s/json/")
	go buscarAPI("https://ws.apicep.com/cep/%s.json")

	// Espera pelo resultado mais rápido e retorna o endereço correspondente
	var endereco *Endereco
	for i := 0; i < 3; i++ {
		select {
		case e := <-enderecos:
			if e != nil {
				if endereco == nil || len(e.CEP) < len(endereco.CEP) {
					endereco = e
				}
			}
		case <-time.After(2 * time.Second):
			break
		}
	}

	if endereco == nil {
		return nil, fmt.Errorf("Não foi possível obter o endereço do CEP %s", cep)
	}

	return endereco, nil
}

func main() {
	endereco, err := BuscaEndereco("89160362")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s, %s - %s/%s\n", endereco.Logradouro, endereco.Bairro, endereco.Localidade, endereco.UF)
	}
}
