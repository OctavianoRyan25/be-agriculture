package article

import (
	"net/http"
	"strconv"
	"time"

	"github.com/OctavianoRyan25/be-agriculture/base"
	"github.com/OctavianoRyan25/be-agriculture/constants"
	"github.com/labstack/echo/v4"
)

type ArticleController struct {
	useCase UseCase
}

func NewArticleController(useCase UseCase) *ArticleController {
	return &ArticleController{
		useCase: useCase,
	}
}

func (c *ArticleController) StoreArticle(e echo.Context) error {
	role := e.Get("admin").(string)
	if role != "user" {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: "Forbidden access",
			Code:    http.StatusForbidden,
		}
		return e.JSON(http.StatusForbidden, errRes)
	}

	title := e.FormValue("title")
	content := e.FormValue("content")
	image, err := e.FormFile("image")
	if err != nil {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		}
		return e.JSON(http.StatusBadRequest, errRes)
	}

	file, err := image.Open()
	if err != nil {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		}
		return e.JSON(http.StatusBadRequest, errRes)
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	imagePath := "public/" + timestamp

	url, err := UploadToCloudinary(file, imagePath)
	if err != nil {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
		return e.JSON(http.StatusInternalServerError, errRes)
	}

	article := &Article{
		Title:   title,
		Content: content,
		Image:   url,
	}
	article, err = c.useCase.StoreArticle(article)
	if err != nil {
		errResponse := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    constants.ErrCodeBadRequest,
		}
		return e.JSON(constants.ErrCodeBadRequest, errResponse)
	}
	res := ArticleTOResponse(article)
	return e.JSON(constants.Created, res)
}