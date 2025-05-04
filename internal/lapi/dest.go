package lapi

// SetDest define o destino para a resposta HTTP.
// Esta função permite especificar uma estrutura ou variável
// que receberá os dados da resposta após a deserialização.
//
// Parâmetros:
//   - dest: Ponteiro para a estrutura ou variável que receberá a resposta
//
// Exemplo:
//
//	var response struct {
//	    ID    int    `json:"id"`
//	    Name  string `json:"name"`
//	    Email string `json:"email"`
//	}
//	r.SetDest(&response)
//
// Retorna a própria requisição para permitir encadeamento de métodos.
//
// TODO: Implementar a função SetDest
func (r *request) SetDest(dest interface{}) *request {
	// TODO: Implementar a função SetDest
	return r
}
