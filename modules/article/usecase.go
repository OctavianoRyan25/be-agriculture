package article

type UseCase interface {
	StoreArticle(*Article) (*Article, error)
	GetArticle(int) (*Article, error)
	GetAllArticles() ([]Article, error)
	UpdateArticle(*Article) (*Article, error)
	DeleteArticle(int) error
}

type useCase struct {
	repo Repository
}

func NewUseCase(repo Repository) *useCase {
	return &useCase{
		repo: repo,
	}
}

func (uc *useCase) StoreArticle(a *Article) (*Article, error) {
	return uc.repo.StoreArticle(a)
}

func (uc *useCase) GetArticle(id int) (*Article, error) {
	return uc.repo.GetArticle(id)
}

func (uc *useCase) GetAllArticles() ([]Article, error) {
	return uc.repo.GetAllArticles()
}

func (uc *useCase) UpdateArticle(a *Article) (*Article, error) {
	return uc.repo.UpdateArticle(a)
}

func (uc *useCase) DeleteArticle(id int) error {
	return uc.repo.DeleteArticle(id)
}
