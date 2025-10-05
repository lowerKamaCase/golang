package verify

import (
	"fmt"
	"lowerkamacase/golang/configs"
	"net/http"
)

type VerifierDeps struct {
	*configs.Config
}

type Verifier struct {
	*configs.Config
}

func NewVerifierHandler(router *http.ServeMux, deps VerifierDeps) {
	handler := &Verifier{
		Config: deps.Config,
	}

	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("GET /verify/{hash}", handler.Verify())

}

func (verifier *Verifier) Send() http.HandlerFunc {
	return func(rw http.ResponseWriter, request *http.Request) {
		fmt.Println("Send")
	}
}

func (verifier *Verifier) Verify() http.HandlerFunc {
	return func(rw http.ResponseWriter, request *http.Request) {
		fmt.Println("Verify")
	}
}
