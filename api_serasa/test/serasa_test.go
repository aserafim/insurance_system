package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type CpfConsulta struct {
	Cpf string `json:"cpf"`
}

type ScoreRetorno struct {
	Cpf   string
	Score string
}

func TestMetodoDiferenteDeGet(t *testing.T) {

	ret, err := http.Get("http://localhost:8080/consulta-score")
	if err != nil {
		fmt.Println("Erro ao realizar chamada GET")
	}

	body, err := io.ReadAll(ret.Body)
	if err != nil {
		fmt.Println("Erro ao consumir body")
	}
	assert.Equal(t, "Método não disponível. Use POST!\n", string(body))
}

func TestRetornoSucesso(t *testing.T) {

	//Cria o objeto de consulta
	cpf := CpfConsulta{
		Cpf: "12345678911",
	}

	//Faz o parse para o JSON da requisição
	payload, err := json.Marshal(cpf)
	if err != nil {
		fmt.Println("Erro ao criar input")
	}

	//Realiza a chamada
	response, err := http.Post("http://localhost:8080/consulta-score", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Erro na chamada")
	}

	//Consome o body de retorno
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Erro ao ler o corpo da resposta")
	}

	//Guarda o retorno em um JSON
	var jsonScore ScoreRetorno
	err = json.Unmarshal(body, &jsonScore)
	if err != nil {
		fmt.Println("Erro ao parsear retorno")
	}

	//Verifica os valores retornados
	assert.Equal(t, "12345678911", jsonScore.Cpf)
	assert.Equal(t, "590", jsonScore.Score)

}
