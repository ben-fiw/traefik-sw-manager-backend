package models

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/ben-fiw/go-database-bundle"
)

// ##############################
// #       MODEL METADATA       #
// ##############################

var AvailableVersionModelMeta = ModelMeta{
	StoreName:             "available_version",
	TableName:             "available_versions",
	DocumentName:          "available-version",
	PrimaryKey:            "id",
	DefaultOrderColumn:    "version",
	DefaultOrderDirection: "desc",
	DefaultPageSize:       10,
	DatabaseColumns:       []string{"id", "version", "display_name", "created_at", "updated_at"},
}

// ##############################
// #         MODEL STORE        #
// ##############################

func InitAvailableVersionModelStore(db *sql.DB) {
	database.GetModelStoreRegistry().Register(
		database.NewModelStore(db, AvailableVersionModelMeta.StoreName, &AvailableVersionModel{}, AvailableVersionModelMeta.TableName, AvailableVersionModelMeta.PrimaryKey),
	)
}

// ################################
// # AVAILABLE VERSION MODEL LIST #
// ################################

type AvailableVersionModelList []*AvailableVersionModel

func NewAvailableVersionModelList() *AvailableVersionModelList {
	return &AvailableVersionModelList{}
}

func (s *AvailableVersionModelList) Paginate(po PaginationParams) error {
	*s = make(AvailableVersionModelList, 0)

	offset := (po.Page - 1) * po.Limit

	// TODO: fix orderBy issues in the database package
	// paginateQueryBuilder := database.NewQueryBuilder(database.QueryTypeSelect, AvailableVersionModelMeta.TableName).
	// 	SetFields("id", "version", "display_name", "created_at", "updated_at").
	// 	SetLimit(po.Limit).
	// 	SetOffset(offset).
	// 	SetOrder(po.OrderBy)

	// query, params, err := paginateQueryBuilder.Build()
	// if err != nil {
	// 	return err
	// }

	query := fmt.Sprintf(
		"SELECT id, version, display_name, created_at, updated_at FROM %s ORDER BY %s %s LIMIT ? OFFSET ?",
		AvailableVersionModelMeta.TableName, po.OrderBy, po.OrderDirection,
	)
	params := []interface{}{po.Limit, offset}

	db, err := database.GetConnection()
	if err != nil {
		return err
	}

	rows, err := db.Query(query, params...)
	if err != nil {
		return err
	}

	for rows.Next() {
		availableVersion := &AvailableVersionModel{}
		err = availableVersion.FillFromRow(rows)
		if err != nil {
			return err
		}

		*s = append(*s, availableVersion)
	}

	return nil
}

func (s *AvailableVersionModelList) GetTotalCount() (int, error) {
	db, err := database.GetConnection()
	if err != nil {
		return 0, err
	}

	totalCountQueryBuilder := database.NewQueryBuilder(database.QueryTypeSelect, AvailableVersionModelMeta.TableName).
		SetFields("COUNT(id)")

	query, params, err := totalCountQueryBuilder.Build()
	if err != nil {
		return 0, err
	}

	row := db.QueryRow(query, params...)

	var totalCount int
	err = row.Scan(&totalCount)
	if err != nil {
		return 0, err
	}

	return totalCount, nil
}

// ##############################
// #   AVAILABLE VERSION MODEL  #
// ##############################

type AvailableVersionModel struct {
	Id          string `json:"id" xml:"id" yaml:"id"`
	Version     string `json:"version" xml:"version" yaml:"version"`
	DisplayName string `json:"displayName" xml:"displayName" yaml:"displayName"`
	CreatedAt   string `json:"createdAt" xml:"createdAt" yaml:"createdAt"`
	UpdatedAt   string `json:"updatedAt" xml:"updatedAt" yaml:"updatedAt"`
}

func (s *AvailableVersionModel) GetID() string {
	return s.Id
}

func (s *AvailableVersionModel) GetStore() *database.ModelStore {
	store, err := database.GetModelStoreRegistry().GetStore(AvailableVersionModelMeta.StoreName)
	if err != nil {
		panic(err)
	}
	return store
}

func (s *AvailableVersionModel) GetFieldNames() []string {
	return AvailableVersionModelMeta.DatabaseColumns
}

func (s *AvailableVersionModel) GetFieldValues() []interface{} {
	return []interface{}{
		s.Id,
		s.Version,
		s.DisplayName,
		s.CreatedAt,
		s.UpdatedAt,
	}
}

func (s *AvailableVersionModel) FillFromRow(row database.Scannable) error {
	return row.Scan(
		&s.Id,
		&s.Version,
		&s.DisplayName,
		&s.CreatedAt,
		&s.UpdatedAt,
	)
}

func (s *AvailableVersionModel) Create() error {
	s.CreatedAt = time.Now().Format(MySqlDateTimeFormat)
	s.UpdatedAt = time.Now().Format(MySqlDateTimeFormat)
	return s.GetStore().Create(context.TODO(), s)
}

func (s *AvailableVersionModel) Delete() error {
	return s.GetStore().Delete(context.TODO(), s.Id)
}

func (s *AvailableVersionModel) Update() error {
	s.UpdatedAt = time.Now().Format(MySqlDateTimeFormat)
	return s.GetStore().Update(context.TODO(), s)
}

func (s *AvailableVersionModel) Load() error {
	service, err := s.GetStore().Get(context.TODO(), s.Id)
	if err != nil {
		return err
	}

	loadedService, ok := service.(*AvailableVersionModel)
	if !ok {
		return fmt.Errorf("failed to load available version: not found")
	}

	*s = *loadedService
	return nil
}

// ##############################
// #     SERVICE RELATIONS      #
// ##############################

// TBD
