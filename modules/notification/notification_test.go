package notification

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type mockRepository struct {
	storeNotificationFunc               func(*Notification) (*Notification, error)
	readNotificationFunc                func(int) (*Notification, error)
	getAllNotificationsFunc             func(uint) ([]Notification, error)
	deleteAllNotificationsFunc          func(uint) error
	createCustomizeWateringReminderFunc func(*CustomizeWateringReminder) (*CustomizeWateringReminder, error)
}

func (m *mockRepository) StoreNotification(n *Notification) (*Notification, error) {
	return m.storeNotificationFunc(n)
}

func (m *mockRepository) ReadNotification(id int) (*Notification, error) {
	return m.readNotificationFunc(id)
}

func (m *mockRepository) GetAllNotifications(userID uint) ([]Notification, error) {
	return m.getAllNotificationsFunc(userID)
}

func (m *mockRepository) DeleteAllNotifications(userID uint) error {
	return m.deleteAllNotificationsFunc(userID)
}

func (m *mockRepository) CreateCustomizeWateringReminder(r *CustomizeWateringReminder) (*CustomizeWateringReminder, error) {
	return m.createCustomizeWateringReminderFunc(r)
}

func TestStoreNotification_Success(t *testing.T) {
	mockRepo := &mockRepository{
		storeNotificationFunc: func(n *Notification) (*Notification, error) {
			n.Id = 1
			n.CreatedAt = time.Now()
			n.UpdatedAt = time.Now()
			return n, nil
		},
	}

	uc := NewUseCase(mockRepo)
	newNotification := &Notification{
		UserId: 1,
		Body:   "Test Notification",
	}

	storedNotification, err := uc.StoreNotification(newNotification)
	assert.NoError(t, err)
	assert.NotNil(t, storedNotification)
	assert.Equal(t, newNotification.Body, storedNotification.Body)
	assert.NotZero(t, storedNotification.Id)
}

func TestStoreNotification_Error(t *testing.T) {
	expectedError := errors.New("store error")
	mockRepo := &mockRepository{
		storeNotificationFunc: func(n *Notification) (*Notification, error) {
			return nil, expectedError
		},
	}

	uc := NewUseCase(mockRepo)
	newNotification := &Notification{
		UserId: 1,
		Body:   "Test Notification",
	}

	storedNotification, err := uc.StoreNotification(newNotification)
	assert.Error(t, err)
	assert.Nil(t, storedNotification)
	assert.Equal(t, expectedError, err)
}

func TestReadNotification_Success(t *testing.T) {
	mockRepo := &mockRepository{
		readNotificationFunc: func(id int) (*Notification, error) {
			return &Notification{
				Id:     id,
				UserId: 1,
				Body:   "Test Notification",
			}, nil
		},
	}

	uc := NewUseCase(mockRepo)
	notificationID := 1

	readNotification, err := uc.ReadNotification(notificationID)
	assert.NoError(t, err)
	assert.NotNil(t, readNotification)
	assert.Equal(t, notificationID, readNotification.Id)
}

func TestReadNotification_Error(t *testing.T) {
	expectedError := errors.New("not found")
	mockRepo := &mockRepository{
		readNotificationFunc: func(id int) (*Notification, error) {
			return nil, expectedError
		},
	}

	uc := NewUseCase(mockRepo)
	notificationID := 1

	readNotification, err := uc.ReadNotification(notificationID)
	assert.Error(t, err)
	assert.Nil(t, readNotification)
	assert.Equal(t, expectedError, err)
}

func TestGetAllNotifications_Success(t *testing.T) {
	mockRepo := &mockRepository{
		getAllNotificationsFunc: func(userID uint) ([]Notification, error) {
			return []Notification{
				{Id: 1, UserId: int(userID), Body: "Test Notification 1"},
				{Id: 2, UserId: int(userID), Body: "Test Notification 2"},
			}, nil
		},
	}

	uc := NewUseCase(mockRepo)
	userID := uint(1)

	notifications, err := uc.GetAllNotifications(userID)
	assert.NoError(t, err)
	assert.Len(t, notifications, 2)
}

func TestGetAllNotifications_Error(t *testing.T) {
	expectedError := errors.New("no notifications found")
	mockRepo := &mockRepository{
		getAllNotificationsFunc: func(userID uint) ([]Notification, error) {
			return nil, expectedError
		},
	}

	uc := NewUseCase(mockRepo)
	userID := uint(1)

	notifications, err := uc.GetAllNotifications(userID)
	assert.Error(t, err)
	assert.Nil(t, notifications)
	assert.Equal(t, expectedError, err)
}

func TestCreateCustomizeWateringReminder_Success(t *testing.T) {
	mockRepo := &mockRepository{
		createCustomizeWateringReminderFunc: func(r *CustomizeWateringReminder) (*CustomizeWateringReminder, error) {
			r.Id = 1
			r.CreatedAt = time.Now()
			r.UpdatedAt = time.Now()
			return r, nil
		},
	}

	uc := NewUseCase(mockRepo)
	newReminder := &CustomizeWateringReminder{
		UserId: 1,
		Type:   "Watering Reminder",
	}

	storedReminder, err := uc.CreateCustomizeWateringReminder(newReminder)
	assert.NoError(t, err)
	assert.NotNil(t, storedReminder)
	assert.Equal(t, newReminder.Type, storedReminder.Type)
	assert.NotZero(t, storedReminder.Id)
}

func TestCreateCustomizeWateringReminder_Error(t *testing.T) {
	expectedError := errors.New("store error")
	mockRepo := &mockRepository{
		createCustomizeWateringReminderFunc: func(r *CustomizeWateringReminder) (*CustomizeWateringReminder, error) {
			return nil, expectedError
		},
	}

	uc := NewUseCase(mockRepo)
	newReminder := &CustomizeWateringReminder{
		UserId: 1,
		Type:   "Watering Reminder",
	}

	storedReminder, err := uc.CreateCustomizeWateringReminder(newReminder)
	assert.Error(t, err)
	assert.Nil(t, storedReminder)
	assert.Equal(t, expectedError, err)
}
