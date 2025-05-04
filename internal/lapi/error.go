package lapi

// HttpError é uma interface que representa um erro HTTP.
// Ela fornece métodos para acessar o código de status, a requisição e a resposta associadas ao erro.
//
// Exemplo de uso:
//
//	var err HttpError
//	if err != nil {
//	    fmt.Printf("Erro %d: %s\n", err.StatusCode(), err.Message())
//	}
type HttpError interface {
	// StatusCode retorna o código de status HTTP do erro.
	// Exemplo: 404 para Not Found, 500 para Internal Server Error
	StatusCode() int

	// Message retorna a mensagem de erro associada ao erro.
	// Esta mensagem é amigável para o usuário e pode ser exibida diretamente.
	Message() string

	// Request retorna a requisição que gerou o erro.
	// Pode ser usado para debug ou logging.
	Request() interface{}

	// Response retorna a resposta associada ao erro.
	// Pode conter detalhes adicionais sobre o erro retornado pela API.
	Response() interface{}
}

// httpError é uma implementação concreta de HttpError.
// Ela armazena o código de status, a mensagem de erro, a requisição e a resposta associada ao erro.
type httpError struct {
	// statusCode é o código de status HTTP do erro.
	// Exemplo: 404 para Not Found, 500 para Internal Server Error
	statusCode int

	// message é a mensagem de erro amigável para o usuário.
	message string

	// request é a requisição que gerou o erro.
	// Pode ser usado para debug ou logging.
	request interface{}

	// response é a resposta associada ao erro.
	// Pode conter detalhes adicionais sobre o erro retornado pela API.
	response interface{}
}

// Message retorna a mensagem de erro associada ao erro.
// Esta mensagem é amigável para o usuário e pode ser exibida diretamente.
func (e *httpError) Message() string {
	return e.message
}

// Request retorna a requisição que gerou o erro.
// Pode ser usado para debug ou logging.
func (e *httpError) Request() interface{} {
	return e.request
}

// Response retorna a resposta associada ao erro.
// Pode conter detalhes adicionais sobre o erro retornado pela API.
func (e *httpError) Response() interface{} {
	return e.response
}

// StatusCode retorna o código de status HTTP do erro.
// Exemplo: 404 para Not Found, 500 para Internal Server Error
func (e *httpError) StatusCode() int {
	return e.statusCode
}

// NewError cria um novo erro HTTP com todos os detalhes.
//
// Parâmetros:
//   - statusCode: Código de status HTTP (ex: 404, 500)
//   - message: Mensagem de erro amigável para o usuário
//   - request: Requisição que gerou o erro
//   - response: Resposta associada ao erro
//
// Exemplo:
//
//	err := m.NewError(404, "Usuário não encontrado", request, response)
func (m *model) NewError(statusCode int, message string, request interface{}, response interface{}) HttpError {
	return &httpError{
		statusCode: statusCode,
		message:    message,
		request:    request,
		response:   response,
	}
}

// MakeError cria um novo erro HTTP com código de status e mensagens.
//
// Parâmetros:
//   - statusCode: Código de status HTTP (ex: 404, 500)
//   - err: Mensagem de erro técnica (para debug)
//   - message: Mensagem de erro amigável para o usuário
//
// Exemplo:
//
//	err := m.MakeError(500, "database connection failed", "Erro interno do servidor")
func (m *model) MakeError(statusCode int, err, message string) *httpError {
	return &httpError{
		statusCode: statusCode,
		message:    message,
	}
}

// Error implementa a interface error do Go.
// Retorna a mensagem de erro amigável para o usuário.
func (e *httpError) Error() string {
	return e.message
}
