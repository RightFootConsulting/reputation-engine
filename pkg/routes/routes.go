package routes

import (
	"rightfoot-consulting/reputation-engine/pkg/models"
	"rightfoot-consulting/reputation-engine/pkg/repos"
	"rightfoot-consulting/reputation-engine/pkg/service"

	"github.com/gin-gonic/gin"
)

type RouteManager struct {
	ServiceManager     *service.ServiceManager
	Router             *gin.Engine
	BlockEventsApi     *CrudqApi[*models.BlockEvent]
	BlockSummariesApi  *CrudqApi[*models.BlockSummary]
	KarmaEventsApi     *CrudqApi[*models.KarmaEvent]
	KarmaEventTypesApi *CrudqApi[*models.KarmaEventType]
	SocialEntitiesApi  *CrudqApi[*models.SocialEntity]
}

func Create(r *gin.Engine, sm *service.ServiceManager) (result *RouteManager, err error) {

	blockEvents := r.Group("/block-events")
	blockSummaries := r.Group("/block-summaries")
	karmaEvents := r.Group("/karma-events")
	karmaEventTypes := r.Group("/karma-event-types")
	socialEntities := r.Group("/social-entites")

	result = &RouteManager{
		ServiceManager:     sm,
		Router:             r,
		BlockEventsApi:     MapApi(blockEvents, &sm.BlockEventRepo),
		BlockSummariesApi:  MapApi(blockSummaries, &sm.BlockSummaryRepo),
		KarmaEventsApi:     MapApi(karmaEvents, &sm.KarmaEventRepo),
		KarmaEventTypesApi: MapApi(karmaEventTypes, &sm.KarmaEventTypeRepo),
		SocialEntitiesApi:  MapApi(socialEntities, &sm.SocialEntityRepo),
	}
	return
}

func MapApi[MT models.Model](rg *gin.RouterGroup, repo *repos.Repository[MT]) *CrudqApi[MT] {
	api := &CrudqApi[MT]{
		Repo: repo,
	}
	rg.POST("/", api.Create)
	rg.PUT("/:id", api.Update)
	rg.DELETE("/:id", api.Delete)
	rg.GET("/:id", api.Get)
	rg.GET("/", api.Query)

	return api
}
