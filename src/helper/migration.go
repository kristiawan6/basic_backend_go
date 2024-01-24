package helper

import (
	"backend_golang/src/config"
	"backend_golang/src/models"
)

func Migration() {
	config.DB.AutoMigrate(&models.Product{})
}
