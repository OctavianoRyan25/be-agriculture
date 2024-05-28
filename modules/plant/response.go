package plant

import "time"

type PlantResponse struct {
	ID               			int                        		`json:"id"`
	Name             			string                     		`json:"name"`
	Description      			string                     		`json:"description"`
	IsToxic          			bool                       		`json:"is_toxic"`
	HarvestDuration  			int                        		`json:"harvest_duration"`
	PlantCategory    			PlantCategoryClimateResponse  `json:"plant_category"`
	ClimateCondition 			string											  `json:"climate_condition"`
	PlantingTime     			string                        `json:"planting_time"`
	Sunlight         			string                        `json:"sunlight"`
	PlantCharacteristic 		PlantCharacteristicResponse 		`json:"plant_characteristic"`
	WateringSchedule 			PlantReminderResponse      		`json:"watering_schedule"`
	PlantInstruction 			[]PlantInstructionResponse 		`json:"plant_instructions"`
	PlantFAQ         			[]PlantFAQResponse         		`json:"plant_faqs"`
	PlantImages      			[]PlantImageResponse       		`json:"plant_images"`
	CreatedAt        			time.Time                  		`json:"created_at"`
}

type PlantCategoryClimateResponse struct {
	ID   									int    				`json:"id"`
	Name 									string 				`json:"name"`
	ImageURL 							string 				`json:"image_url"`
}

type PlantCharacteristicResponse struct {
	ID         						int    				`json:"id"`
	Height     						int    				`json:"height"`
	HeightUnit 						string 				`json:"height_unit"`
	Wide       						int    				`json:"wide"`
	WideUnit   						string 				`json:"wide_unit"`
	LeafColor  						string 				`json:"leaf_color"`
}

type PlantImageResponse struct {
	ID        						int    				`json:"id"`
	PlantID   						int    				`json:"plant_id"`
	FileName  						string 				`json:"file_name"`
	IsPrimary 						int    				`json:"is_primary"`
}

type PlantReminderResponse struct {
	ID                    int       		`json:"id"`
	PlantID               int       		`json:"plant_id"`
	WateringFrequency     int       		`json:"watering_frequency"`
	Each                  string    		`json:"each"`
	WateringAmount        int       		`json:"watering_amount"`
	Unit                  string    		`json:"unit"`
	WateringTime          string    		`json:"watering_time"`
	WeatherCondition      string    		`json:"weather_condition"`
	ConditionDescription  string    		`json:"condition_description"`
}

type PlantInstructionResponse struct {
	ID               			int       		`json:"id"`
	PlantID          			int       		`json:"plant_id"`
	StepNumber       			int       		`json:"step_number"`
	StepTitle        			string    		`json:"step_title"`
	StepDescription  			string    		`json:"step_description"`
	StepImageURL     			string    		`json:"step_image_url"`
	AdditionalTips   			string    		`json:"additional_tips"`
}

type PlantFAQResponse struct {
	ID         						int       		`json:"id"`
	PlantID    						int       		`json:"plant_id"`
	Question   						string    		`json:"question"`
	Answer     						string    		`json:"answer"`
	CreatedAt  						time.Time 		`json:"created_at"`
}

func NewPlantResponse(plant Plant) PlantResponse {
	return PlantResponse{
		ID								 : plant.ID,
		Name							 : plant.Name,
		Description				 : plant.Description,
		IsToxic						 : plant.IsToxic,
		HarvestDuration		 : plant.HarvestDuration,
		ClimateCondition	 : plant.ClimateCondition,
		PlantingTime			 : plant.PlantingTime,
		Sunlight					 : plant.Sunlight,
		PlantCategory			 : NewPlantCategoryResponse(plant.PlantCategory),
		PlantCharacteristic : NewPlantCharacteristicResponse(plant.PlantCharacteristic),
		WateringSchedule	 : NewPlantReminderResponse(plant.WateringSchedule),
		PlantInstruction	 : NewPlantInstructionResponses(plant.PlantInstructions),
		PlantFAQ					 : NewPlantFAQResponses(plant.PlantFAQs),
		PlantImages				 : NewPlantImageResponses(plant.PlantImages),
		CreatedAt					 : plant.CreatedAt,
	}
}

func NewPlantCategoryResponse(category PlantCategory) PlantCategoryClimateResponse {
	return PlantCategoryClimateResponse{
		ID			 : category.ID,
		Name		 : category.Name,
		ImageURL : category.ImageURL,
	}
}

func NewPlantCharacteristicResponse(characteristic PlantCharacteristic) PlantCharacteristicResponse {
	return PlantCharacteristicResponse{
		ID				 : characteristic.ID,
		Height		 : characteristic.Height,
		HeightUnit : characteristic.HeightUnit,
		Wide			 : characteristic.Wide,
		WideUnit	 : characteristic.WideUnit,
		LeafColor	 : characteristic.LeafColor,
	}
}

func NewPlantImageResponses(images []PlantImage) []PlantImageResponse {
	var responses []PlantImageResponse

	for _, img := range images {
		responses = append(responses, NewPlantImageResponse(img))
	}

	return responses
}

func NewPlantImageResponse(image PlantImage) PlantImageResponse {
	return PlantImageResponse{
		ID			  : image.ID,
		PlantID	  : image.PlantID,
		FileName  : image.FileName,
		IsPrimary : image.IsPrimary,
	}
}

func NewPlantReminderResponse(reminder PlantReminder) PlantReminderResponse {
	return PlantReminderResponse{
		ID									 : reminder.ID,
		PlantID							 : reminder.PlantID,
		WateringFrequency		 : reminder.WateringFrequency,
		Each								 : reminder.Each,
		WateringAmount			 : reminder.WateringAmount,
		Unit								 : reminder.Unit,
		WateringTime				 : reminder.WateringTime,
		WeatherCondition		 : reminder.WeatherCondition,
		ConditionDescription : reminder.ConditionDescription,
	}
}

func NewPlantInstructionResponses(instructions []PlantInstruction) []PlantInstructionResponse {
	var responses []PlantInstructionResponse

	for _, instruction := range instructions {
		responses = append(responses, NewPlantInstructionResponse(instruction))
	}

	return responses
}

func NewPlantInstructionResponse(instruction PlantInstruction) PlantInstructionResponse {
	return PlantInstructionResponse{
		ID							: instruction.ID,
		PlantID					: instruction.PlantID,
		StepNumber			: instruction.StepNumber,
		StepTitle				: instruction.StepTitle,
		StepDescription	: instruction.StepDescription,
		StepImageURL		: instruction.StepImageURL,
		AdditionalTips	: instruction.AdditionalTips,
	}
}

func NewPlantFAQResponses(faqs []PlantFAQ) []PlantFAQResponse {
	var responses []PlantFAQResponse

	for _, faq := range faqs {
		responses = append(responses, NewPlantFAQResponse(faq))
	}

	return responses
}

func NewPlantFAQResponse(faq PlantFAQ) PlantFAQResponse {
	return PlantFAQResponse{
		ID 				: faq.ID,
		PlantID 	: faq.PlantID,
		Question 	: faq.Question,
		Answer 		: faq.Answer,
		CreatedAt : faq.CreatedAt,
	}
}
