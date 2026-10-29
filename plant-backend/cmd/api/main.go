package main

import (
	auditLog "plant-backend/internal/modules/audit-log"
	"plant-backend/internal/modules/plant"
	"plant-backend/internal/platform/httpserver"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	plantModule := plant.NewModule()
	auditLogModule := auditLog.NewModule()
	apiV1 := router.Group("/api/v1")
	{
		apiV1.GET("/health", func(c *gin.Context) {
			httpserver.OK(c, gin.H{
				"message": "Server is running",
			})
		})
		plantModule.RegisterRoutes(apiV1)
		auditLogModule.RegisterRoutes(apiV1)
	}

	router.Run(":5773")
}
