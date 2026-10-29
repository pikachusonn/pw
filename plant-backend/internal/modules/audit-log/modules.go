package auditlog

import (
	"plant-backend/internal/modules/audit-log/adapter"
	"plant-backend/internal/modules/audit-log/application"
	auditLogHttp "plant-backend/internal/modules/audit-log/transport/http"

	"github.com/gin-gonic/gin"
)

type Module struct {
	handler *auditLogHttp.AuditLogHandler
}

func NewModule() *Module {
	repo := adapter.NewAditLogRepository()
	useCase := application.NewAuditLogHandler(repo)
	handler := auditLogHttp.NewAuditLogHandler(useCase)

	return &Module{handler: handler}
}
func (m *Module) RegisterRoutes(router *gin.RouterGroup) {
	auditLogHttp.RegisterRoutes(router, m.handler)
}
