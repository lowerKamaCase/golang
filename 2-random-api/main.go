package main

import (
	"fmt"
	"math/rand/v2"
	"net/http"
)

const PORT = 8087

func main() {
	serveMux := http.NewServeMux()

	serveMux.HandleFunc("/random", random)

	Addr := fmt.Sprintf(":%d", PORT)

	server := http.Server{
		Addr:    Addr,
		Handler: serveMux,
	}

	fmt.Println("Server started at port: ", PORT)

	err := server.ListenAndServe()

	if err != nil {
		panic(err.Error())
	}

}

func random(rw http.ResponseWriter, request *http.Request) {
	fmt.Println("Got request ", *request)
	randomFrom1To6 := rand.IntN(6) + 1
	result := fmt.Sprintf("%d", randomFrom1To6)

	rw.Write([]byte(result))
}
