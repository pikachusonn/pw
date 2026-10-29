package plant

import (
	"plant-backend/internal/modules/plant/adapter"
	"plant-backend/internal/modules/plant/application"
	wateringhttp "plant-backend/internal/modules/plant/transport/http"

	"github.com/gin-gonic/gin"
)

type Module struct {
	handler *wateringhttp.Handler
}

func NewModule() *Module {
	repo := adapter.NewPlantRepo()
	useCase := application.NewListPlantHandler(repo)
	handler := wateringhttp.NewHandler(useCase)

	return &Module{handler: handler}
}
func (m *Module) RegisterRoutes(router *gin.RouterGroup) {
	wateringhttp.RegisterRoutes(router, m.handler)
}
