package main

import (
	"fmt"
	"lowerkamacase/golang/configs"
	"lowerkamacase/golang/internal/auth"
	"lowerkamacase/golang/pkg/db"
	"lowerkamacase/golang/pkg/link"
	"lowerkamacase/golang/pkg/middleware"
	"lowerkamacase/golang/pkg/product"
	"net/http"
)

const PORT = 8081

func main() {
	conf := configs.LoadConfig()
	database := db.NewDb(conf)
	fmt.Println(conf)

	if conf == nil {
		fmt.Print("Config cannot be nil")
		panic("Config cannot be nil")
	}

	// Repositories
	linkRepository := link.NewLinkRepository(database)
	productRepository := product.NewProductRepository(database)

	serveMux := http.NewServeMux()

	auth.NewAuthHandler(serveMux, auth.AuthHandlerDeps{
		Config: conf,
	})

	link.NewLinkHandler(serveMux, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
	})

	product.NewProductHandler(serveMux, product.ProductHandlerDeps{
		ProductRepository: productRepository,
	})

	Addr := fmt.Sprintf(":%d", PORT)

	stackMiddlewaresFn := middleware.Chain(
		middleware.Logging,
		middleware.CORS,
	)
	server := http.Server{
		Addr: Addr,
		Handler: stackMiddlewaresFn(serveMux),
	}

	fmt.Println("Server started at port: ", PORT)

	err := server.ListenAndServe()

	if err != nil {
		panic(err.Error())
	}

}
