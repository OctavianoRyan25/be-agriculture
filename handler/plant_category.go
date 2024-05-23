package handler

import (
	"net/http"
	"strconv"

	"github.com/OctavianoRyan25/be-agriculture/modules/plant"
	"github.com/OctavianoRyan25/be-agriculture/utils/helper"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type PlantCategoryHandler struct {
	service plant.PlantCategoryService
}

func NewPlantCategoryHandler(service plant.PlantCategoryService) *PlantCategoryHandler {
	return &PlantCategoryHandler{service}
}

func (h *PlantCategoryHandler) GetAll(c echo.Context) error {
	categories, err := h.service.FindAll()
	
	if err != nil {
		response := helper.APIResponse("Failed to get plant categories", http.StatusInternalServerError, "error", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := helper.APIResponse("Plant categories fetched successfully", http.StatusOK, "success", categories)
	return c.JSON(http.StatusOK, response)
}

func (h *PlantCategoryHandler) GetByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		response := helper.APIResponse("Invalid ID", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	category, err := h.service.FindByID(id)
	if err != nil {
		response := helper.APIResponse("Plant category not found", http.StatusNotFound, "error", nil)
		return c.JSON(http.StatusNotFound, response)
	}

	response := helper.APIResponse("Plant category fetched successfully", http.StatusOK, "success", category)
	return c.JSON(http.StatusOK, response)
}

func (h *PlantCategoryHandler) Create(c echo.Context) error {
	var input plant.PlantCategoryClimateInput

	if err := c.Bind(&input); err != nil {
		response := helper.APIResponse("Failed to bind input", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.APIResponse("Validation error", http.StatusBadRequest, "error", errors)
		return c.JSON(http.StatusBadRequest, response)
	}

	category, err := h.service.Create(input)
	if err != nil {
		response := helper.APIResponse("Failed to create plant category", http.StatusInternalServerError, "error", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := helper.APIResponse("Plant category created successfully", http.StatusCreated, "success", category)
	return c.JSON(http.StatusCreated, response)
}

func (h *PlantCategoryHandler) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		response := helper.APIResponse("Invalid ID", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	var input plant.PlantCategoryClimateInput

	if err := c.Bind(&input); err != nil {
		response := helper.APIResponse("Failed to bind input", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.APIResponse("Validation error", http.StatusBadRequest, "error", errors)
		return c.JSON(http.StatusBadRequest, response)
	}

	category, err := h.service.Update(id, input)
	if err != nil {
		response := helper.APIResponse("Failed to update plant category", http.StatusInternalServerError, "error", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := helper.APIResponse("Plant category updated successfully", http.StatusOK, "success", category)
	return c.JSON(http.StatusOK, response)
}

func (h *PlantCategoryHandler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		response := helper.APIResponse("Invalid ID", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	category, err := h.service.FindByID(id)
	if err != nil {
		response := helper.APIResponse("Plant category not found", http.StatusNotFound, "error", nil)
		return c.JSON(http.StatusNotFound, response)
	}

	err = h.service.Delete(id)
	if err != nil {
		response := helper.APIResponse("Failed to delete plant category", http.StatusInternalServerError, "error", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := helper.APIResponse("Plant category deleted successfully", http.StatusOK, "success", category)
	return c.JSON(http.StatusOK, response)
}

