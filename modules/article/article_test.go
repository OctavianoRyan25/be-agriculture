package article

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type mockRepository struct {
	storeArticleFunc   func(*Article) (*Article, error)
	getArticleFunc     func(int) (*Article, error)
	getAllArticlesFunc func() ([]Article, error)
	updateArticleFunc  func(*Article, int) (*Article, error)
	deleteArticleFunc  func(int) error
}

func (m *mockRepository) StoreArticle(a *Article) (*Article, error) {
	return m.storeArticleFunc(a)
}

func (m *mockRepository) GetArticle(id int) (*Article, error) {
	return m.getArticleFunc(id)
}

func (m *mockRepository) GetAllArticles() ([]Article, error) {
	return m.getAllArticlesFunc()
}

func (m *mockRepository) UpdateArticle(a *Article, id int) (*Article, error) {
	return m.updateArticleFunc(a, id)
}

func (m *mockRepository) DeleteArticle(id int) error {
	return m.deleteArticleFunc(id)
}

func TestStoreArticle_Success(t *testing.T) {
	mockRepo := &mockRepository{
		storeArticleFunc: func(a *Article) (*Article, error) {
			a.ID = 1
			a.CreatedAt = time.Now()
			a.UpdatedAt = time.Now()
			return a, nil
		},
	}

	uc := NewUseCase(mockRepo)
	newArticle := &Article{
		Title:   "Test Title",
		Content: "Test Content",
	}

	storedArticle, err := uc.StoreArticle(newArticle)
	assert.NoError(t, err)
	assert.NotNil(t, storedArticle)
	assert.Equal(t, newArticle.Title, storedArticle.Title)
	assert.Equal(t, newArticle.Content, storedArticle.Content)
	assert.NotZero(t, storedArticle.ID)
}

func TestStoreArticle_Error(t *testing.T) {
	expectedError := errors.New("store error")
	mockRepo := &mockRepository{
		storeArticleFunc: func(a *Article) (*Article, error) {
			return nil, expectedError
		},
	}

	uc := NewUseCase(mockRepo)
	newArticle := &Article{
		Title:   "Test Title",
		Content: "Test Content",
	}

	storedArticle, err := uc.StoreArticle(newArticle)
	assert.Error(t, err)
	assert.Nil(t, storedArticle)
	assert.Equal(t, expectedError, err)
}

func TestGetArticle_Success(t *testing.T) {
	mockRepo := &mockRepository{
		getArticleFunc: func(id int) (*Article, error) {
			return &Article{
				ID:      id,
				Title:   "Test Title",
				Content: "Test Content",
			}, nil
		},
	}

	uc := NewUseCase(mockRepo)
	articleID := 1

	article, err := uc.GetArticle(articleID)
	assert.NoError(t, err)
	assert.NotNil(t, article)
	assert.Equal(t, articleID, article.ID)
}

func TestGetArticle_Error(t *testing.T) {
	expectedError := errors.New("not found")
	mockRepo := &mockRepository{
		getArticleFunc: func(id int) (*Article, error) {
			return nil, expectedError
		},
	}

	uc := NewUseCase(mockRepo)
	articleID := 1

	article, err := uc.GetArticle(articleID)
	assert.Error(t, err)
	assert.Nil(t, article)
	assert.Equal(t, expectedError, err)
}

func TestGetAllArticles_Success(t *testing.T) {
	mockRepo := &mockRepository{
		getAllArticlesFunc: func() ([]Article, error) {
			return []Article{
				{ID: 1, Title: "Test Title 1", Content: "Test Content 1"},
				{ID: 2, Title: "Test Title 2", Content: "Test Content 2"},
			}, nil
		},
	}

	uc := NewUseCase(mockRepo)

	articles, err := uc.GetAllArticles()
	assert.NoError(t, err)
	assert.Len(t, articles, 2)
}

func TestGetAllArticles_Error(t *testing.T) {
	expectedError := errors.New("no articles found")
	mockRepo := &mockRepository{
		getAllArticlesFunc: func() ([]Article, error) {
			return nil, expectedError
		},
	}

	uc := NewUseCase(mockRepo)

	articles, err := uc.GetAllArticles()
	assert.Error(t, err)
	assert.Nil(t, articles)
	assert.Equal(t, expectedError, err)
}

func TestUpdateArticle_Success(t *testing.T) {
	mockRepo := &mockRepository{
		updateArticleFunc: func(a *Article, id int) (*Article, error) {
			a.ID = id
			a.UpdatedAt = time.Now()
			return a, nil
		},
	}

	uc := NewUseCase(mockRepo)
	updateData := &Article{
		Title:   "Updated Title",
		Content: "Updated Content",
	}
	articleID := 1

	updatedArticle, err := uc.UpdateArticle(updateData, articleID)
	assert.NoError(t, err)
	assert.NotNil(t, updatedArticle)
	assert.Equal(t, articleID, updatedArticle.ID)
	assert.Equal(t, updateData.Title, updatedArticle.Title)
	assert.Equal(t, updateData.Content, updatedArticle.Content)
}

func TestUpdateArticle_Error(t *testing.T) {
	expectedError := errors.New("update error")
	mockRepo := &mockRepository{
		updateArticleFunc: func(a *Article, id int) (*Article, error) {
			return nil, expectedError
		},
	}

	uc := NewUseCase(mockRepo)
	updateData := &Article{
		Title:   "Updated Title",
		Content: "Updated Content",
	}
	articleID := 1

	updatedArticle, err := uc.UpdateArticle(updateData, articleID)
	assert.Error(t, err)
	assert.Nil(t, updatedArticle)
	assert.Equal(t, expectedError, err)
}

func TestDeleteArticle_Success(t *testing.T) {
	mockRepo := &mockRepository{
		deleteArticleFunc: func(id int) error {
			return nil
		},
	}

	uc := NewUseCase(mockRepo)
	articleID := 1

	err := uc.DeleteArticle(articleID)
	assert.NoError(t, err)
}

func TestDeleteArticle_Error(t *testing.T) {
	expectedError := errors.New("delete error")
	mockRepo := &mockRepository{
		deleteArticleFunc: func(id int) error {
			return expectedError
		},
	}

	uc := NewUseCase(mockRepo)
	articleID := 1

	err := uc.DeleteArticle(articleID)
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
}
