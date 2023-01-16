package dependencies

import (
	"compression/src/controllers"
	"compression/src/database"
	"compression/src/repository"
	"compression/src/services"
)

func CompressionDependencies() *controllers.CompressionController {

	mongo := database.MongoDB{
		Uri:            "mongodb://127.0.0.1:27017",
		DatabaseName:   "go",
		CollectionName: "compression",
	}

	mongo.Connect()

	repository := repository.CompressionRepository{Collection: mongo.Collection}

	service := services.CompressionService{CompressionRepository: &repository}

	controller := controllers.CompressionController{CompressionService: &service}

	return &controller
}
