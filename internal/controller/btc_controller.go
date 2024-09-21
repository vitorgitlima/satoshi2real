package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"satoshi2real/internal/usecase" // Corrigido para o caminho correto
	"strconv"
)

type BTCController struct {
	ConvertUseCase *usecase.ConvertSatoshiUseCase
}

func (c *BTCController) ConvertHandler(w http.ResponseWriter, r *http.Request) {
	// obtendo satoshis via query parameter
	satoshisStr := r.URL.Query().Get("satoshis")
	if satoshisStr == "" {
		http.Error(w, "Faltando o parametro satoshis", http.StatusBadRequest)
		return
	}

	satoshis, err := strconv.ParseInt(satoshisStr, 10, 64)
	if err != nil {
		http.Error(w, "Parâmetro inválido", http.StatusBadRequest)
		return
	}

	// Convertendo satoshis para reais usando o caso de uso
	reais, err := c.ConvertUseCase.Convert(satoshis)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro na conversão: %v", err), http.StatusInternalServerError)
		return
	}

	// Respondendo com o valor em reais
	response := map[string]float64{"reais": reais}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
