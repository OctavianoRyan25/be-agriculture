package handler

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/OctavianoRyan25/be-agriculture/modules/plant"
	"github.com/OctavianoRyan25/be-agriculture/utils/helper"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type PlantHandler struct {
	service plant.PlantService
	cloudinary  *cloudinary.Cloudinary
}

func NewPlantHandler(service plant.PlantService, 	cloudinary  *cloudinary.Cloudinary) *PlantHandler {
	return &PlantHandler{service, cloudinary}
}

func (h *PlantHandler) GetAll(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page <= 0 {
			page = 0 // Default value
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit <= 0 {
			limit = 0 // Default value
	}

	if page > 0 && limit > 0 {
			totalCount, err := h.service.CountAll()
			if err != nil {
					response := helper.APIResponse("Failed to count plants", http.StatusInternalServerError, "error", nil)
					return c.JSON(http.StatusInternalServerError, response)
			}

			if int64((page-1)*limit) >= totalCount {
					response := helper.APIResponse("Page exceeds available data", http.StatusBadRequest, "error", nil)
					return c.JSON(http.StatusBadRequest, response)
			}
	}

	plants, err := h.service.FindAll(page, limit)
	if err != nil {
			response := helper.APIResponse("Failed to fetch plants", http.StatusInternalServerError, "error", nil)
			return c.JSON(http.StatusInternalServerError, response)
	}

	responseData := struct {
			Plants      []plant.PlantResponse `json:"plants"`
			Limit       int                   `json:"limit"`
			Page        int                   `json:"page"`
			TotalCount  int64                 `json:"total_count,omitempty"`
			TotalPages  int                   `json:"total_pages,omitempty"`
	}{
			Plants:     plants,
			Limit:      limit,
			Page:       page,
	}

	if page > 0 && limit > 0 {
			totalCount, err := h.service.CountAll()
			if err == nil {
					responseData.TotalCount = totalCount
					responseData.TotalPages = int((totalCount + int64(limit) - 1) / int64(limit))
			}
	}

	response := helper.APIResponse("Plants fetched successfully", http.StatusOK, "success", responseData)
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
	form, err := c.MultipartForm()
	if err != nil {
		response := helper.APIResponse("Invalid multipart form data", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	var input plant.CreatePlantInput

	input.Name = form.Value["name"][0]
	input.Description = form.Value["description"][0]
	input.IsToxic = form.Value["is_toxic"][0] == "true"
	input.HarvestDuration, _ = strconv.Atoi(form.Value["harvest_duration"][0])
	input.Sunlight = form.Value["sunlight"][0]
	input.PlantingTime = form.Value["planting_time"][0]
	input.PlantCategoryID, _ = strconv.Atoi(form.Value["plant_category_id"][0])
	input.ClimateCondition = form.Value["climate_condition"][0]

	// Parsing plant characteristic
	input.PlantCharacteristic = plant.CreatePlantCharacteristicInput{
		Height: atoi(form.Value["plant_characteristic.height"][0]),
		HeightUnit: form.Value["plant_characteristic.height_unit"][0],
		Wide: atoi(form.Value["plant_characteristic.wide"][0]),
		WideUnit: form.Value["plant_characteristic.wide_unit"][0],
		LeafColor: form.Value["plant_characteristic.leaf_color"][0],
	}

	// Parsing watering schedule
	input.WateringSchedule = plant.CreateWateringScheduleInput{
		WateringFrequency:    atoi(form.Value["watering_schedule.watering_frequency"][0]),
		Each:                 form.Value["watering_schedule.each"][0],
		WateringAmount:       atoi(form.Value["watering_schedule.watering_amount"][0]),
		Unit:                 form.Value["watering_schedule.unit"][0],
		WateringTime:         form.Value["watering_schedule.watering_time"][0],
		WeatherCondition:     form.Value["watering_schedule.weather_condition"][0],
		ConditionDescription: form.Value["watering_schedule.condition_description"][0],
	}

	// Parsing plant instructions with image upload
	instructionFiles := form.File["plant_instructions.step_image_url"]
	for i := 0; i < len(form.Value["plant_instructions.step_number"]); i++ {
		instruction := plant.CreatePlantInstructionInput{
			InstructionCategoryID: atoi(form.Value["plant_instructions.instruction_category_id"][i]),
			StepNumber:      atoi(form.Value["plant_instructions.step_number"][i]),
			StepTitle:       form.Value["plant_instructions.step_title"][i],
			StepDescription: form.Value["plant_instructions.step_description"][i],
			StepImageURL:    "",
			AdditionalTips:  form.Value["plant_instructions.additional_tips"][i],
		}

		// Handle file upload for step_image_url
		if i < len(instructionFiles) {
			file := instructionFiles[i]
			src, err := file.Open()
			if err != nil {
				return err
			}
			defer src.Close()

			uploadResult, err := h.cloudinary.Upload.Upload(context.Background(), src, uploader.UploadParams{Folder: "be-agriculture"})
			if err != nil {
				return err
			}

			instruction.StepImageURL = uploadResult.SecureURL
		}

		input.PlantInstructions = append(input.PlantInstructions, instruction)
	}

	// Parsing plant FAQs
	for i := 0; i < len(form.Value["plant_faqs.question"]); i++ {
		faq := plant.CreatePlantFAQInput{
			Question: form.Value["plant_faqs.question"][i],
			Answer:   form.Value["plant_faqs.answer"][i],
		}
		input.PlantFAQs = append(input.PlantFAQs, faq)
	}

	// Parsing plant images
	files := form.File["plant_images"]
	for i, file := range files {
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		uploadResult, err := h.cloudinary.Upload.Upload(context.Background(), src, uploader.UploadParams{Folder: "be-agriculture"})
		if err != nil {
			return err
		}

		plantImage := plant.CreatePlantImageInput{
			FileName:  uploadResult.SecureURL,
			IsPrimary: atoi(form.Value["plant_images.is_primary"][i]),
		}
		input.PlantImages = append(input.PlantImages, plantImage)
	}

	// Perform validation
	validate := validator.New()
	if err := validate.Struct(&input); err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.APIResponse(strings.Join(errors, ", "), http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	createdPlant, err := h.service.CreatePlant(input)
	if err != nil {
		response := helper.APIResponse("Failed to create plant", http.StatusInternalServerError, "error", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := helper.APIResponse("Plant created successfully", http.StatusCreated, "success", createdPlant)
	return c.JSON(http.StatusCreated, response)
}

func atoi(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

func (h *PlantHandler) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := helper.APIResponse("Invalid plant ID", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	form, err := c.MultipartForm()
	if err != nil {
		response := helper.APIResponse("Invalid multipart form data", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	var input plant.UpdatePlantInput

	input.Name = form.Value["name"][0]
	input.Description = form.Value["description"][0]
	input.IsToxic = form.Value["is_toxic"][0] == "true"
	input.HarvestDuration, _ = strconv.Atoi(form.Value["harvest_duration"][0])
	input.Sunlight = form.Value["sunlight"][0]
	input.PlantingTime = form.Value["planting_time"][0]
	input.PlantCategoryID, _ = strconv.Atoi(form.Value["plant_category_id"][0])
	input.ClimateCondition = form.Value["climate_condition"][0]

	// Parsing plant characteristic
	input.PlantCharacteristic = plant.CreatePlantCharacteristicInput{
		Height: atoi(form.Value["plant_characteristic.height"][0]),
		HeightUnit: form.Value["plant_characteristic.height_unit"][0],
		Wide: atoi(form.Value["plant_characteristic.wide"][0]),
		WideUnit: form.Value["plant_characteristic.wide_unit"][0],
		LeafColor: form.Value["plant_characteristic.leaf_color"][0],
	}

	// Parsing watering schedule
	input.WateringSchedule = plant.CreateWateringScheduleInput{
		WateringFrequency:    atoi(form.Value["watering_schedule.watering_frequency"][0]),
		Each:                 form.Value["watering_schedule.each"][0],
		WateringAmount:       atoi(form.Value["watering_schedule.watering_amount"][0]),
		Unit:                 form.Value["watering_schedule.unit"][0],
		WateringTime:         form.Value["watering_schedule.watering_time"][0],
		WeatherCondition:     form.Value["watering_schedule.weather_condition"][0],
		ConditionDescription: form.Value["watering_schedule.condition_description"][0],
	}

	// Parsing plant instructions with image upload
	instructionFiles := form.File["plant_instructions.step_image_url"]
	for i := 0; i < len(form.Value["plant_instructions.step_number"]); i++ {
		instruction := plant.CreatePlantInstructionInput{
			InstructionCategoryID: atoi(form.Value["plant_instructions.instruction_category_id"][i]),
			StepNumber:      atoi(form.Value["plant_instructions.step_number"][i]),
			StepTitle:       form.Value["plant_instructions.step_title"][i],
			StepDescription: form.Value["plant_instructions.step_description"][i],
			StepImageURL:    "",
			AdditionalTips:  form.Value["plant_instructions.additional_tips"][i],
		}

		// Handle file upload for step_image_url
		if i < len(instructionFiles) {
			file := instructionFiles[i]
			src, err := file.Open()
			if err != nil {
				return err
			}
			defer src.Close()

			uploadResult, err := h.cloudinary.Upload.Upload(context.Background(), src, uploader.UploadParams{Folder: "be-agriculture"})
			if err != nil {
				return err
			}

			instruction.StepImageURL = uploadResult.SecureURL
		}

		input.PlantInstructions = append(input.PlantInstructions, instruction)
	}

	// Parsing plant FAQs
	for i := 0; i < len(form.Value["plant_faqs.question"]); i++ {
		faq := plant.CreatePlantFAQInput{
			Question: form.Value["plant_faqs.question"][i],
			Answer:   form.Value["plant_faqs.answer"][i],
		}
		input.PlantFAQs = append(input.PlantFAQs, faq)
	}

	// Parsing plant images
	files := form.File["plant_images"]
	for i, file := range files {
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		uploadResult, err := h.cloudinary.Upload.Upload(context.Background(), src, uploader.UploadParams{Folder: "be-agriculture"})
		if err != nil {
			return err
		}

		plantImage := plant.CreatePlantImageInput{
			FileName:  uploadResult.SecureURL,
			IsPrimary: atoi(form.Value["plant_images.is_primary"][i]),
		}
		
		input.PlantImages = append(input.PlantImages, plantImage)
	}

	// Perform validation
	validate := validator.New()
	if err := validate.Struct(&input); err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.APIResponse(strings.Join(errors, ", "), http.StatusBadRequest, "error", nil)
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

func (h *PlantHandler) SearchPlantsByName(c echo.Context) error {
	name := c.QueryParam("name")
	if name == "" {
		response := helper.APIResponse("Name parameter is required", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit <= 0 {
		limit = 10
	}

	plants, totalCount, err := h.service.SearchPlantsByName(name, page, limit)
	if err != nil {
		response := helper.APIResponse("Failed to search plants", http.StatusInternalServerError, "error", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	if totalCount == 0 {
		response := helper.APIResponse("No plants found with the given name", http.StatusNotFound, "error", nil)
		return c.JSON(http.StatusNotFound, response)
	}

	responseData := struct {
		Plants     []plant.Plant `json:"plants"`
		TotalCount int64         `json:"total_count"`
		Limit      int           `json:"limit"`
		Page       int           `json:"page"`
		TotalPages int           `json:"total_pages"`
	}{
		Plants:     plants,
		TotalCount: totalCount,
		Limit:      limit,
		Page:       page,
		TotalPages: int((totalCount + int64(limit) - 1) / int64(limit)),
	}

	response := helper.APIResponse("Plants fetched successfully", http.StatusOK, "success", responseData)
	return c.JSON(http.StatusOK, response)
}


