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

	// Sementara gadipake karena katanya mau statis
	// climateConditionRepository := plant.NewClimateConditionRepository(db)
	// climateConditionHandler := handler.NewClimateConditionHandler(climateConditionService, cloudinary)
	// climateConditionService := plant.NewClimateConditionService(climateConditionRepository)


	plantRepository := plant.NewPlantRepository(db)
	plantService := plant.NewPlantService(plantRepository, plantCategoryRepository)
	plantHandler := handler.NewPlantHandler(plantService, cloudinary)

	v1 := e.Group("/api/v1")
	{
		v1.GET("/plants/categories", plantCategoryHandler.GetAll)
		v1.GET("/plants/categories/:id", plantCategoryHandler.GetByID)
		v1.POST("/plants/categories", plantCategoryHandler.Create)
		v1.PUT("/plants/categories/:id", plantCategoryHandler.Update)
		v1.DELETE("/plants/categories/:id", plantCategoryHandler.Delete)
		
		v1.GET("/plants", plantHandler.GetAll)            
		v1.GET("/plants/:id", plantHandler.GetByID)        
		v1.POST("/plants", plantHandler.Create)           
		v1.PUT("/plants/:id", plantHandler.Update)         
		v1.DELETE("/plants/:id", plantHandler.Delete) 

		// Sementara gadipake karena katanya mau statis
    
		// v1.GET("/climate-conditions", climateConditionHandler.GetAll)
		// v1.GET("/climate-conditions/:id", climateConditionHandler.GetByID)
		// v1.POST("/climate-conditions", climateConditionHandler.Create)
		// v1.PUT("/climate-conditions/:id", climateConditionHandler.Update)
		// v1.DELETE("/climate-conditions/:id", climateConditionHandler.Delete)
	}
}

func initCloudinary() (*cloudinary.Cloudinary, error) {
	cloudinaryURL := "cloudinary://985586974845469:PGaIo6o1qfPg54_o_zA4poyj49o@dxrz0cg5z"
	cloudinary, err := cloudinary.NewFromURL(cloudinaryURL)
	if err != nil {
			return nil, err
	}
	return cloudinary, nil
}