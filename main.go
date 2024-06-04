package main

import (
	"fmt"

	"github.com/OctavianoRyan25/be-agriculture/configs"
	"github.com/OctavianoRyan25/be-agriculture/handler"
	"github.com/OctavianoRyan25/be-agriculture/modules/admin"
	"github.com/OctavianoRyan25/be-agriculture/modules/plant"
	"github.com/OctavianoRyan25/be-agriculture/modules/search"
	"github.com/OctavianoRyan25/be-agriculture/modules/user"
	"github.com/OctavianoRyan25/be-agriculture/modules/weather"
	"github.com/OctavianoRyan25/be-agriculture/router"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	fmt.Println("Ini Branch Development!")

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	db, err := configs.InitDB()
	if err != nil {
		panic("Failed to connect database")
	}

	err = configs.AutoMigrate(db)
	if err != nil {
		panic("Failed to migrate database")
	}

	cloudinary, err := initCloudinary()
	if err != nil {
		fmt.Println("Failed to initialize Cloudinary:", err)
		return
	}

	repo := user.NewRepository(db)
	useCase := user.NewUseCase(repo)
	controller := user.NewUserController(useCase)

	repoAdmin := admin.NewRepository(db)
	useCaseAdmin := admin.NewUseCase(repoAdmin)
	controllerAdmin := admin.NewUserController(*useCaseAdmin)

	plantCategoryRepository := plant.NewPlantCategoryRepository(db)
	plantCategoryService := plant.NewPlantCategoryService(plantCategoryRepository)
	plantCategoryHandler := handler.NewPlantCategoryHandler(plantCategoryService, cloudinary)

	plantProgressRepository := plant.NewPlantProgressRepository(db)
	plantProgressService := plant.NewPlantProgressService(plantProgressRepository)
	plantProgressHandler := handler.NewPlantProgressHandler(plantProgressService, cloudinary)

	plantInstructionCategoryRepository := plant.NewPlantInstructionCategoryRepository(db)
	plantInstructionCategoryService := plant.NewPlantInstructionCategoryService(plantInstructionCategoryRepository)
	plantInstructionCategoryHandler := handler.NewPlantInstructionCategoryHandler(plantInstructionCategoryService, cloudinary)

	plantRepository := plant.NewPlantRepository(db)
	plantService := plant.NewPlantService(plantRepository, plantCategoryRepository)
	plantHandler := handler.NewPlantHandler(plantService, cloudinary)

	plantUserRepository := plant.NewUserPlantRepository(db)
	plantUserService := plant.NewUserPlantService(plantUserRepository)
	plantUserHandler := handler.NewUserPlantHandler(plantUserService)

	weatherService := weather.NewWeatherService()
	weatherHandler := handler.NewWeatherHandler(weatherService)

	searchRepository := search.NewRepository(db)
	searchUsecase := search.NewUsecase(searchRepository)
	searchController := search.NewSearchController(searchUsecase)

	router.InitRoutes(e, controller, controllerAdmin, plantCategoryHandler, plantHandler, plantUserHandler, weatherHandler, plantInstructionCategoryHandler, plantProgressHandler, searchController)

	e.Logger.Fatal(e.Start(":8080"))
}

func initCloudinary() (*cloudinary.Cloudinary, error) {
	//Production
	//cloudinaryURL := os.Getenv("CLOUDINARY_URL")
	//Development
	cloudinaryURL := "cloudinary://985586974845469:PGaIo6o1qfPg54_o_zA4poyj49o@dxrz0cg5z"
	cloudinary, err := cloudinary.NewFromURL(cloudinaryURL)
	if err != nil {
		return nil, err
	}
	return cloudinary, nil
}
