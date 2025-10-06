package main

import (
	"fmt"
	"lowerkamacase/golang/configs"
	"lowerkamacase/golang/internal/auth"
	"lowerkamacase/golang/internal/verify"
	"net/http"
)

const PORT = 8087

func main() {
	conf := configs.LoadConfig()
	fmt.Println(conf)

	if conf == nil {
		fmt.Print("Config cannot be nil")
		panic("Config cannot be nil")
	}

	serveMux := http.NewServeMux()

	auth.NewAuthHandler(serveMux, auth.AuthHandlerDeps{
		Config: conf,
	})

	verify.NewVerifierHandler(serveMux, verify.VerifierDeps{
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
