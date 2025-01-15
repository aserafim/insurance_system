package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type SininstrosChassi struct {
	IdConsulta   string
	Chassi       string
	QtdSinistros int
}

type ChassiRequest struct {
	Chassi string `json:"chassi"`
}

func main() {
	http.HandleFunc("/consulta-chassi", GetQtdSinistros)
	http.ListenAndServe(":8080", nil)
}

func GetQtdSinistros(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido. Use GET!", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erro ao ler o corpo da requisição", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var req ChassiRequest
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, "Erro ao decodificar JSON Request ", http.StatusBadRequest)
		return
	}

	if req.Chassi == "" {
		http.Error(w, "O parâmetro 'chassi' é obrigatório", http.StatusBadRequest)
		return
	}

	// FALTA CRIAR A CONEXAO COM O BANCO E
	// REALIZAR A CONSULTA DA QTD DE SINISTROS
	// PARA PREENCHER OS VALORES "REAIS"

	var resposta SininstrosChassi
	resposta.Chassi = req.Chassi
	resposta.IdConsulta = "123"
	resposta.QtdSinistros = 2

	retorno, err := json.Marshal(resposta)

	if err != nil {
		http.Error(w, "Erro ao processar JSON", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(retorno)
}
