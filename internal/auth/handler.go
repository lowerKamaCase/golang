package auth

import (
	"fmt"
	"lowerkamacase/golang/configs"
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
		fmt.Println("Login Secret: ", handler.Config.Auth.Secret)
		loginResponse := LoginResponse{
			Token: "1111",
		}

		res.Json(rw, loginResponse, 200)
	}

}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(rw http.ResponseWriter, request *http.Request) {
		fmt.Println("Register Dsn: ", handler.Config.Db.Dsn)
	}
}
