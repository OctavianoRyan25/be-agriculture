package router

import (
	"net/http"

	"github.com/OctavianoRyan25/be-agriculture/handler"
	"github.com/OctavianoRyan25/be-agriculture/middlewares"
	"github.com/OctavianoRyan25/be-agriculture/modules/admin"
	"github.com/OctavianoRyan25/be-agriculture/modules/user"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, userController *user.UserController, adminController *admin.AdminController, plantCategoryHandler *handler.PlantCategoryHandler, plantHandler *handler.PlantHandler, plantUserHandler *handler.UserPlantHandler, weatherHandler *handler.WeatherHandler, plantInstructionCategoryHandler *handler.PlantInstructionCategoryHandler, plantProgressHandler *handler.PlantProgressHandler) {
	group := e.Group("/api/v1")
	group.POST("/register", userController.RegisterUser)
	group.POST("/check-email", userController.CheckEmail)
	group.POST("/verify", userController.VerifyEmail)
	group.POST("/login", userController.Login)
	group.GET("/profile", userController.GetUserProfile, middlewares.Authentication())

	groupAdmin := e.Group("/api/v1/admin")
	groupAdmin.POST("/register", adminController.RegisterUser)
	groupAdmin.POST("/login", adminController.Login)
	groupAdmin.GET("/profile", adminController.GetUserProfile, middlewares.Authentication())

	group.GET("/plants/categories", plantCategoryHandler.GetAll)
	group.GET("/plants/categories/:id", plantCategoryHandler.GetByID)
	groupAdmin.POST("/plants/categories", plantCategoryHandler.Create, middlewares.Authentication())
	groupAdmin.PUT("/plants/categories/:id", plantCategoryHandler.Update, middlewares.Authentication())
	groupAdmin.DELETE("/plants/categories/:id", plantCategoryHandler.Delete, middlewares.Authentication())

	group.GET("/plants/progress/:plant_id", plantProgressHandler.GetAllByUserIDAndPlantID, middlewares.Authentication())
	group.POST("/plants/progress", plantProgressHandler.UploadProgress, middlewares.Authentication())

	group.GET("/plants/instructions/categories", plantInstructionCategoryHandler.GetAll)
	group.GET("/plants/instructions/categories/:id", plantInstructionCategoryHandler.GetByID)
	groupAdmin.POST("/plants/instructions/categories", plantInstructionCategoryHandler.Create, middlewares.Authentication())
	groupAdmin.PUT("/plants/instructions/categories/:id", plantInstructionCategoryHandler.Update, middlewares.Authentication())
	groupAdmin.DELETE("/plants/instructions/categories/:id", plantInstructionCategoryHandler.Delete, middlewares.Authentication())

	group.GET("/plants", plantHandler.GetAll)            
	group.GET("/plants/:id", plantHandler.GetByID)        
	group.GET("/plants/search", plantHandler.SearchPlantsByName)        
	groupAdmin.POST("/plants", plantHandler.Create, middlewares.Authentication())           
	groupAdmin.PUT("/plants/:id", plantHandler.Update, middlewares.Authentication())         
	groupAdmin.DELETE("/plants/:id", plantHandler.Delete, middlewares.Authentication())

	group.GET("/my/plants/:user_id", plantUserHandler.GetUserPlants, middlewares.Authentication())
	group.POST("/my/plants/add", plantUserHandler.AddUserPlant, middlewares.Authentication())
	group.DELETE("/my/plants/:user_plant_id", plantUserHandler.DeleteUserPlantByID, middlewares.Authentication())

	group.GET("/weather/current/:city", weatherHandler.GetCurrentWeather, middlewares.Authentication())
  group.GET("/weather/hourly/:city", weatherHandler.GetHourlyWeather, middlewares.Authentication())
  group.GET("/weather/daily/:city", weatherHandler.GetDailyWeather, middlewares.Authentication())

	group.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
}
