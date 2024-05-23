package plant

type PlantService interface {
	FindAll() ([]PlantResponse, error)
	FindByID(id int) (PlantResponse, error)
	CreatePlant(request CreatePlantInput) (PlantResponse, error)
	UpdatePlant(id int, request CreatePlantInput) (PlantResponse, error)
	DeletePlant(id int) (PlantResponse, error)
}

type plantService struct {
	repository PlantRepository
}

func NewPlantService(repository PlantRepository) PlantService {
	return &plantService{repository}
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

func (s *plantService) CreatePlant(request CreatePlantInput) (PlantResponse, error) {
	createdPlant, err := s.repository.Create(request)
	if err != nil {
		return PlantResponse{}, err
	}
	return NewPlantResponse(createdPlant), nil
}

func (s *plantService) UpdatePlant(id int, request CreatePlantInput) (PlantResponse, error) {
	updatedPlant, err := s.repository.Update(id, request)
	if err != nil {
		return PlantResponse{}, err
	}
	return NewPlantResponse(updatedPlant), nil
}

func (s *plantService) DeletePlant(id int) (PlantResponse, error) {
	deletedPlant, err := s.repository.Delete(id)
	if err != nil {
		return PlantResponse{}, err
	}
	return NewPlantResponse(deletedPlant), nil
}
