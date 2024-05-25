package controller

import (
	"net/http"

	"strconv"
	"time"

	"github.com/OctavianoRyan25/be-agriculture/modules"
	"github.com/labstack/echo"
)

func GetFertilizer(c echo.Context) error {
	fertilizer := []modules.Fertilizer{{Id: 1, Name: "NPK", Compostition: "Natrium, Pottasium, Carbon", CreateAt: time.Now()}}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"fertilizer": fertilizer,
	})
}

func GetFertilizerById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("Id"))
	fertilizer := []modules.Fertilizer{{Id: id, Name: "NPK", Compostition: "Natrium, Pottasium, Carbon", CreateAt: time.Now()}}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"fertilizer": fertilizer,
	})
}

// Create a new Fertilizer
func CreateFertilizer(c echo.Context) error {
	var fertilizer modules.Fertilizer
	if err := c.Bind(&fertilizer); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"fertilizer": fertilizer,
	})
}

// Update a Fertilizer by Id
func UpdateFertilizer(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("Id"))
	fertilizerCatalog := modules.Fertilizer{}
	c.Bind(&fertilizerCatalog)

	fertilizer := []modules.Fertilizer{}

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
	id, _ := strconv.Atoi(c.Param("Id"))

	fertilizer := []modules.Fertilizer{}

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
