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

	//assert.Equal(t, "404 page not found", resp.Body)
	//assert.Error(t, err, http.StatusBadRequest)

}

func TestRetornoSucesso(t *testing.T) {

	cpf := CpfConsulta{
		Cpf: "12345678911",
	}

	payload, err := json.Marshal(cpf)
	if err != nil {
		fmt.Println("Erro ao criar input")
	}

	response, err := http.Post("http://localhost:8080/consulta-score", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Erro na chamada")
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Erro ao ler o corpo da resposta")
	}

	var retornoJson ScoreRetorno
	err = json.Unmarshal(body, &retornoJson)
	if err != nil {
		fmt.Println("Erro ao parsear retorno")
	}
	assert.Equal(t, "12345678911", retornoJson.Cpf)
	assert.Equal(t, "590", retornoJson.Score)

}
