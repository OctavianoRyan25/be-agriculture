package fertilizer

import (
	"net/http"

	"github.com/OctavianoRyan25/be-agriculture/base"
	"github.com/labstack/echo/v4"
	"strconv"
	"time"
)

func GetFertilizer(c echo.Context) error {
	fertilizer := []Fertilizer{{Id: 1, Name: "NPK", Compostition: "Natrium, Pottasium, Carbon", CreateAt: time.Now()}}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"fertilizer": fertilizer,
	})
}

func GetFertilizerById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("Id"))
	fertilizer := []Fertilizer{{Id: id, Name: "NPK", Compostition: "Natrium, Pottasium, Carbon", CreateAt: time.Now()}}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"fertilizer": fertilizer,
	})
}

// Create a new Fertilizer
func CreateFertilizer(c echo.Context) error {
	userId := c.Get("user_id").(uint)
	if userId == 0 {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: "Bad request",
			Code:    http.StatusBadRequest,
		}
		return c.JSON(http.StatusBadRequest, errRes)
	}
	role := c.Get("role").(string)
	if role != "admin" {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: "Forbidden access",
			Code:    http.StatusForbidden,
		}
		return c.JSON(http.StatusForbidden, errRes)

	}
	var fertilizer Fertilizer
	if err := c.Bind(&fertilizer); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"fertilizer": fertilizer,
	})
}

// Update a Fertilizer by Id
func UpdateFertilizer(c echo.Context) error {
	role := c.Get("role").(string)
	if role != "admin" {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: "Forbidden access",
			Code:    http.StatusForbidden,
		}
		return c.JSON(http.StatusForbidden, errRes)

	}
	userId := c.Get("user_id").(uint)
	if userId == 0 {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: "Bad request",
			Code:    http.StatusBadRequest,
		}
		return c.JSON(http.StatusBadRequest, errRes)
	}
	id, _ := strconv.Atoi(c.Param("Id"))
	fertilizerCatalog := Fertilizer{}
	c.Bind(&fertilizerCatalog)

	fertilizer := []Fertilizer{}

	for i, fertilizers := range fertilizer {
		if id == fertilizers.Id {
			fertilizer[i].Id = fertilizerCatalog.Id
			fertilizer[i].Name = fertilizerCatalog.Name
			fertilizer[i].Compostition = fertilizerCatalog.Compostition
		}
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"fertilizer": fertilizer,
	})
}

// Delete a Fertilizer by Id
func DeleteFertilizer(c echo.Context) error {
	role := c.Get("role").(string)
	if role != "admin" {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: "Forbidden access",
			Code:    http.StatusForbidden,
		}
		return c.JSON(http.StatusForbidden, errRes)

	}
	userId := c.Get("user_id").(uint)
	if userId == 0 {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: "Bad request",
			Code:    http.StatusBadRequest,
		}
		return c.JSON(http.StatusBadRequest, errRes)
	}
	id, _ := strconv.Atoi(c.Param("Id"))

	fertilizer := []Fertilizer{}

	indexToDelete := -1
	for i, fertilizers := range fertilizer {
		if fertilizers.Id == id {
			indexToDelete = i
			break
		}
	}

	if indexToDelete != -1 {
		fertilizer = append(fertilizer[:indexToDelete], fertilizer[indexToDelete+1:]...)
	}

	return c.NoContent(http.StatusNoContent)
}
