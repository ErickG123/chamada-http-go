package main

import (
	"io"
	"net/http"
)

func main() {
	// Fazendo uma chamada HTTP
	request, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}

	// Eu preciso finalizar a stream dos dados
	// O defer é statement que faz atrasar a execução de alguma coisa
	// O defer é chamado por último no nosso código
	defer request.Body.Close()

	// O ReadAll faz a leitura do corpo da requisição
	result, err := io.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}

	// Preciso converter o resultado para String
	println(string(result))
}
