package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CpfConsulta struct {
	Cpf string `json:"cpf"`
}

type ScoreRetorno struct {
	Cpf   string
	Score string
}

func ConsultaScore(w http.ResponseWriter, r *http.Request) {

	// Verifica se o método está correto
	if r.Method != http.MethodPost {
		http.Error(w, "Método não disponível. Use POST!", http.StatusBadRequest)
	}

	// Lê o corpo da requisição
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erro ao ler a requisição.", http.StatusBadRequest)
	}
	defer r.Body.Close()

	// Faz o parse do JSON
	var req CpfConsulta
	err = json.Unmarshal(body, &req)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Erro ao desserializar JSON (Request).", http.StatusBadRequest)
	}

	// Cria um MOCK do resultado
	score := ScoreRetorno{Cpf: req.Cpf, Score: "590"}

	// Faz o Marshal do resultado
	ret, err := json.Marshal(score)
	if err != nil {
		http.Error(w, "Erro ao tentar serializar Response", http.StatusBadRequest)
	}

	// Grava retorno na saída
	w.Header().Set("Content-Type", "application/json")
	w.Write(ret)

}

func main() {

	http.HandleFunc("/consulta-score", ConsultaScore)
	http.ListenAndServe(":8080", nil)

}
