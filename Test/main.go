package main

import (
	endpointcount "delos/endpointCount"
	"delos/farm"
	"delos/handler"
	"delos/pounds"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/delos?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Db Connection Error")
	}

	err = db.AutoMigrate(&farm.Farm{}, &pounds.Pounds{}, &endpointcount.Statistics{})
	if err != nil {
		log.Fatal("AutoMigrate Error")
		}
	

	statisticsRepository := endpointcount.NewStatisticsRepository(db)

	// Inisialisasi service
	statisticsService := endpointcount.NewStatisticsService(statisticsRepository)

	// Inisialisasi handler
	statisticsHandler := handler.NewStatisticsHandler(statisticsService)

	farmRepository := farm.NewRepository(db)
	farmService := farm.NewService(farmRepository)
	farmHandler := handler.NewFarmHandler(farmService, statisticsService)

	poundsRepository := pounds.NewRepository(db)
	poundsService := pounds.NewService(poundsRepository, farmRepository)
	poundsHandler := handler.NewPoundsHandler(poundsService, statisticsService)

	router := gin.Default()
	api := router.Group("/farm")
	api1 := router.Group("/pounds")
	
	//farm
	api.POST("/create", farmHandler.CreateFarm)
	api.PUT("/:id", farmHandler.UpdatedFarm)
	api.DELETE("/:id", farmHandler.DeletedFarm)
	// api.DELETE("cek/:id", farmHandler.FindDeletedFarm)
	api.GET("/", farmHandler.GetFarms)
	api.GET("/:id", farmHandler.GetOneFarm)

	//pounds
	api1.POST("/create", poundsHandler.CreatePounds)
	api1.PUT("/:id", poundsHandler.UpdatedPounds)
	api1.DELETE("/:id", poundsHandler.DeletedPounds)
	api1.GET("/", poundsHandler.GetPounds)
	api1.GET("/:id", poundsHandler.GetOnePounds)

	router.GET("/statistics", statisticsHandler.GetStatisticsHandler)



	router.Run(":8080")

}
