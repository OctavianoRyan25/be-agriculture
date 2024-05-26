package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/OctavianoRyan25/be-agriculture/modules/plant"
	"github.com/OctavianoRyan25/be-agriculture/utils/helper"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type PlantHandler struct {
	service plant.PlantService
}

func NewPlantHandler(service plant.PlantService) *PlantHandler {
	return &PlantHandler{service}
}

func (h *PlantHandler) GetAll(c echo.Context) error {
	plants, err := h.service.FindAll()
	if err != nil {
		response := helper.APIResponse("Failed to fetch plants", http.StatusInternalServerError, "error", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	response := helper.APIResponse("Plants fetched successfully", http.StatusOK, "success", plants)
	return c.JSON(http.StatusOK, response)
}

func (h *PlantHandler) GetByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := helper.APIResponse("Invalid ID", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	plant, err := h.service.FindByID(id)
	if err != nil {
		response := helper.APIResponse("Plant not found", http.StatusNotFound, "error", nil)
		return c.JSON(http.StatusNotFound, response)
	}
	response := helper.APIResponse("Plant fetched successfully", http.StatusOK, "success", plant)
	return c.JSON(http.StatusOK, response)
}

func (h *PlantHandler) Create(c echo.Context) error {
	var request plant.CreatePlantInput
	if err := c.Bind(&request); err != nil {
		response := helper.APIResponse("Invalid request", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	validate := validator.New()
	if err := validate.Struct(&request); err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.APIResponse(strings.Join(errors, ", "), http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	createdPlant, err := h.service.CreatePlant(request)
	if err != nil {
		response := helper.APIResponse("Failed to create plant", http.StatusInternalServerError, "error", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := helper.APIResponse("Plant created successfully", http.StatusCreated, "success", createdPlant)
	return c.JSON(http.StatusCreated, response)
}

func (h *PlantHandler) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
			response := helper.APIResponse("Invalid ID", http.StatusBadRequest, "error", nil)
			return c.JSON(http.StatusBadRequest, response)
	}

	var input plant.CreatePlantInput
	if err := c.Bind(&input); err != nil {
			response := helper.APIResponse("Failed to bind input", http.StatusBadRequest, "error", nil)
			return c.JSON(http.StatusBadRequest, response)
	}

	validate := validator.New()
	if err := validate.Struct(&input); err != nil {
			errors := helper.FormatValidationError(err)
			response := helper.APIResponse("Validation error", http.StatusBadRequest, "error", errors)
			return c.JSON(http.StatusBadRequest, response)
	}

	updatedPlant, err := h.service.UpdatePlant(id, input)
	if err != nil {
			response := helper.APIResponse("Failed to update plant", http.StatusInternalServerError, "error", nil)
			return c.JSON(http.StatusInternalServerError, response)
	}

	response := helper.APIResponse("Plant updated successfully", http.StatusOK, "success", updatedPlant)
	return c.JSON(http.StatusOK, response)
}

func (h *PlantHandler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
			response := helper.APIResponse("Invalid ID", http.StatusBadRequest, "error", nil)
			return c.JSON(http.StatusBadRequest, response)
	}

	deletedPlant, err := h.service.DeletePlant(id)
	if err != nil {
			response := helper.APIResponse("Failed to delete plant", http.StatusInternalServerError, "error", nil)
			return c.JSON(http.StatusInternalServerError, response)
	}

	response := helper.APIResponse("Plant deleted successfully", http.StatusOK, "success", deletedPlant)
	return c.JSON(http.StatusOK, response)
}
