package main

import (
	"fmt"
	"lowerkamacase/golang/configs"
	"lowerkamacase/golang/internal/auth"
	"net/http"
)

const PORT = 8087

func main() {
	conf := configs.LoadConfig()

	serveMux := http.NewServeMux()

	auth.NewAuthHandler(serveMux, auth.AuthHandlerDeps{
		Config: conf,
	})

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
