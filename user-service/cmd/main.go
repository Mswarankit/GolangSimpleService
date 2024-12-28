package main

import (
	"flag"
	"fmt"

	"github.com/Mswarankit/user-service/internal/api/handlers"
	"github.com/Mswarankit/user-service/internal/api/middleware"
	"github.com/Mswarankit/user-service/internal/store"
	"github.com/gin-gonic/gin"
)

func main() {
	port := flag.Int("port", 8080, "Server port")
	flag.Parse()

	router := gin.Default()
	userStore := store.NewMemoryStore()
	userHandler := handlers.NewUserHandler(userStore)

	router.Use(middleware.BasicAuth())

	router.POST("/users", userHandler.CreateUser)
	router.POST("/users/:id", userHandler.GetUser)
	router.POST("/users", userHandler.ListUser)

	addr := fmt.Sprintf(":%d", *port)
	router.Run(addr)
}
