package application

import (
	"context"
	"plant-backend/internal/modules/plant/domain"
)

type ListPlantsHandler struct {
	repo PlantRepository
}

func NewListPlantHandler(repo PlantRepository) *ListPlantsHandler {
	return &ListPlantsHandler{repo: repo}
}

func (h *ListPlantsHandler) Execute(ctx context.Context) ([]*domain.Plant, error) {
	return h.repo.ListAll(ctx)
}
