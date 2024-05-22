package main

import (
	"fmt"
	"net/http"

	"github.com/OctavianoRyan25/be-agriculture/configs"
	"github.com/labstack/echo"
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

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
