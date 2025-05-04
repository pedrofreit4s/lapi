package main

import (
	"log"

	"github.com/pedrofreit4s/lapi/internal/lapi"
)

func main() {
	// Criando uma requisição fora do contexto.
	// r := lapi.OutOfContext()

	// r.SetMethod("GET")

	// r.SetBaseURL("https://example.com")
	// r.SetHeader("Content-Type", "application/json")
	// r.SetHeader("Authorization", "Bearer 1234567890")

	// r.SetQuery(map[string]string{
	// 	"key": "value",
	// })
	// r.SetBody(strings.NewReader("{}"))

	// r.SetDest(nil)

	// // Enviando a requisição.
	// resp, err := r.Send()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(string(body))

	api := lapi.NewRequest("https://jsonplaceholder.typicode.com",
		map[string]string{
			"Content-Type": "application/json",
		},
		10,
	)

	// api.SetAuth("access_token", "refresh_token")

	var dest struct {
		ID     int    `json:"id"`
		Title  string `json:"title"`
		UserID int    `json:"userId"`
	}

	err := api.Post("/todos/1", nil, &dest)
	if err != nil {
		log.Fatal(err)
	}
}
