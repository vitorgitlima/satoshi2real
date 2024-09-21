package main

import (
	"log"
	"net/http"
	"satoshi2real/internal/controller" // Importando o controller
	"satoshi2real/internal/usecase"
	"satoshi2real/pkg/api"
)

func main() {
	// Inicializando o repositório da CoinGecko
	coinGeckoRepo := &api.CoinGeckoRepository{}

	// Inicializando o caso de uso de conversão
	convertUseCase := &usecase.ConvertSatoshiUseCase{
		BTCPriceRepo: coinGeckoRepo,
	}

	// Inicializando o controlador
	btcController := &controller.BTCController{
		ConvertUseCase: convertUseCase,
	}

	// Configurando o handler da rota
	http.HandleFunc("/convert", btcController.ConvertHandler)

	// Rodando o servidor
	log.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
