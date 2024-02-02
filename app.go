package main

import (
	"fmt"
	"io"
	"node/database"
	"node/handlers"
	"node/middlewares"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("config.env")

	GIN_MODE := os.Getenv("GIN_MODE")
	GIN_ADDR := os.Getenv("GIN_ADDR")

	f, _ := os.Create("logs.txt")

	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	db, err := database.NewDb("clients.sql")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer db.Close()

	gin.SetMode(GIN_MODE)

	serv := gin.New()
	serv.Use(gin.Recovery(), gin.Logger(), middlewares.ApiTokenAuthMiddleware(db))

	serv.GET("/ping", handlers.PingHandler)
	serv.GET("/getClient", handlers.GetClientHandler(db))
	serv.GET("/getClients", handlers.GetClientsHandler(db))
	serv.GET("/addClient", handlers.AddClient(db))
	serv.GET("/updateExpiryDate", handlers.UpdateExpiryDateHandler(db))

	serv.Run(GIN_ADDR)
}
