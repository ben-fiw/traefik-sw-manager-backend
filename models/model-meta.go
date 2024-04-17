package models

type ModelMeta struct {
	StoreName             string
	TableName             string
	DocumentName          string
	PrimaryKey            string
	DefaultOrderColumn    string
	DefaultOrderDirection string
	DefaultPageSize       int
	DatabaseColumns       []string
}
