package fertilizer

import (
	"net/http"

	"github.com/OctavianoRyan25/be-agriculture/base"
	"github.com/OctavianoRyan25/be-agriculture/utils/helper"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"strconv"
)

type FertilizerController struct {
	useCase FertilizerUseCase
}

func NewFertilizerController(useCase FertilizerUseCase) *FertilizerController {
	return &FertilizerController{
		useCase: useCase,
	}
}

func (c *FertilizerController) CreateFertilizer(ctx echo.Context) error {

	req := new(FertilizerRequest)
	err := ctx.Bind(&req)
	if err != nil {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusUnprocessableEntity,
		}
		return ctx.JSON(http.StatusUnprocessableEntity, errRes)
	}

	validate := validator.New()

	err = validate.Struct(req)
	if err != nil {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		}
		return ctx.JSON(http.StatusBadRequest, errRes)
	}

	mapped := &Fertilizer{
		Id: req.Id,
	}

	f, err := c.useCase.CreateFertilizer(mapped)
	if err != nil {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
		return ctx.JSON(http.StatusInternalServerError, errRes)
	}

	mappedres := &FertilizerResponse{
		Id:           f.Id,
		Name:         f.Name,
		Compostition: f.Compostition,
		CreateAt:     f.CreateAt,
	}

	res := base.SuccessResponse{
		Status:  "success",
		Message: "Watering history created",
		Data:    mappedres,
	}

	return ctx.JSON(http.StatusCreated, res)
}

func (c *FertilizerController) GetFertilizer(ctx echo.Context) error {
	userID := ctx.Get("Id").(uint)

	f, err := c.useCase.GetFertilizer(userID)
	if err != nil {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
		return ctx.JSON(http.StatusInternalServerError, errRes)
	}

	var mappedres []FertilizerResponse
	for _, v := range f {
		mappedres = append(mappedres, FertilizerResponse{
			Id:           v.Id,
			Name:         v.Name,
			Compostition: v.Compostition,
			CreateAt:     v.CreateAt,
		})
	}

	res := base.SuccessResponse{
		Status:  "success",
		Message: "Fertilizer fetched",
		Data:    mappedres,
	}

	return ctx.JSON(http.StatusOK, res)
}

func (c *FertilizerController) GetFertilizerById(ctx echo.Context) error {
	userID := ctx.Get("Id").(uint)

	f, err := c.useCase.GetFertilizer(userID)
	if err != nil {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
		return ctx.JSON(http.StatusInternalServerError, errRes)
	}

	var mappedres []FertilizerResponse
	for _, v := range f {
		mappedres = append(mappedres, FertilizerResponse{
			Id:           v.Id,
			Name:         v.Name,
			Compostition: v.Compostition,
			CreateAt:     v.CreateAt,
		})
	}

	res := base.SuccessResponse{
		Status:  "success",
		Message: "Fertilizer fetched",
		Data:    mappedres,
	}

	return ctx.JSON(http.StatusOK, res)
}

// Update a Fertilizer by Id
func (h *FertilizerController) UpdateFertilizer(c echo.Context) error {
	_, err := strconv.Atoi(c.Param("Id"))

	if err != nil {
		response := helper.APIResponse("Invalid ID", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	role := c.Get("role").(string)
	if role != "admin" {
		response := helper.APIResponse("Only admin can update", http.StatusUnauthorized, "error", nil)
		return c.JSON(http.StatusUnauthorized, response)
	}

	var input Fertilizer

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

	response := helper.APIResponse("Plant category updated successfully", http.StatusOK, "success", nil)
	return c.JSON(http.StatusOK, response)
}

// Delete a Fertilizer by Id
func (c *FertilizerController) DeleteFertilizer(ctx echo.Context) error {
	userId := ctx.Get("user_id").(uint)
	if userId == 0 {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: "Id cannot be empty",
			Code:    http.StatusBadRequest,
		}
		return ctx.JSON(http.StatusBadRequest, errRes)
	}
	err := c.useCase.DeleteFertilizer(userId)
	if err != nil {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
		return ctx.JSON(http.StatusInternalServerError, errRes)
	}
	res := base.SuccessResponse{
		Status:  "success",
		Message: "Fertilizer deleted",
	}
	return ctx.JSON(http.StatusOK, res)
}
