package main

import (
	"fmt"
	"net/http"

	"github.com/OctavianoRyan25/be-agriculture/configs"
	"github.com/OctavianoRyan25/be-agriculture/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	fmt.Println("Ini Branch Development!")

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db, err := configs.InitDB()
	if err != nil {
		panic("Failed to connect database")
	}

	err = configs.AutoMigrate(db)
	if err != nil {
		panic("Failed to migrate database")
	}

	router.SetupRoutes(e, db)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":8080"))
}

