package plant

type UserPlantService interface {
	AddUserPlant(input AddUserPlantInput) (UserPlantResponse, error)
	GetUserPlantsByUserID(userID int, limit int, offset int) (map[int][]UserPlantResponse, error)
	DeleteUserPlantByID(userPlantID int) (UserPlantResponse, error)
	GetUserPlantByID(userPlantID int) (UserPlant, error)
	CountByUserID(userID int) (int64, error)
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

func (s *userPlantService) GetUserPlantsByUserID(userID int, limit int, page int) (map[int][]UserPlantResponse, error) {
	offset := (page - 1) * limit

	if limit <= 0 {
		limit = -1
		offset = -1
	}

	userPlants, err := s.repository.GetUserPlantsByUserID(userID, limit, offset)
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

func (s *userPlantService) GetUserPlantByID(userPlantID int) (UserPlant, error) {
	return s.repository.GetUserPlantByID(userPlantID)
}

func (s *userPlantService) CountByUserID(userID int) (int64, error) {
	var count int64
	err := s.repository.CountByUserID(userID, &count)
	return count, err
}
