package lapi

import "encoding/json"

// SetHeaders define múltiplos cabeçalhos HTTP para a requisição de uma vez.
// Esta função substitui todos os cabeçalhos existentes.
//
// Parâmetros:
//   - headers: Mapa de cabeçalhos HTTP (chave -> valor)
//
// Exemplo:
//
//	r.SetHeaders(map[string]string{
//	    "Content-Type": "application/json",
//	    "Accept": "application/json",
//	    "X-Custom-Header": "valor"
//	})
//
// Retorna a própria requisição para permitir encadeamento de métodos.
func (r *request) SetHeaders(headers map[string]string) *request {
	r.headers = headers
	return r
}

// SetHeader define um único cabeçalho HTTP para a requisição.
// Se o cabeçalho já existir, seu valor será substituído.
//
// Parâmetros:
//   - key: Nome do cabeçalho HTTP
//   - value: Valor do cabeçalho HTTP
//
// Exemplo:
//
//	r.SetHeader("Content-Type", "application/json")
//
// Retorna a própria requisição para permitir encadeamento de métodos.
func (r *request) SetHeader(key, value string) *request {
	r.headers[key] = value
	return r
}

// SetHeaderJSON define um cabeçalho HTTP com um valor JSON para a requisição.
// O valor será automaticamente convertido para uma string JSON.
//
// Parâmetros:
//   - key: Nome do cabeçalho HTTP
//   - value: Valor a ser convertido para JSON
//
// Exemplo:
//
//	r.SetHeaderJSON("X-Custom-Data", map[string]string{
//	    "key": "value"
//	})
//
// Retorna a própria requisição para permitir encadeamento de métodos.
// Se houver erro na conversão para JSON, o cabeçalho não será definido.
func (r *request) SetHeaderJSON(key string, value interface{}) *request {
	jsonBody, err := json.Marshal(value)
	if err != nil {
		return r
	}
	r.headers[key] = string(jsonBody)
	return r
}
