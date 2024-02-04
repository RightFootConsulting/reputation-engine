package service

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Holds shared services to be leveraged by api endpoints.
type ServiceManager struct {
	db *gorm.DB
}

var serviceManager *ServiceManager = nil

func GetServiceManager() (result *ServiceManager, err error) {
	if serviceManager == nil {
		serviceManager, err = createServiceManager()
	}
	result = serviceManager
	return
}

func createServiceManager() (result *ServiceManager, err error) {
	dsn := os.Getenv("POSTGRESQL_DSN")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
}
