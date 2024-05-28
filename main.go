package main

import (
	"fmt"

	"github.com/OctavianoRyan25/be-agriculture/configs"
	"github.com/OctavianoRyan25/be-agriculture/modules/admin"
	"github.com/OctavianoRyan25/be-agriculture/modules/user"
	"github.com/OctavianoRyan25/be-agriculture/router"
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

	router.InitRoutes(e, controller, controllerAdmin)

	e.Logger.Fatal(e.Start(":8080"))
}
