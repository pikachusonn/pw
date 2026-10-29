package adapter

import (
	"context"
	"plant-backend/internal/modules/audit-log/domain"
)

type AuditlopRepository struct {
	auditLogs []*domain.AuditLog
}

func NewAditLogRepository() *AuditlopRepository {
	return &AuditlopRepository{
		auditLogs: []*domain.AuditLog{
			domain.NewAuditLog("watering.water", "Tưới cây chậu A thành công, tưới 250ml nước", "1"),
		},
	}
}

func (r *AuditlopRepository) ListAll(ctx context.Context) ([]*domain.AuditLog, error) {
	return r.auditLogs, nil
}
