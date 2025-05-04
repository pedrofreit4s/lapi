package lapi

import (
	"net/http"
	"net/url"
)

// OutOfContext retorna uma nova requisição HTTP fora do contexto.
// Esta função é útil quando você precisa fazer uma requisição HTTP
// sem usar o contexto padrão do modelo.
//
// Exemplo:
//
//	r := OutOfContext()
//	r.SetBaseURL("https://api.exemplo.com")
//	r.SetMethod("GET")
//	resp, err := r.Send()
//
// Retorna uma nova instância de request com headers e query parameters vazios.
func OutOfContext() *request {
	return &request{
		headers: make(map[string]string),
		query:   make(url.Values),
	}
}

// Send envia a requisição HTTP e retorna a resposta.
// Esta função é responsável por:
// 1. Criar uma nova requisição HTTP com os parâmetros configurados
// 2. Adicionar os headers definidos
// 3. Enviar a requisição usando o cliente HTTP padrão
// 4. Retornar a resposta ou um erro, se ocorrer algum problema
//
// Exemplo:
//
//	r := OutOfContext()
//	r.SetBaseURL("https://api.exemplo.com")
//	r.SetMethod("GET")
//	resp, err := r.Send()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer resp.Body.Close()
//
// Retorna:
//   - *http.Response: Resposta HTTP
//   - error: Erro, se ocorrer algum problema durante a requisição
func (r *request) Send() (*http.Response, error) {
	req, err := http.NewRequest(r.method, r.baseURL, r.body)
	if err != nil {
		return nil, err
	}
	for key, value := range r.headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
