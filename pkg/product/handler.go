package product

import (
	"lowerkamacase/golang/pkg/req"
	"lowerkamacase/golang/pkg/res"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type ProductHandlerDeps struct {
	ProductRepository *ProductRepository
}

type ProductHandler struct {
	ProductRepository *ProductRepository
}

func NewProductHandler(router *http.ServeMux, deps ProductHandlerDeps) {
	handler := &ProductHandler{
		ProductRepository: deps.ProductRepository,
	}
	router.HandleFunc("POST /product", handler.Create())
	router.HandleFunc("PUT /product/{id}", handler.Update())
	router.HandleFunc("DELETE /product/{id}", handler.Delete())
	router.HandleFunc("GET /product/{id}", handler.GetById())

}

func (handler *ProductHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		body, err := req.HandleBody[ProductCreateRequest](&w, request)
		if err != nil {
			return
		}

		product := NewProduct(body.Name)

		createdProduct, err := handler.ProductRepository.Create(product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res.Json(w, createdProduct, http.StatusCreated)
	}
}

func (handler *ProductHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		body, err := req.HandleBody[ProductUpdateRequest](&w, request)
		if err != nil {
			return
		}

		idString := request.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		link, err := handler.ProductRepository.Update(&Product{
			Model: gorm.Model{ID: uint(id)},
			Name:  body.Name,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Json(w, link, http.StatusAccepted)
	}
}

func (handler *ProductHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		idString := request.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		_, err = handler.ProductRepository.GetById(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		err = handler.ProductRepository.Delete(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Json(w, nil, http.StatusOK)
	}
}

func (handler *ProductHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		idString := request.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		product, err := handler.ProductRepository.GetById(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		res.Json(w, product, http.StatusOK)
	}
}
