package main

import (
	"fmt"
	"net/http"

	"github.com/OctavianoRyan25/be-agriculture/configs"
	"github.com/OctavianoRyan25/be-agriculture/middlewares"
	"github.com/OctavianoRyan25/be-agriculture/modules/admin"
	"github.com/OctavianoRyan25/be-agriculture/modules/user"
	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Ini Branch Development!")

	e := echo.New()

	db, err := configs.InitDB()
	if err != nil {
		panic("Failed to connect database")
	}

	// Auto migrate schema
	err = configs.AutoMigrate(db)
	if err != nil {
		panic("Failed to migrate database")
	}

	repo := user.NewRepository(db)
	useCase := user.NewUseCase(repo)
	controller := user.NewUserController(useCase)

	repoAdmin := admin.NewRepository(db)
	useCaseAdmin := admin.NewUseCase(repoAdmin)
	controllerAdmin := admin.NewUserController(*useCaseAdmin)

	group := e.Group("/api/v1")
	group.POST("/register", controller.RegisterUser)
	group.POST("/check-email", controller.CheckEmail)
	group.POST("/verify", controller.VerifyEmail)
	group.POST("/login", controller.Login)
	group.GET("/profile", controller.GetUserProfile, middlewares.Authentication())

	groupAdmin := e.Group("/api/v1/admin")
	groupAdmin.POST("/register", controllerAdmin.RegisterUser)
	groupAdmin.POST("/login", controllerAdmin.Login)
	groupAdmin.GET("/profile", controllerAdmin.GetUserProfile, middlewares.Authentication())

	group.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
