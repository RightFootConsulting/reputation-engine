package routes

import (
	"net/http"
	"rightfoot-consulting/reputation-engine/pkg/service"

	"github.com/gin-gonic/gin"
)

type RouteManager struct {
	serviceManager *service.ServiceManager
	router         *gin.Engine
}

func Create(serviceManage *service.ServiceManager) (result *RouteManager, err error) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
