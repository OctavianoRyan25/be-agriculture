// handler.go
package handler

import (
	"net/http"

	"github.com/OctavianoRyan25/be-agriculture/modules/plant"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/labstack/echo/v4"
)

type PlantEarliestWateringHandler struct {
	service plant.PlantEarliestWateringService
	cloudinary  *cloudinary.Cloudinary
}

func NewPlantEarliestWateringHandler(service plant.PlantEarliestWateringService, cloudinary  *cloudinary.Cloudinary) *PlantEarliestWateringHandler {
	return &PlantEarliestWateringHandler{service , cloudinary}
}

func (h *PlantEarliestWateringHandler) GetEarliestWateringTime(c echo.Context) error {
    plantID := c.Param("plant_id")
    if plantID == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "plant_id is required"})
    }

    schedule, err := h.service.FindEarliestWateringTime(plantID)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }

    return c.JSON(http.StatusOK, schedule)
}
