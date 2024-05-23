package handler

import (
	"net/http"
	"strconv"

	"github.com/OctavianoRyan25/be-agriculture/modules/plant"
	"github.com/OctavianoRyan25/be-agriculture/utils/helper"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ClimateConditionHandler struct {
	service plant.ClimateConditionService
}

func NewClimateConditionHandler(service plant.ClimateConditionService) *ClimateConditionHandler {
	return &ClimateConditionHandler{service}
}

func (h *ClimateConditionHandler) GetAll(c echo.Context) error {
	climateConditions, err := h.service.FindAll()

	if err != nil {
		response := helper.APIResponse("Failed to get climate conditions", http.StatusInternalServerError, "error", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	
	response := helper.APIResponse("Climate conditions fetched successfully", http.StatusOK, "success", climateConditions)
	return c.JSON(http.StatusOK, response)
}

func (h *ClimateConditionHandler) GetByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		response := helper.APIResponse("Invalid ID", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	climateCondition, err := h.service.FindByID(id)
	if err != nil {
		response := helper.APIResponse("Climate condition not found", http.StatusNotFound, "error", nil)
		return c.JSON(http.StatusNotFound, response)
	}

	response := helper.APIResponse("Climate condition fetched successfully", http.StatusOK, "success", climateCondition)
	return c.JSON(http.StatusOK, response)
}

func (h *ClimateConditionHandler) Create(c echo.Context) error {
	var input plant.PlantCategoryClimateInput

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

	climateCondition, err := h.service.Create(input)
	if err != nil {
		response := helper.APIResponse("Failed to create climate condition", http.StatusInternalServerError, "error", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := helper.APIResponse("Climate condition created successfully", http.StatusCreated, "success", climateCondition)
	return c.JSON(http.StatusCreated, response)
}

func (h *ClimateConditionHandler) Update(c echo.Context) error {
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
	if err := validate.Struct(&input); err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.APIResponse("Validation error", http.StatusBadRequest, "error", errors)
		return c.JSON(http.StatusBadRequest, response)
	}

	climateCondition, err := h.service.Update(id, input)
	if err != nil {
		response := helper.APIResponse("Failed to update climate condition", http.StatusInternalServerError, "error", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := helper.APIResponse("Climate condition updated successfully", http.StatusOK, "success", climateCondition)
	return c.JSON(http.StatusOK, response)
}

func (h *ClimateConditionHandler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		response := helper.APIResponse("Invalid ID", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	condition, err := h.service.Delete(id)
	if err != nil {
		response := helper.APIResponse("Failed to delete climate condition", http.StatusInternalServerError, "error", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := helper.APIResponse("Climate condition deleted successfully", http.StatusOK, "success", condition)
	return c.JSON(http.StatusOK, response)
}