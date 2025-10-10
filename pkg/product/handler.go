package product

import (
	"fmt"
	"lowerkamacase/golang/configs"
	"lowerkamacase/golang/pkg/req"
	"lowerkamacase/golang/pkg/res"
	"net/http"
)

type PostProductDeps struct {
	*configs.Config
}

type PostProduct struct {
	*configs.Config
}

func NewProductHandler(router *http.ServeMux, deps PostProductDeps) {
	handler := &PostProduct{
		Config: deps.Config,
	}
	router.HandleFunc("POST /product", handler.CreateProduct())

}

func (*PostProduct) CreateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {

		body, err := req.HandleBody[CreateProduct](&w, request)
		if err !=nil {
			res.Json(w, err.Error(), 400)
			return
		}

		res.Json(w, fmt.Sprintf("Successful creation of product %s", body.Name), 200)
	}
}