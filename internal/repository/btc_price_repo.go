package repository

// Interface para o repositório de preço do Bitcoin

type BTCPriceRepository interface {
	GetBTCPrice() (float64, error)
}
