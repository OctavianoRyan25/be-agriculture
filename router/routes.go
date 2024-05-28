package router

import (
	"fmt"

	"github.com/OctavianoRyan25/be-agriculture/handler"
	"github.com/OctavianoRyan25/be-agriculture/modules/plant"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB) {
	cloudinary, err := initCloudinary()
		if err != nil {
			fmt.Println("Failed to initialize Cloudinary:", err)
			return
		}

  plantCategoryRepository := plant.NewPlantCategoryRepository(db)
  plantCategoryService := plant.NewPlantCategoryService(plantCategoryRepository)
  plantCategoryHandler := handler.NewPlantCategoryHandler(plantCategoryService, cloudinary)

	plantRepository := plant.NewPlantRepository(db)
	plantService := plant.NewPlantService(plantRepository, plantCategoryRepository)
	plantHandler := handler.NewPlantHandler(plantService, cloudinary)

	v1 := e.Group("/api/v1")
	{
		v1.GET("/admin/plants/categories", plantCategoryHandler.GetAll)
		v1.GET("/admin/plants/categories/:id", plantCategoryHandler.GetByID)
		v1.POST("/admin/plants/categories", plantCategoryHandler.Create)
		v1.PUT("/admin/plants/categories/:id", plantCategoryHandler.Update)
		v1.DELETE("/admin/plants/categories/:id", plantCategoryHandler.Delete)
		
		v1.GET("/admin/plants", plantHandler.GetAll)            
		v1.GET("/admin/plants/:id", plantHandler.GetByID)        
		v1.POST("/admin/plants", plantHandler.Create)           
		v1.PUT("/admin/plants/:id", plantHandler.Update)         
		v1.DELETE("/admin/plants/:id", plantHandler.Delete) 
	}
}

func initCloudinary() (*cloudinary.Cloudinary, error) {
	cloudinaryURL := ""
	cloudinary, err := cloudinary.NewFromURL(cloudinaryURL)
	if err != nil {
			return nil, err
	}
	return cloudinary, nil
}