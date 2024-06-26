package fertilizer

type FertilizerInput struct {
	Name         string `json:"name"`
	PlantID      int    `json:"plantId"`
	Compostition string `json:"compostition"`
}

func NewFertilizerInput(fertilizer FertilizerInput) *Fertilizer {
	return &Fertilizer{
		Name:         fertilizer.Name,
		PlantID:      fertilizer.PlantID,
		Compostition: fertilizer.Compostition,
	}
}
