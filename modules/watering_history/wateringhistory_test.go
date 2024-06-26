package wateringhistory

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) StoreWateringHistory(wh *WateringHistory) (*WateringHistory, error) {
	args := m.Called(wh)
	return args.Get(0).(*WateringHistory), args.Error(1)
}

func (m *MockRepository) GetAllWateringHistories(userID uint) ([]WateringHistory, error) {
	args := m.Called(userID)
	return args.Get(0).([]WateringHistory), args.Error(1)
}

func (m *MockRepository) GetLateWateringHistories(userID uint) (Notification, error) {
	args := m.Called(userID)
	return args.Get(0).(Notification), args.Error(1)
}

func TestStoreWateringHistory(t *testing.T) {
	mockRepo := new(MockRepository)
	usecase := NewUseCase(mockRepo)

	wh := &WateringHistory{
		PlantID: 1,
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("StoreWateringHistory", wh).Return(wh, nil).Once()

		storedWH, err := usecase.StoreWateringHistory(wh)

		assert.NoError(t, err)
		assert.NotNil(t, storedWH)
		assert.Equal(t, wh, storedWH)
		mockRepo.AssertExpectations(t)
	})

}

func TestGetAllWateringHistories(t *testing.T) {
	mockRepo := new(MockRepository)
	usecase := NewUseCase(mockRepo)

	userID := uint(1)
	wateringHistories := []WateringHistory{
		{PlantID: 1},
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("GetAllWateringHistories", userID).Return(wateringHistories, nil).Once()

		histories, err := usecase.GetAllWateringHistories(userID)

		assert.NoError(t, err)
		assert.NotNil(t, histories)
		assert.Equal(t, wateringHistories, histories)
		mockRepo.AssertExpectations(t)
	})

}

func TestGetLateWateringHistories(t *testing.T) {
	mockRepo := new(MockRepository)
	usecase := NewUseCase(mockRepo)

	userID := uint(1)
	notification := Notification{
		Body: "You have late watering tasks.",
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("GetLateWateringHistories", userID).Return(notification, nil).Once()

		lateNotification, err := usecase.GetLateWateringHistories(userID)

		assert.NoError(t, err)
		assert.NotNil(t, lateNotification)
		assert.Equal(t, notification, lateNotification)
		mockRepo.AssertExpectations(t)
	})

}
