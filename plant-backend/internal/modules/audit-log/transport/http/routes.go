package http

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, handler *AuditLogHandler) {
	router.GET("/audit-log", handler.List)
}
