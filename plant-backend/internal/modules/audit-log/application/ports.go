package application

import (
	"context"
	"plant-backend/internal/modules/audit-log/domain"
)

type AuditlogRepository interface {
	ListAll(ctx context.Context) ([]*domain.AuditLog, error)
}
