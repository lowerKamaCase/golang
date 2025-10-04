package main

import (
	"fmt"
	"math"
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

	server.ListenAndServe()

}

func random(rw http.ResponseWriter, request *http.Request) {
	randomFrom1To6 := math.Ceil(rand.Float64() * 6)
	result := fmt.Sprintf("%.0f", randomFrom1To6)

	rw.Write([]byte(result))
}
