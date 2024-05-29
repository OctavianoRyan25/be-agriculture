package router

import (
	"net/http"

	"github.com/OctavianoRyan25/be-agriculture/handler"
	"github.com/OctavianoRyan25/be-agriculture/middlewares"
	"github.com/OctavianoRyan25/be-agriculture/modules/admin"
	"github.com/OctavianoRyan25/be-agriculture/modules/user"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, userController *user.UserController, adminController *admin.AdminController, plantCategoryHandler *handler.PlantCategoryHandler, plantHandler *handler.PlantHandler, plantUserHandler *handler.UserPlantHandler) {
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

	group.GET("/plants", plantHandler.GetAll)            
	group.GET("/plants/:id", plantHandler.GetByID)        
	groupAdmin.POST("/plants", plantHandler.Create, middlewares.Authentication())           
	groupAdmin.PUT("/plants/:id", plantHandler.Update, middlewares.Authentication())         
	groupAdmin.DELETE("/plants/:id", plantHandler.Delete, middlewares.Authentication())

	group.GET("/my/plants/:user_id", plantUserHandler.GetUserPlants)
	group.POST("/my/plants/add", plantUserHandler.AddUserPlant)
	group.DELETE("/my/plants/:user_plant_id", plantUserHandler.DeleteUserPlantByID)

	group.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
}
