package main

import (
	"chilley.nam2507/config"
	"chilley.nam2507/routes"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	router := gin.Default()
	cfg, err := config.LoadAllAppConfig()
	if err != nil {
		log.Fatal(err)
		return
	}
	corsConfig := cors.Config{
		AllowOrigins:     []string{cfg.CorsProd, cfg.CorsDev},
		AllowMethods:     config.AllowMethods,
		AllowHeaders:     config.AllowHeaders,
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	router.Use(cors.New(corsConfig))

	//Load routes
	routes.RegisterRoutes(router)

	serverAddr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	log.Println("Server is listening on " + serverAddr)
	if err := router.Run(serverAddr); err != nil {
		log.Fatal("Failed to start application")
		return
	}
}
