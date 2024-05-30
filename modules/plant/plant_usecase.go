package plant

import (
	"time"
)

type PlantService interface {
	FindAll() ([]PlantResponse, error)
	FindByID(id int) (PlantResponse, error)
	CreatePlant(input CreatePlantInput) (PlantResponse, error)
	UpdatePlant(id int, input UpdatePlantInput) (PlantResponse, error)
	DeletePlant(id int) (PlantResponse, error)
}

type plantService struct {
	repository             PlantRepository
	plantCategoryRepository PlantCategoryRepository
}

func NewPlantService(repository PlantRepository, plantCategoryRepository PlantCategoryRepository) PlantService {
	return &plantService{repository, plantCategoryRepository}
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

	plant := Plant{
		Name:               input.Name,
		Description:        input.Description,
		IsToxic:            input.IsToxic,
		HarvestDuration:    input.HarvestDuration,
		Sunlight: input.Sunlight,
		PlantingTime: input.PlantingTime,
		PlantCategoryID:    input.PlantCategoryID,
		PlantCategory:      category, 
		ClimateCondition:   input.ClimateCondition,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
		PlantInstructions:  make([]PlantInstruction, len(input.PlantInstructions)),
		PlantFAQs:          make([]PlantFAQ, len(input.PlantFAQs)),
		PlantImages:        make([]PlantImage, len(input.PlantImages)),
		PlantCharacteristic: PlantCharacteristic{
			Height:     input.PlantCharacteristic.Height,
			HeightUnit: input.PlantCharacteristic.HeightUnit,
			Wide:       input.PlantCharacteristic.Wide,
			WideUnit:   input.PlantCharacteristic.WideUnit,
			LeafColor:  input.PlantCharacteristic.LeafColor,
		},
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

func (s *plantService) UpdatePlant(id int, input UpdatePlantInput) (PlantResponse, error) {
	plant, err := s.repository.FindByID(id)
	if err != nil {
		return PlantResponse{}, err
	}

	category, err := s.plantCategoryRepository.FindByID(input.PlantCategoryID)
	if err != nil {
		return PlantResponse{}, err
	}

	plant.Name = input.Name
	plant.Description = input.Description
	plant.IsToxic = input.IsToxic
	plant.HarvestDuration = input.HarvestDuration
	plant.Sunlight = input.Sunlight
	plant.PlantingTime = input.PlantingTime
	plant.PlantCategoryID = input.PlantCategoryID
	plant.PlantCategory = category
	plant.ClimateCondition = input.ClimateCondition
	plant.UpdatedAt = time.Now()
	
	plant.PlantCharacteristic = PlantCharacteristic{
		Height:     input.PlantCharacteristic.Height,
		HeightUnit: input.PlantCharacteristic.HeightUnit,
		Wide:       input.PlantCharacteristic.Wide,
		WideUnit:   input.PlantCharacteristic.WideUnit,
		LeafColor:  input.PlantCharacteristic.LeafColor,
	}
	
	plant.WateringSchedule = PlantReminder{
		WateringFrequency:    input.WateringSchedule.WateringFrequency,
		Each:                 input.WateringSchedule.Each,
		WateringAmount:       input.WateringSchedule.WateringAmount,
		Unit:                 input.WateringSchedule.Unit,
		WateringTime:         input.WateringSchedule.WateringTime,
		WeatherCondition:     input.WateringSchedule.WeatherCondition,
		ConditionDescription: input.WateringSchedule.ConditionDescription,
	}

	for i, instruction := range input.PlantInstructions {
		plant.PlantInstructions[i] = PlantInstruction{
			StepNumber:      instruction.StepNumber,
			StepTitle:       instruction.StepTitle,
			StepDescription: instruction.StepDescription,
			StepImageURL:    instruction.StepImageURL,
			AdditionalTips:  instruction.AdditionalTips,
			UpdatedAt:       time.Now(),
		}
	}

	for i, faq := range input.PlantFAQs {
		plant.PlantFAQs[i] = PlantFAQ{
			Question:  faq.Question,
			Answer:    faq.Answer,
			UpdatedAt: time.Now(),
		}
	}

	for i, image := range input.PlantImages {
		plant.PlantImages[i] = PlantImage{
			FileName:  image.FileName,
			IsPrimary: image.IsPrimary,
			UpdatedAt: time.Now(),
		}
	}

	updatedPlant, err := s.repository.Update(plant)
	if err != nil {
		return PlantResponse{}, err
	}

	return NewPlantResponse(updatedPlant), nil
}

func (s *plantService) DeletePlant(id int) (PlantResponse, error) {
	plant, err := s.repository.FindByIDWithRelations(id)
	if err != nil {
			return PlantResponse{}, err
	}

	if err := s.repository.Delete(id); err != nil {
			return PlantResponse{}, err
	}

	return NewPlantResponse(plant), nil
}


