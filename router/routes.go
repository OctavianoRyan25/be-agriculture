package router

import (
	"github.com/OctavianoRyan25/be-agriculture/handler"
	"github.com/OctavianoRyan25/be-agriculture/modules/plant"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB) {
	plantCategoryRepository := plant.NewPlantCategoryRepository(db)
	plantCategoryService := plant.NewPlantCategoryService(plantCategoryRepository)
	plantCategoryHandler := handler.NewPlantCategoryHandler(plantCategoryService)

	climateConditionRepository := plant.NewClimateConditionRepository(db)
	climateConditionService := plant.NewClimateConditionService(climateConditionRepository)
	climateConditionHandler := handler.NewClimateConditionHandler(climateConditionService)

	plantRepository := plant.NewPlantRepository(db)
	plantService := plant.NewPlantService(plantRepository)
	plantHandler := handler.NewPlantHandler(plantService)

	v1 := e.Group("/api/v1")
	{
		v1.GET("/plants/categories", plantCategoryHandler.GetAll)
		v1.GET("/plants/categories/:id", plantCategoryHandler.GetByID)
		v1.POST("/plants/categories", plantCategoryHandler.Create)
		v1.PUT("/plants/categories/:id", plantCategoryHandler.Update)
		v1.DELETE("/plants/categories/:id", plantCategoryHandler.Delete)

		v1.GET("/climate-conditions", climateConditionHandler.GetAll)
		v1.GET("/climate-conditions/:id", climateConditionHandler.GetByID)
		v1.POST("/climate-conditions", climateConditionHandler.Create)
		v1.PUT("/climate-conditions/:id", climateConditionHandler.Update)
		v1.DELETE("/climate-conditions/:id", climateConditionHandler.Delete)

		v1.GET("/plants", plantHandler.GetAll)            
		v1.GET("/plants/:id", plantHandler.GetByID)        
		v1.POST("/plants", plantHandler.Create)           
		v1.PUT("/plants/:id", plantHandler.Update)         
		v1.DELETE("/plants/:id", plantHandler.Delete)     
	}
}
