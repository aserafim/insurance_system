package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMetodoDiferenteDeGet(t *testing.T) {

	_, err := http.Post("http://localhost:8080/consulta-score" + "/45069844753", "application/json", nil)

	//assert.Equal(t, "404 page not found", resp.Body)
	assert.Error(t, err, http.StatusBadRequest)
		
}
