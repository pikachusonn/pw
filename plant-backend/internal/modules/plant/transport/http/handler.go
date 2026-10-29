package http

import (
	"net/http"
	"plant-backend/internal/modules/plant/application"
	"plant-backend/internal/platform/httpserver"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	listPlants *application.ListPlantsHandler
}

func NewHandler(listPlants *application.ListPlantsHandler) *Handler {
	return &Handler{
		listPlants: listPlants,
	}
}

func (h *Handler) List(c *gin.Context) {
	// 1. Call the application use case (passing the request's context)
	plants, err := h.listPlants.Execute(c.Request.Context())
	if err != nil {
		httpserver.WriteError(c, "erros.internal_eror", http.StatusInternalServerError, err.Error())
		return
	}
	httpserver.OK(c, plants)
}
