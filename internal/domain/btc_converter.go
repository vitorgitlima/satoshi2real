package domain

// Converter responsável por armazenar a lógica de conversão

type BTCConverter struct {
	Satoshis int64
	BTCRate  float64
}

// Método para converter satoshis para reais
func (c *BTCConverter) ConvertToReais() float64 {
	btc := float64(c.Satoshis) / 100_000_000
	reais := btc * c.BTCRate

	return reais
}
