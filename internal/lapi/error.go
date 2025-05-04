package lapi

// HttpError é uma interface que representa um erro HTTP.
// Ela fornece métodos para acessar o código de status, a requisição e a resposta associadas ao erro.
type HttpError interface {
	// StatusCode retorna o código de status HTTP do erro.
	StatusCode() int
	// Error retorna a mensagem de erro associada ao erro.
	Message() string
	// Request retorna a requisição que gerou o erro.
	Request() interface{}
	// Response retorna a resposta associada ao erro.
	Response() interface{}
}

// error é uma implementação concreta de HttpError.
// Ela armazena o código de status, a mensagem de erro, a requisição e a resposta associada ao erro.
type httpError struct {
	statusCode int
	message    string
	request    interface{}
	response   interface{}
}

// Message implements HttpError.
func (e *httpError) Message() string {
	return e.message
}

// Request implements HttpError.
func (e *httpError) Request() interface{} {
	return e.request
}

// Response implements HttpError.
func (e *httpError) Response() interface{} {
	return e.response
}

// StatusCode implements HttpError.
func (e *httpError) StatusCode() int {
	return e.statusCode
}

// NewError cria um novo erro HTTP.
func (m *model) NewError(statusCode int, message string, request interface{}, response interface{}) HttpError {
	return &httpError{
		statusCode: statusCode,
		message:    message,
		request:    request,
		response:   response,
	}
}

// MakeError cria um novo erro HTTP.
func (m *model) MakeError(statusCode int, err, message string) *httpError {
	return &httpError{
		statusCode: statusCode,
		message:    message,
	}
}

func (e *httpError) Error() string {
	return e.message
}
