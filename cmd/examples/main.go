package main

import (
	"fmt"
	"io"
	"log"
	"strings"

	lapi "github.com/pedrofreit4s/lapi/internal/lapi/request"
)

func main() {
	// Criando uma requisição fora do contexto.
	r := lapi.OutOfContext()

	r.SetMethod("GET")

	r.SetBaseURL("https://example.com")

	r.SetHeader("Content-Type", "application/json")
	r.SetHeader("Authorization", "Bearer 1234567890")

	r.SetQuery(map[string]string{
		"key": "value",
	})
	r.SetBody(strings.NewReader("{}"))

	r.SetDest(nil)

	// Enviando a requisição.
	resp, err := r.Send()
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
