package models

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/ben-fiw/go-database-bundle"
)

// ##############################
// #       MODEL METADATA       #
// ##############################

var DemoInstanceModelMeta = ModelMeta{
	StoreName:             "demo_instance",
	TableName:             "demo_instances",
	DocumentName:          "demo-instance",
	PrimaryKey:            "id",
	DefaultOrderColumn:    "name",
	DefaultOrderDirection: "asc",
	DefaultPageSize:       10,
	DatabaseColumns:       []string{"id", "version_id", "name", "slug", "status", "docker_id", "domain", "path", "created_at", "updated_at"},
}

type DemoInstanceModelStatusId struct {
	Id   int64  `json:"id" xml:"id" yaml:"id"`
	Key  string `json:"key" xml:"key" yaml:"key"`
	Type string `json:"type" xml:"type" yaml:"type"`
}

var DemoInstanceModelStatusIdes = map[int64]DemoInstanceModelStatusId{
	-1: {Id: -1, Key: "unknown", Type: "danger"},
	0:  {Id: 0, Key: "stopped", Type: "danger"},
	1:  {Id: 1, Key: "running", Type: "success"},
	2:  {Id: 2, Key: "starting", Type: "warning"},
	3:  {Id: 3, Key: "stopping", Type: "warning"},
}

// ##############################
// #         MODEL STORE        #
// ##############################

func InitDemoInstanceModelStore(db *sql.DB) {
	database.GetModelStoreRegistry().Register(
		database.NewModelStore(
			db,
			DemoInstanceModelMeta.StoreName,
			&DemoInstanceModel{},
			DemoInstanceModelMeta.TableName,
			DemoInstanceModelMeta.PrimaryKey,
			DemoInstanceModelFactory,
		),
	)
}

// ################################
// #   DEMO INSTANCE MODEL LIST   #
// ################################

type DemoInstanceModelList []*DemoInstanceModel

func NewDemoInstanceModelList() *DemoInstanceModelList {
	return &DemoInstanceModelList{}
}

func (s *DemoInstanceModelList) Paginate(po PaginationParams) error {
	*s = make(DemoInstanceModelList, 0)

	offset := (po.Page - 1) * po.Limit

	// TODO: fix orderBy issues in the database package
	// paginateQueryBuilder := database.NewQueryBuilder(database.QueryTypeSelect, DemoInstanceModelMeta.TableName).
	// 	SetFields(DemoInstanceModelMeta.DatabaseColumns...).
	// 	SetLimit(po.Limit).
	// 	SetOffset(offset).
	// 	SetOrder(po.OrderBy)

	// query, params, err := paginateQueryBuilder.Build()
	// if err != nil {
	// 	return err
	// }

	query := fmt.Sprintf(
		"SELECT %s FROM %s ORDER BY %s %s LIMIT ? OFFSET ?",
		strings.Join(DemoInstanceModelMeta.DatabaseColumns, ", "),
		DemoInstanceModelMeta.TableName,
		po.OrderBy, po.OrderDirection,
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
		demoInstance := &DemoInstanceModel{}
		err = demoInstance.FillFromRow(rows)
		if err != nil {
			return err
		}

		*s = append(*s, demoInstance)
	}

	return nil
}

func (s *DemoInstanceModelList) GetTotalCount() (int, error) {
	db, err := database.GetConnection()
	if err != nil {
		return 0, err
	}

	totalCountQueryBuilder := database.NewQueryBuilder(database.QueryTypeSelect, DemoInstanceModelMeta.TableName).
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

func (s *DemoInstanceModelList) LoadVersions() error {
	// Collect all version IDs
	versionIds := make([]interface{}, 0)
	for _, demoInstance := range *s {
		// skip, if the version ID is already collected
		for _, versionId := range versionIds {
			if versionId == demoInstance.VersionId {
				continue
			}
		}

		versionIds = append(versionIds, demoInstance.VersionId)
	}

	// Load all required versions
	versionModelStore, err := database.GetModelStoreRegistry().GetStore(AvailableVersionModelMeta.StoreName)
	if err != nil {
		return err
	}

	dbQuery := database.NewEqualsAnyQuery("id", false, versionIds...)
	fmt.Printf("Query: %+v\n", dbQuery)
	versions, err := versionModelStore.Search(context.TODO(), dbQuery)
	if err != nil {
		return err
	}

	// Set the versions for the models
	for _, demoInstance := range *s {
		for _, version := range versions {
			if version.GetID() == demoInstance.VersionId {
				demoInstance.Version = version.(*AvailableVersionModel)
				break
			}
		}
	}

	return nil
}

// ##############################
// #     DEMO INSTANCE MODEL  #
// ##############################

type DemoInstanceModel struct {
	Id        string                    `json:"id" xml:"id" yaml:"id"`
	VersionId string                    `json:"-" xml:"-" yaml:"-"`
	Version   *AvailableVersionModel    `json:"version" xml:"version" yaml:"version"`
	Name      string                    `json:"name" xml:"name" yaml:"name"`
	Slug      string                    `json:"slug" xml:"slug" yaml:"slug"`
	StatusId  int64                     `json:"-" xml:"-" yaml:"-"`
	Status    DemoInstanceModelStatusId `json:"status" xml:"status" yaml:"status"`
	DockerId  string                    `json:"dockerId" xml:"dockerId" yaml:"dockerId"`
	Domain    string                    `json:"domain" xml:"domain" yaml:"domain"`
	Path      string                    `json:"path" xml:"path" yaml:"path"`
	CreatedAt string                    `json:"createdAt" xml:"createdAt" yaml:"createdAt"`
	UpdatedAt string                    `json:"updatedAt" xml:"updatedAt" yaml:"updatedAt"`
}

func DemoInstanceModelFactory() database.Model {
	return &DemoInstanceModel{}
}

func (s *DemoInstanceModel) GetID() string {
	return s.Id
}

func (s *DemoInstanceModel) GetStore() *database.ModelStore {
	store, err := database.GetModelStoreRegistry().GetStore(DemoInstanceModelMeta.StoreName)
	if err != nil {
		panic(err)
	}
	return store
}

func (s *DemoInstanceModel) GetFieldNames() []string {
	return DemoInstanceModelMeta.DatabaseColumns
}

func (s *DemoInstanceModel) GetFieldValues() []interface{} {
	return []interface{}{
		s.Id,
		s.VersionId,
		s.Name,
		s.Slug,
		s.StatusId,
		s.DockerId,
		s.Domain,
		s.Path,
		s.CreatedAt,
		s.UpdatedAt,
	}
}

func (s *DemoInstanceModel) FillFromRow(row database.Scannable) error {
	err := row.Scan(
		&s.Id,
		&s.VersionId,
		&s.Name,
		&s.Slug,
		&s.StatusId,
		&s.DockerId,
		&s.Domain,
		&s.Path,
		&s.CreatedAt,
		&s.UpdatedAt,
	)
	if err != nil {
		return err
	}

	status, ok := DemoInstanceModelStatusIdes[s.StatusId]
	if !ok {
		status = DemoInstanceModelStatusIdes[-1]
	}

	s.Status = status
	return nil
}

func (s *DemoInstanceModel) Create() error {
	s.CreatedAt = time.Now().Format(MySqlDateTimeFormat)
	s.UpdatedAt = time.Now().Format(MySqlDateTimeFormat)
	return s.GetStore().Create(context.TODO(), s)
}

func (s *DemoInstanceModel) Delete() error {
	return s.GetStore().Delete(context.TODO(), s.Id)
}

func (s *DemoInstanceModel) Update() error {
	s.UpdatedAt = time.Now().Format(MySqlDateTimeFormat)
	return s.GetStore().Update(context.TODO(), s)
}

func (s *DemoInstanceModel) Load() error {
	service, err := s.GetStore().Get(context.TODO(), s.Id)
	if err != nil {
		return err
	}

	loadedService, ok := service.(*DemoInstanceModel)
	if !ok {
		return fmt.Errorf("failed to load DEMO INSTANCE: not found")
	}

	*s = *loadedService
	return nil
}

// ##############################
// #     SERVICE RELATIONS      #
// ##############################

func (s *DemoInstanceModel) LoadVersion() (*AvailableVersionModel, error) {
	version := &AvailableVersionModel{Id: s.VersionId}
	err := version.Load()
	if err != nil {
		return nil, err
	}

	s.Version = version

	return version, nil
}
