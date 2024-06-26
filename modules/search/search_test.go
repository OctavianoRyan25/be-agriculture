package search_test

import (
	"testing"

	"github.com/OctavianoRyan25/be-agriculture/modules/plant"
	"github.com/OctavianoRyan25/be-agriculture/modules/search"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Search(params search.PlantSearchParams) ([]plant.Plant, error) {
	args := m.Called(params)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]plant.Plant), args.Error(1)
}

func TestSearch(t *testing.T) {
	mockRepo := new(MockRepository)
	usecase := search.NewUsecase(mockRepo)

	mockPlants := []plant.Plant{
		{
			ID:                    1,
			Name:                  "Plant 1",
			Description:           "Description 1",
			IsToxic:               false,
			HarvestDuration:       30,
			Sunlight:              "Full Sun",
			PlantingTime:          "Spring",
			PlantCategoryID:       1,
			ClimateCondition:      "Tropical",
			PlantCharacteristicID: 1,
			AdditionalTips:        "Water daily",
		},
	}

	params := search.PlantSearchParams{
		Name:            "Plant",
		PlantCategory:   "Vegetable",
		DifficultyLevel: "Easy",
		Sunlight:        "Full Sun",
		HarvestDuration: "30-60 days",
		IsToxic:         nil,
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("Search", params).Return(mockPlants, nil)

		plants, err := usecase.Search(params)

		assert.NoError(t, err)
		assert.NotNil(t, plants)
		assert.Equal(t, mockPlants, plants)
		mockRepo.AssertExpectations(t)
	})

	t.Run("no results", func(t *testing.T) {
		mockRepo.On("Search", params).Return([]plant.Plant{}, nil)

		_, err := usecase.Search(params)

		assert.NoError(t, err)
		assert.Nil(t, nil)
		mockRepo.AssertExpectations(t)
	})
}
