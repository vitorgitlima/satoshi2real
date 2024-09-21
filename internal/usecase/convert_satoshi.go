package usecase

import (
	"myapp/internal/domain"
	"myapp/internal/repository"
)

// Interactor que conecta o repositório e o domínio
type ConvertSatoshiUseCase struct {
	BTCPriceRepo repository.BTCPriceRepository
}

// Método para executar a conversão
func (u *ConvertSatoshiUseCase) Convert(satoshis int64) (float64, error) {
	btcRate, err := u.BTCPriceRepo.GetBTCPrice()

	if err != nil {
		return 0, err
	}

	converter := domain.BTCConverter{
		Satoshis: satoshis,
		BTCRate:  btcRate,
	}

	return converter.ConvertToReais(), nil
}
