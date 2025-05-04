package lapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Request representa uma requisição HTTP.
// Contém informações necessárias para realizar uma requisição HTTP como
// URL base, método HTTP, cabeçalhos e corpo da requisição.
type model struct {
	// request representa a requisição HTTP.
	request *request
	// Auth representa as informações de autenticação.
	Auth struct {
		// Token representa o token de autenticação.
		Token string
		// RefreshToken representa o token de atualização.
		RefreshToken string
	}
	// inDevelopment representa se a requisição está em ambiente de desenvolvimento.
	inDevelopment bool
}

// NewRequest cria uma nova requisição HTTP.
func NewRequest(baseURL string, headers map[string]string, timeout int) *model {
	return &model{
		request: &request{
			baseURL: baseURL,
			headers: headers,
			method:  "GET",
			timeout: time.Duration(timeout) * time.Second,
			query:   make(url.Values),
		},
	}
}

func (m *model) MakeRequest(method string, path string, payload *interface{}, dest interface{}) *httpError {
	m.request.method = method
	// Parse the body
	if payload != nil {
		bodyJson, err := json.Marshal(payload)
		if err != nil {
			return m.MakeError(http.StatusInternalServerError, err.Error(), "Houve um erro interno no servidor! C: 01")
		}
		m.request.body = bytes.NewBuffer(bodyJson)
	}

	// Parse the query
	qp := ""
	if strings.Contains(m.request.baseURL, "?") {
		qp = "&"
	} else {
		qp = "?"
	}

	// Make the request
	req, err := http.NewRequest(
		m.request.method,
		fmt.Sprintf("%s%s%s%s", m.request.baseURL, path, qp, m.request.query.Encode()),
		m.request.body,
	)
	if err != nil {
		return m.MakeError(http.StatusInternalServerError, err.Error(), "Houve um erro interno no servidor! C: 02")
	}

	// Parse the headers
	for k, v := range m.request.headers {
		req.Header.Add(k, v)
	}

	fmt.Println(req.Header)

	// Parse the auth
	if m.Auth.Token != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", m.Auth.Token))
	}

	// Perform request
	client := &http.Client{
		Timeout: m.request.timeout,
	}
	start := time.Now()
	resp, err := client.Do(req)

	responseTime := time.Since(start).Milliseconds()
	status := "(408 timeout)"

	if resp != nil {
		status = resp.Status
	}

	// Log request
	logMsg := fmt.Sprintf("[%s] %s (%s) %d ms", req.Method, req.URL, status, responseTime)
	log.Println(logMsg)

	if err != nil {
		return m.MakeError(http.StatusInternalServerError, err.Error(), "Houve um erro interno no servidor! C: 03")
	}
	defer resp.Body.Close()

	// Parse response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return m.MakeError(http.StatusInternalServerError, err.Error(), "Houve um erro interno no servidor! C: 04")
	}

	if resp.Header.Get("Content-Type") == "application/pdf" {
		*dest.(*[]byte) = body
		return nil
	}

	if err := json.Unmarshal(body, dest); err != nil {
		fmt.Println(string(body))
		log.Println(err.Error())
		return nil
	}

	// Check status code
	if resp.StatusCode >= 400 {
		return m.MakeError(resp.StatusCode, string(body), "Status code >= 400")
	}

	return nil
}

func (m *model) Request() *request {
	return m.request
}

// Get faz uma requisição GET.
func (m *model) Get(path string, dest interface{}) *httpError {
	r := m.MakeRequest("GET", path, nil, dest)

	return r
}

// Post faz uma requisição POST.
func (m *model) Post(path string, payload *interface{}, dest interface{}) *httpError {
	r := m.MakeRequest("POST", path, payload, dest)

	return r
}

// Put faz uma requisição PUT.
func (m *model) Put(path string, payload *interface{}, dest interface{}) *httpError {
	r := m.MakeRequest("PUT", path, payload, dest)

	return r
}

// Delete faz uma requisição DELETE.
func (m *model) Delete(path string, dest interface{}) *httpError {
	r := m.MakeRequest("DELETE", path, nil, dest)

	return r
}

// Patch faz uma requisição PATCH.
func (m *model) Patch(path string, payload *interface{}, dest interface{}) *httpError {
	r := m.MakeRequest("PATCH", path, payload, dest)

	return r
}
