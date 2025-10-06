package auth

import (
	"fmt"
	"lowerkamacase/golang/configs"
	"lowerkamacase/golang/pkg/req"
	"lowerkamacase/golang/pkg/res"
	"net/http"
)

type AuthHandlerDeps struct {
	*configs.Config
}

type AuthHandler struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())

}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(rw http.ResponseWriter, request *http.Request) {
		body, err := req.HandleBody[LoginRequest](&rw, request)
		if err != nil {
			return
		}
		fmt.Println(body)

		loginResponse := LoginResponse{
			Token: "1111",
		}

		res.Json(rw, loginResponse, 200)
	}

}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(rw http.ResponseWriter, request *http.Request) {
		body, err := req.HandleBody[RegisterRequest](&rw, request)
		if err != nil {
			return
		}
		fmt.Println(body)

	}
}
