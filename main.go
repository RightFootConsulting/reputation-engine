package main

import (
	"net/http"
	logger "rightfoot-consulting/reputation-engine/pkg/logging"
	"rightfoot-consulting/reputation-engine/pkg/routes"
	"rightfoot-consulting/reputation-engine/pkg/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env.server")
	if err != nil {
		panic(err)
	}
	sm, err := service.GetServiceManager()
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	routeManager, err := routes.Create(r, sm)
	if err != nil {
		panic(err)
	}
	if routeManager != nil {
		logger.Println("Created all routes ready to serve.")
	}
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
