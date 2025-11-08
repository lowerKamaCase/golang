package main

import (
	"fmt"
	"lowerkamacase/golang/configs"
	"lowerkamacase/golang/internal/auth"
	"lowerkamacase/golang/internal/user"
	"lowerkamacase/golang/pkg/db"
	"lowerkamacase/golang/pkg/event"
	"lowerkamacase/golang/pkg/link"
	"lowerkamacase/golang/pkg/middleware"
	"lowerkamacase/golang/pkg/stat"
	"net/http"
)

const PORT = 8081

func main() {
	conf := configs.LoadConfig()
	database := db.NewDb(conf)
	fmt.Println(conf)

	eventBus := event.NewEventBus()

	if conf == nil {
		fmt.Print("Config cannot be nil")
		panic("Config cannot be nil")
	}

	// Repositories
	linkRepository := link.NewLinkRepository(database)
	userRepository := user.NewUserRepository(database)
	statRepository := stat.NewStatRepository(database)

	// Services
	authService := auth.NewAuthService(userRepository)
	statService := stat.NewStatService(&stat.StatServiceDeps{
		EventBus:       eventBus,
		StatRepository: statRepository,
	})

	serveMux := http.NewServeMux()

	// Handlers
	auth.NewAuthHandler(serveMux, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})
	link.NewLinkHandler(serveMux, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
		Config:         conf,
		EventBus:       eventBus,
	})
	stat.NewStatHandler(serveMux, stat.StatHandlerDeps{
		StatRepository: statRepository,
		Config:         conf,
	})

	Addr := fmt.Sprintf(":%d", PORT)

	stackMiddlewaresFn := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)
	server := http.Server{
		Addr:    Addr,
		Handler: stackMiddlewaresFn(serveMux),
	}

	fmt.Println("Server started at port: ", PORT)

	err := server.ListenAndServe()

	if err != nil {
		panic(err.Error())
	}

	go statService.AddClick()

}
