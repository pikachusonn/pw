package application

import (
	"context"
	"plant-backend/internal/modules/plant/domain"
)

// The Port: Application says "I need someone who can fetch all plants"
type PlantRepository interface {
	ListAll(ctx context.Context) ([]*domain.Plant, error)
}
