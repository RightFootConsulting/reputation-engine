package routes

import (
	"fmt"
	"net/http"
	"net/url"
	"rightfoot-consulting/reputation-engine/pkg/models"
	"rightfoot-consulting/reputation-engine/pkg/query"
	"rightfoot-consulting/reputation-engine/pkg/repos"
	"rightfoot-consulting/reputation-engine/pkg/routes/errors"

	logger "rightfoot-consulting/reputation-engine/pkg/logging"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DeleteResult struct {
	RowsAffected int64 `json:"rows_affected"`
}

// We will assume that we don't need to verify that Sm has been
// properly initialized herein.  If it hasn't though then
// we will get null pointer exceptions.
type CrudqApi[MT models.Model] struct {
	Repo    *repos.Repository[MT]
	Niluuid uuid.UUID
}

func (api *CrudqApi[MT]) Create(c *gin.Context) {
	repo := api.Repo
	body := repo.NewInstance()
	err := c.BindJSON(body)
	if err != nil {
		logger.Errorf("Failed to bind to request json: %v", err)
		result := errors.InvalidRequest(err, true)
		c.JSON(result.Status, result)
		return
	}
	body.SetId(api.Niluuid)
	saved, err := repo.Save(body)
	if err != nil {
		logger.Errorf("Failed to save request json: %v", err)
		result := errors.Forbidden(err, false)
		c.JSON(result.Status, result)
		return
	}
	c.JSON(http.StatusOK, saved)
}

func (api *CrudqApi[MT]) Update(c *gin.Context) {
	repo := api.Repo
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		logger.Errorf("Failed to bind to request id param: %v", err)
		result := errors.InvalidRequest(err, true)
		c.JSON(result.Status, result)
		return
	}
	body := repo.NewInstance()
	err = c.BindJSON(body)
	if err != nil {
		logger.Errorf("Failed to bind to request json: %v", err)
		result := errors.InvalidRequest(err, true)
		c.JSON(result.Status, result)
		return
	}
	if body.GetId() != id {
		logger.Errorf("Path parameter does not match body id: %v", err)
		err = fmt.Errorf("path parameter id '%s' does not match body id '%s'", id.String(), body.GetId().String())
		result := errors.InvalidRequest(err, true)
		c.JSON(result.Status, result)
		return
	}
	saved, err := repo.Save(body)
	if err != nil {
		logger.Errorf("Failed to save request json: %v", err)
		result := errors.Forbidden(err, false)
		c.JSON(result.Status, result)
		return
	}
	c.JSON(http.StatusOK, saved)
}

func (api *CrudqApi[MT]) Delete(c *gin.Context) {
	repo := api.Repo
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		logger.Errorf("Failed to bind to request id param: %v", err)
		result := errors.InvalidRequest(err, true)
		c.JSON(result.Status, result)
		return
	}

	rows, err := repo.Delete(id)
	if err != nil {
		logger.Errorf("Failed to delete: %v", err)
		result := errors.Forbidden(err, false)
		c.JSON(result.Status, result)
		return
	}
	c.JSON(http.StatusOK, &DeleteResult{RowsAffected: rows})
}

func (api *CrudqApi[MT]) Get(c *gin.Context) {
	repo := api.Repo
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		logger.Errorf("Failed to bind to request id param: %v", err)
		result := errors.InvalidRequest(err, true)
		c.JSON(result.Status, result)
		return
	}

	found, err := repo.Get(id)
	if err != nil {
		logger.Errorf("Failed to get by id %s: %v", id.String(), err)
		result := errors.NotFoundError(err, false)
		c.JSON(result.Status, result)
		return
	}
	c.JSON(http.StatusOK, found)
}

func (api *CrudqApi[MT]) Query(c *gin.Context) {
	repo := api.Repo
	values, err := url.ParseQuery(c.Request.URL.RawQuery)
	if err != nil {
		logger.Errorf("Failed to parse query string '%s': %v", c.Request.URL.RawQuery, err)
		result := errors.InvalidRequest(err, false)
		c.JSON(result.Status, result)
		return
	}
	qcfg, err := query.NewQueryConfig(values, models.BlockEvent{})
	if err != nil {
		logger.Errorf("Failed to parse query values '%s': %v", c.Request.URL.RawQuery, err)
		result := errors.InvalidRequest(err, false)
		c.JSON(result.Status, result)
		return
	}
	found, err := repo.Query(qcfg)
	if err != nil {
		logger.Errorf("Failed to execute query: %v", c.Request.URL.RawQuery, err)
		result := errors.InvalidRequest(err, false)
		c.JSON(result.Status, result)
		return
	}
	c.JSON(http.StatusOK, found)
}
