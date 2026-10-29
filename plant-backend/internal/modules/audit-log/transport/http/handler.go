package http

import (
	"net/http"
	"plant-backend/internal/modules/audit-log/application"
	"plant-backend/internal/platform/httpserver"

	"github.com/gin-gonic/gin"
)

type AuditLogHandler struct {
	listAuditLogHandler *application.ListAuditLogHandler
}

func NewAuditLogHandler(h *application.ListAuditLogHandler) *AuditLogHandler {
	return &AuditLogHandler{
		listAuditLogHandler: h,
	}
}

func (h *AuditLogHandler) List(c *gin.Context) {
	auditLogs, err := h.listAuditLogHandler.Execute(c.Request.Context())
	if err != nil {
		httpserver.WriteError(c, "erros.internal_eror", http.StatusInternalServerError, err.Error())
		return
	}
	httpserver.OK(c, auditLogs)
}
