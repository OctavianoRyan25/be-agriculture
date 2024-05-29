package handler

import (
	"net/http"
	"strconv"

	"github.com/OctavianoRyan25/be-agriculture/modules/plant"
	"github.com/OctavianoRyan25/be-agriculture/utils/helper"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UserPlantHandler struct {
	service plant.UserPlantService
}

func NewUserPlantHandler(service plant.UserPlantService) *UserPlantHandler {
	return &UserPlantHandler{service}
}

func (h *UserPlantHandler) AddUserPlant(c echo.Context) error {
	var request plant.AddUserPlantInput
	if err := c.Bind(&request); err != nil {
		response := helper.APIResponse("Invalid request", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.APIResponse(errors[0], http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	userPlant, err := h.service.AddUserPlant(request)
	if err != nil {
		response := helper.APIResponse("Failed to add plant to user", http.StatusInternalServerError, "error", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := helper.APIResponse("Plant added to user successfully", http.StatusCreated, "success", userPlant)
	return c.JSON(http.StatusCreated, response)
}

func (h *UserPlantHandler) GetUserPlants(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		response := helper.APIResponse("Invalid user ID", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	userPlantResponses, err := h.service.GetUserPlantsByUserID(userID)
	if err != nil {
		response := helper.APIResponse("Failed to fetch user plants", http.StatusInternalServerError, "error", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	var plants []plant.UserPlantResponse
	for _, userPlants := range userPlantResponses {
		for _, userPlant := range userPlants {
			plants = append(plants, userPlant)
		}
	}

	responseData := struct {
		UserID int                              	`json:"user_id"`
		Plants []plant.UserPlantResponse `json:"plants"`
	}{
		UserID: userID,
		Plants: plants,
	}

	response := helper.APIResponse("User plants fetched successfully", http.StatusOK, "success", responseData)
	return c.JSON(http.StatusOK, response)
}

func (h *UserPlantHandler) DeleteUserPlantByID(c echo.Context) error {
	userPlantID, err := strconv.Atoi(c.Param("user_plant_id"))
	if err != nil {
			response := helper.APIResponse("Invalid user plant ID", http.StatusBadRequest, "error", nil)
			return c.JSON(http.StatusBadRequest, response)
	}

	deletedUserPlant, err := h.service.DeleteUserPlantByID(userPlantID)
	if err != nil {
			response := helper.APIResponse("Failed to delete user plant", http.StatusInternalServerError, "error", nil)
			return c.JSON(http.StatusInternalServerError, response)
	}

	response := helper.APIResponse("User plant deleted successfully", http.StatusOK, "success", deletedUserPlant)
	return c.JSON(http.StatusOK, response)
}






// func (h *UserPlantHandler) RemoveUserPlant(c echo.Context) error {
// 	var request plant.RemoveUserPlantRequest
// 	if err := c.Bind(&request); err != nil {
// 		response := helper.APIResponse("Invalid request", http.StatusBadRequest, "error", nil)
// 		return c.JSON(http.StatusBadRequest, response)
// 	}

// 	validate := validator.New()
// 	if err := validate.Struct(request); err != nil {
// 		errors := helper.FormatValidationError(err)
// 		response := helper.APIResponse(errors[0], http.StatusBadRequest, "error", nil)
// 		return c.JSON(http.StatusBadRequest, response)
// 	}

// 	if err := h.service.RemoveUserPlant(request); err != nil {
// 		response := helper.APIResponse("Failed to remove plant from user", http.StatusInternalServerError, "error", nil)
// 		return c.JSON(http.StatusInternalServerError, response)
// 	}

// 	response := helper.APIResponse("Plant removed from user successfully", http.StatusOK, "success", nil)
// 	return c.JSON(http.StatusOK, response)
// }
