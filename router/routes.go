package router

import (
	"net/http"

	"github.com/OctavianoRyan25/be-agriculture/middlewares"
	"github.com/OctavianoRyan25/be-agriculture/modules/admin"
	"github.com/OctavianoRyan25/be-agriculture/modules/fertilizer"
	"github.com/OctavianoRyan25/be-agriculture/modules/user"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, userController *user.UserController, adminController *admin.AdminController) {
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

	groupFertilizer := e.Group("/api/v1")
	groupFertilizer.GET("/fertilizer", fertilizer.GetFertilizer)
	groupFertilizer.GET("/fertilizer/:Id", fertilizer.GetFertilizerById)
	groupFertilizer.POST("/fertilizer", fertilizer.CreateFertilizer, middlewares.Authentication())
	groupFertilizer.PUT("/fertilizer/:Id", fertilizer.UpdateFertilizer, middlewares.Authentication())
	groupFertilizer.DELETE("/fertilizer/:Id", fertilizer.DeleteFertilizer, middlewares.Authentication())

	group.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
}
