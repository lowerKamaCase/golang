package main

import (
	"fmt"
	"lowerkamacase/golang/internal/random"
	"net/http"
)

const PORT = 8087

func main() {
	serveMux := http.NewServeMux()

	serveMux.HandleFunc("/random", random.Random)

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
