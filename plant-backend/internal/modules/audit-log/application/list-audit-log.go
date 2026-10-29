package application

import (
	"context"
	"plant-backend/internal/modules/audit-log/domain"
)

type ListAuditLogHandler struct {
	auditLogRepository AuditlogRepository
}

func NewAuditLogHandler(alr AuditlogRepository) *ListAuditLogHandler {
	return &ListAuditLogHandler{
		auditLogRepository: alr,
	}
}

func (h *ListAuditLogHandler) Execute(ctx context.Context) ([]*domain.AuditLog, error) {
	return h.auditLogRepository.ListAll(ctx)
}
