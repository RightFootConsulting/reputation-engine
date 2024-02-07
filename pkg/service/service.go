package service

import (
	"os"
	"rightfoot-consulting/reputation-engine/pkg/models"
	"rightfoot-consulting/reputation-engine/pkg/repos"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Holds shared services to be leveraged by api endpoints.
type ServiceManager struct {
	db                 *gorm.DB
	BlockEventRepo     repos.Repository[*models.BlockEvent]
	BlockSummaryRepo   repos.Repository[*models.BlockSummary]
	KarmaEventRepo     repos.Repository[*models.KarmaEvent]
	KarmaEventTypeRepo repos.Repository[*models.KarmaEventType]
	SocialEntityRepo   repos.Repository[*models.SocialEntity]
}

func (sm *ServiceManager) GetDB() *gorm.DB {
	return sm.db
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

	result = &ServiceManager{
		db,
		repos.Repository[*models.BlockEvent]{
			DB:          db,
			NewInstance: func() *models.BlockEvent { return &models.BlockEvent{} },
		},
		repos.Repository[*models.BlockSummary]{
			DB:          db,
			NewInstance: func() *models.BlockSummary { return &models.BlockSummary{} },
		},
		repos.Repository[*models.KarmaEvent]{
			DB:          db,
			NewInstance: func() *models.KarmaEvent { return &models.KarmaEvent{} },
		},
		repos.Repository[*models.KarmaEventType]{
			DB:          db,
			NewInstance: func() *models.KarmaEventType { return &models.KarmaEventType{} },
		},
		repos.Repository[*models.SocialEntity]{
			DB:          db,
			NewInstance: func() *models.SocialEntity { return &models.SocialEntity{} },
		},
	}
	return
}
