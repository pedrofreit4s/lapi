package lapi

import (
	"net/url"
)

// SetQuery define os parâmetros de query string para a requisição HTTP.
// Esta função substitui todos os parâmetros de query existentes.
// Os valores serão automaticamente codificados para URL.
//
// Parâmetros:
//   - query: Mapa de parâmetros de query (chave -> valor)
//
// Exemplo:
//
//	r.SetQuery(map[string]string{
//	    "page": "1",
//	    "limit": "10",
//	    "sort": "name",
//	    "filter": "active"
//	})
//	// Resultado: ?page=1&limit=10&sort=name&filter=active
//
// Retorna a própria requisição para permitir encadeamento de métodos.
func (r *request) SetQuery(query map[string]string) *request {
	queryString := url.Values{}
	for key, value := range query {
		queryString.Add(key, value)
	}

	r.query = queryString
	return r
}
