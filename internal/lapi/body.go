package lapi

import (
	"bytes"
	"encoding/json"
	"io"
	"net/url"
	"strings"
)

// SetBody define o corpo da requisição HTTP usando um io.Reader.
// Esta função é útil quando você precisa definir um corpo personalizado
// ou quando está trabalhando com streams de dados.
//
// Parâmetros:
//   - body: Implementação de io.Reader contendo o corpo da requisição
//
// Exemplo:
//
//	file, _ := os.Open("arquivo.txt")
//	r.SetBody(file)
//
// Retorna a própria requisição para permitir encadeamento de métodos.
func (r *request) SetBody(body io.Reader) *request {
	r.body = body
	return r
}

// SetBodyString define o corpo da requisição HTTP como uma string.
// Útil para enviar dados em formato texto simples.
//
// Parâmetros:
//   - body: String contendo o corpo da requisição
//
// Exemplo:
//
//	r.SetBodyString("Hello, World!")
//
// Retorna a própria requisição para permitir encadeamento de métodos.
func (r *request) SetBodyString(body string) *request {
	r.body = strings.NewReader(body)
	return r
}

// SetBodyJSON define o corpo da requisição HTTP como um JSON.
// O valor será automaticamente convertido para uma string JSON.
// O header Content-Type será automaticamente definido como application/json.
//
// Parâmetros:
//   - body: Estrutura ou mapa a ser convertido para JSON
//
// Exemplo:
//
//	r.SetBodyJSON(map[string]interface{}{
//	    "name": "John",
//	    "age": 30,
//	    "active": true
//	})
//
// Retorna a própria requisição para permitir encadeamento de métodos.
// Se houver erro na conversão para JSON, o corpo não será definido.
func (r *request) SetBodyJSON(body interface{}) *request {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return r
	}
	r.body = bytes.NewReader(jsonBody)
	return r
}

// SetBodyFormData define o corpo da requisição HTTP como um FormData.
// Os valores serão automaticamente codificados para URL.
// O header Content-Type será automaticamente definido como application/x-www-form-urlencoded.
//
// Parâmetros:
//   - body: Mapa de campos do formulário (chave -> valor)
//
// Exemplo:
//
//	r.SetBodyFormData(map[string]string{
//	    "username": "john",
//	    "password": "secret",
//	    "remember": "true"
//	})
//
// Retorna a própria requisição para permitir encadeamento de métodos.
func (r *request) SetBodyFormData(body map[string]string) *request {
	formData := url.Values{}
	for key, value := range body {
		formData.Add(key, value)
	}

	r.body = strings.NewReader(formData.Encode())
	return r
}
