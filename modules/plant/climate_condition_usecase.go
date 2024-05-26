package plant

type ClimateConditionService interface {
	FindAll() ([]PlantCategoryClimateResponse, error)
	FindByID(id int) (PlantCategoryClimateResponse, error)
	Create(input PlantCategoryClimateInput) (PlantCategoryClimateResponse, error)
	Update(id int, input PlantCategoryClimateInput) (PlantCategoryClimateResponse, error)
	Delete(id int) (PlantCategoryClimateResponse, error)
}

type climateConditionService struct {
	repository ClimateConditionRepository
}

func NewClimateConditionService(repository ClimateConditionRepository) ClimateConditionService {
	return &climateConditionService{repository}
}

func (s *climateConditionService) FindAll() ([]PlantCategoryClimateResponse, error) {
	conditions, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	var response []PlantCategoryClimateResponse
	for _, condition := range conditions {
		response = append(response, NewClimateConditionResponse(condition))
	}

	return response, nil
}

func (s *climateConditionService) FindByID(id int) (PlantCategoryClimateResponse, error) {
	condition, err := s.repository.FindByID(id)
	if err != nil {
		return PlantCategoryClimateResponse{}, err
	}

	return NewClimateConditionResponse(condition), nil
}

func (s *climateConditionService) Create(input PlantCategoryClimateInput) (PlantCategoryClimateResponse, error) {
	condition := ClimateCondition{
		Name:     input.Name,
		ImageURL: input.ImageURL,
	}

	newCondition, err := s.repository.Create(condition)
	if err != nil {
		return PlantCategoryClimateResponse{}, err
	}

	return NewClimateConditionResponse(newCondition), nil
}

func (s *climateConditionService) Update(id int, input PlantCategoryClimateInput) (PlantCategoryClimateResponse, error) {
	condition, err := s.repository.FindByID(id)
	if err != nil {
		return PlantCategoryClimateResponse{}, err
	}

	condition.Name = input.Name
	condition.ImageURL = input.ImageURL
	updatedCondition, err := s.repository.Update(condition)
	if err != nil {
		return PlantCategoryClimateResponse{}, err
	}

	return NewClimateConditionResponse(updatedCondition), nil
}

func (s *climateConditionService) Delete(id int) (PlantCategoryClimateResponse, error) {
	condition, err := s.repository.FindByID(id)
	if err != nil {
		return PlantCategoryClimateResponse{}, err
	}

	err = s.repository.Delete(condition)
	if err != nil {
		return PlantCategoryClimateResponse{}, err
	}

	return NewClimateConditionResponse(condition), nil
}
