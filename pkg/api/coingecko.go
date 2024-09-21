package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Interface para o client da API do CoinGecko

type CoinGeckoClient interface {
	Bitcoin struct {
		Brl float64 `json:"brl"`
	} `json:"bitcoin"`
}

// implementação do repositório usando CoinGecko

type CoinGeckoRepository struct {}

// Implementação de método para obter o preço do Bitcoin

fun (r *CoinGeckoRepository) GetBTCPriceInBRL() (float64, error){
	url := "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=brl"
	resp, err := http.Get(url)

	if err != nil {
		return 0, fmt.Errorf("erro ao fazer requisição: %v", err)
	}
	defer resp.Body.Close()

	var result CoinGeckoResponse

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, fmt.Errorf("erro ao decodificar resposta: %v", err)
	}

	return result.Bitcoin.Brl, nil
}