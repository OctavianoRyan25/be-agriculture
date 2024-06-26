package admin

import (
	"testing"

	"github.com/OctavianoRyan25/be-agriculture/constants"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mocking the repository
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) IsDuplicateEmail(email string) (bool, error) {
	args := m.Called(email)
	return args.Bool(0), args.Error(1)
}

func (m *MockRepository) RegisterUser(user *Admin) (*Admin, error) {
	args := m.Called(user)
	return args.Get(0).(*Admin), args.Error(1)
}

func (m *MockRepository) Login(user *Admin) (*Admin, error) {
	args := m.Called(user)
	return args.Get(0).(*Admin), args.Error(1)
}

func (m *MockRepository) GetUserProfile(id uint) (*Admin, error) {
	args := m.Called(id)
	return args.Get(0).(*Admin), args.Error(1)
}

func TestRegisterUser(t *testing.T) {
	mockRepo := new(MockRepository)
	useCase := NewUseCase(mockRepo)

	mockAdmin := &Admin{
		Email: "test@example.com",
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("IsDuplicateEmail", mockAdmin.Email).Return(false, nil)
		mockRepo.On("RegisterUser", mock.AnythingOfType("*admin.Admin")).Return(mockAdmin, nil)

		admin, code, err := useCase.RegisterUser(mockAdmin)

		assert.NoError(t, err)
		assert.Equal(t, constants.CodeSuccess, code)
		assert.Equal(t, mockAdmin, admin)
		mockRepo.AssertExpectations(t)
	})

}

func TestCheckEmail(t *testing.T) {
	mockRepo := new(MockRepository)
	useCase := NewUseCase(mockRepo)

	email := "test@example.com"

	t.Run("success", func(t *testing.T) {
		mockRepo.On("IsDuplicateEmail", email).Return(false, nil)

		code, err := useCase.CheckEmail(email)

		assert.NoError(t, err)
		assert.Equal(t, constants.CodeSuccess, code)
		mockRepo.AssertExpectations(t)
	})

}

func TestLogin(t *testing.T) {
	mockRepo := new(MockRepository)
	useCase := NewUseCase(mockRepo)

	mockAdmin := &Admin{
		Email:    "test@example.com",
		Password: "password",
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("Login", mockAdmin).Return(mockAdmin, nil)

		admin, code, err := useCase.Login(mockAdmin)

		assert.NoError(t, err)
		assert.Equal(t, constants.CodeSuccess, code)
		assert.Equal(t, mockAdmin, admin)
		mockRepo.AssertExpectations(t)
	})

}
