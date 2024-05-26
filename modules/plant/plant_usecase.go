package plant

import (
	"time"
)

type PlantService interface {
	FindAll() ([]PlantResponse, error)
	FindByID(id int) (PlantResponse, error)
	CreatePlant(input CreatePlantInput) (PlantResponse, error)
	UpdatePlant(id int, input CreatePlantInput) (PlantResponse, error)
	DeletePlant(id int) (PlantResponse, error)
}

type plantService struct {
	repository             PlantRepository
	plantCategoryRepository PlantCategoryRepository
	climateConditionRepository ClimateConditionRepository
}

func NewPlantService(repository PlantRepository, plantCategoryRepository PlantCategoryRepository, climateConditionRepository ClimateConditionRepository) PlantService {
	return &plantService{repository, plantCategoryRepository, climateConditionRepository}
}

func (s *plantService) FindAll() ([]PlantResponse, error) {
	plants, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	var responses []PlantResponse
	for _, plant := range plants {
		responses = append(responses, NewPlantResponse(plant))
	}

	return responses, nil
}

func (s *plantService) FindByID(id int) (PlantResponse, error) {
	plant, err := s.repository.FindByID(id)
	if err != nil {
		return PlantResponse{}, err
	}

	return NewPlantResponse(plant), nil
}

func (s *plantService) CreatePlant(input CreatePlantInput) (PlantResponse, error) {
	category, err := s.plantCategoryRepository.FindByID(input.PlantCategoryID)
	if err != nil {
		return PlantResponse{}, err
	}

	condition, err := s.climateConditionRepository.FindByID(input.ClimateConditionID)
	if err != nil {
		return PlantResponse{}, err
	}

	plant := Plant{
		Name:               input.Name,
		Description:        input.Description,
		IsToxic:            input.IsToxic,
		HarvestDuration:    input.HarvestDuration,
		PlantCategoryID:    input.PlantCategoryID,
		PlantCategory:      category, 
		ClimateConditionID: input.ClimateConditionID,
		ClimateCondition:   condition,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
		PlantInstructions:  make([]PlantInstruction, len(input.PlantInstructions)),
		PlantFAQs:          make([]PlantFAQ, len(input.PlantFAQs)),
		PlantImages:        make([]PlantImage, len(input.PlantImages)),
		WateringSchedule: PlantReminder{
			WateringFrequency:   input.WateringSchedule.WateringFrequency,
			Each:                input.WateringSchedule.Each,
			WateringAmount:      input.WateringSchedule.WateringAmount,
			Unit:                input.WateringSchedule.Unit,
			WateringTime:        input.WateringSchedule.WateringTime,
			WeatherCondition:    input.WateringSchedule.WeatherCondition,
			ConditionDescription: input.WateringSchedule.ConditionDescription,
		},
	}

	for i, instruction := range input.PlantInstructions {
		plant.PlantInstructions[i] = PlantInstruction{
			StepNumber:      instruction.StepNumber,
			StepTitle:       instruction.StepTitle,
			StepDescription: instruction.StepDescription,
			StepImageURL:    instruction.StepImageURL,
			AdditionalTips:  instruction.AdditionalTips,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		}
	}

	for i, faq := range input.PlantFAQs {
		plant.PlantFAQs[i] = PlantFAQ{
			Question:  faq.Question,
			Answer:    faq.Answer,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
	}

	for i, image := range input.PlantImages {
		plant.PlantImages[i] = PlantImage{
			FileName:  image.FileName,
			IsPrimary: image.IsPrimary,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
	}

	newPlant, err := s.repository.Create(plant)
	if err != nil {
		return PlantResponse{}, err
	}

	return NewPlantResponse(newPlant), nil
}

func (s *plantService) UpdatePlant(id int, input CreatePlantInput) (PlantResponse, error) {
	plant, err := s.repository.FindByID(id)
	if err != nil {
		return PlantResponse{}, err
	}

	plant.Name = input.Name
	plant.Description = input.Description
	plant.IsToxic = input.IsToxic
	plant.HarvestDuration = input.HarvestDuration
	plant.PlantCategoryID = input.PlantCategoryID
	plant.ClimateConditionID = input.ClimateConditionID
	plant.UpdatedAt = time.Now()


	updatedPlant, err := s.repository.Update(plant)
	if err != nil {
		return PlantResponse{}, err
	}

	return NewPlantResponse(updatedPlant), nil
}

func (s *plantService) DeletePlant(id int) (PlantResponse, error) {
	plant, err := s.repository.Delete(id)
	if err != nil {
		return PlantResponse{}, err
	}

	return NewPlantResponse(plant), nil
}
