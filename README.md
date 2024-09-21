# satoshi2real

O **satoshi2real** é uma aplicação Go que converte valores de satoshis (a menor unidade do Bitcoin) para reais (BRL) utilizando a API da CoinGecko.

## Funcionalidades

- Converte valores em satoshis para reais.
- Utiliza a API da CoinGecko para obter o preço atual do Bitcoin em reais.
- Responde a requisições HTTP com um formato JSON.

## Instalação

1. Clone o repositório:
   ```bash
   git clone https://github.com/vitorgitlima/satoshi2real.git
   cd satoshi2real
   ```
## Exemplo de Uso

Para converter um valor de satoshis para reais, você pode fazer uma requisição HTTP GET para o endpoint `/convert`. 

### Requisição

- **URL**: `http://localhost:8080/convert?satoshis=1000`
- **Método**: GET
- **Parâmetros**:
  - `satoshis`: A quantidade de satoshis que você deseja converter. Por exemplo, para converter 1.000 satoshis, use `?satoshis=1000`.

```bash
go run main.go
  ```
Você pode testar a API diretamente no navegador ou usar ferramentas como cURL ou Postman. Aqui está um exemplo usando cURL

```bash
curl "http://localhost:8080/convert?satoshis=1000"
 ```

## Resposta
{
"reais": 3486.4700000000003
}

