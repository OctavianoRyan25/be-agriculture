package plant

type UserPlantService interface {
	AddUserPlant(input AddUserPlantInput) (UserPlantResponse, error)
	GetUserPlantsByUserID(userID int) (map[int][]UserPlantResponse, error)
	DeleteUserPlantByID(userPlantID int) (UserPlantResponse, error)
}

type userPlantService struct {
	repository UserPlantRepository
}

func NewUserPlantService(repository UserPlantRepository) UserPlantService {
	return &userPlantService{repository}
}

func (s *userPlantService) AddUserPlant(input AddUserPlantInput) (UserPlantResponse, error) {
	userPlant := UserPlant{
		UserID:  input.UserID,
		PlantID: input.PlantID,
	}

	newUserPlant, err := s.repository.AddUserPlant(userPlant)
	if err != nil {
		return UserPlantResponse{}, err
	}

	return NewUserPlantResponse(newUserPlant), nil
}

func (s *userPlantService) GetUserPlantsByUserID(userID int) (map[int][]UserPlantResponse, error) {
	userPlants, err := s.repository.GetUserPlantsByUserID(userID)
	if err != nil {
		return nil, err
	}

	return NewUserPlantResponses(userPlants), nil
}

func (s *userPlantService) DeleteUserPlantByID(userPlantID int) (UserPlantResponse, error) {
	deletedUserPlant, err := s.repository.GetUserPlantByID(userPlantID)
	if err != nil {
		return UserPlantResponse{}, err
	}

	err = s.repository.DeleteUserPlantByID(userPlantID)
	if err != nil {
		return UserPlantResponse{}, err
	}

	deletedUserPlantResponse := NewUserPlantResponse(deletedUserPlant)

	return deletedUserPlantResponse, nil
}
