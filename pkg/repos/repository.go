package repos

import (
	"errors"
	logger "rightfoot-consulting/reputation-engine/pkg/logging"
	"rightfoot-consulting/reputation-engine/pkg/models"
	"rightfoot-consulting/reputation-engine/pkg/query"
	"strconv"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type InstanceFactory[MT models.Model] func() MT

type Repository[MT models.Model] struct {
	DB          *gorm.DB
	NewInstance InstanceFactory[MT]
}

func (repo *Repository[MT]) Save(data MT) (saved MT, err error) {
	db := repo.DB
	if db == nil {
		err = errors.New("database not initialized")
		return
	}
	tx := db.Clauses(clause.Returning{}).Save(data)
	if tx.Error != nil {
		err = tx.Error
		return
	}
	saved = data
	return
}

func (repo *Repository[MT]) Delete(id uuid.UUID) (count int64, err error) {
	db := repo.DB
	if db == nil {
		err = errors.New("database not initialized")
		return
	}
	tx := db.Where("id = ?", id).Delete(repo.NewInstance())
	if tx.Error != nil {
		err = tx.Error
	}
	count = tx.RowsAffected
	return
}

func (repo *Repository[MT]) Get(id uuid.UUID) (found MT, err error) {
	db := repo.DB
	if db == nil {
		err = errors.New("database not initialized")
		return
	}
	found = repo.NewInstance()
	tx := db.Where("id = ?", id).First(found)
	if tx.Error != nil {
		err = tx.Error
	}
	return
}

func (repo *Repository[MT]) Query(qConfig *query.QueryConfig) (found []MT, err error) {
	db := repo.DB
	if db == nil {
		err = errors.New("database not initialized")
		return
	}
	args := make([]any, len(qConfig.Args))
	i := 0
	for _, v := range qConfig.Args {
		args[i] = v
		i++
	}
	tx := db.Model(repo.NewInstance()).Where(qConfig.Pattern, args...)
	if qConfig.SortBy != "" {
		tx.Order(qConfig.SortBy)
	}
	var limit int64
	var page int64
	if qConfig.Limit != "" {
		limit, err = strconv.ParseInt(qConfig.Limit, 10, 64)
		if err != nil {
			logger.Errorf("cannot parse limit parameter %s: %v", qConfig.Limit, err)
			return
		}
	} else {
		limit = 100
	}
	if qConfig.Page != "" {
		page, err = strconv.ParseInt(qConfig.Page, 10, 64)
		if err != nil {
			logger.Errorf("cannot parse page parameter %s: %v", qConfig.Limit, err)
			return
		}
	} else {
		page = 0
	}
	offset := page * limit
	rows, err := tx.Offset(int(offset)).Rows()
	if rows != nil {
		defer rows.Close()
	}
	if err != nil {
		logger.Errorf("query results in error: %v", err)
		return
	}
	found = make([]MT, 0, limit)
	n := 0
	for rows.Next() {
		var row MT = repo.NewInstance()
		tx.ScanRows(rows, row)
		found = append(found, row)
		n++
	}
	return
}
