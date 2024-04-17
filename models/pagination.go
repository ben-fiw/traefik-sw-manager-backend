package models

import (
	"net/http"
	"strconv"
)

type PaginationParams struct {
	Page           int
	Limit          int
	OrderBy        string
	OrderDirection string
}

func GetPaginationParams(r *http.Request, mm ModelMeta) PaginationParams {
	page := 1
	limit := mm.DefaultPageSize
	orderBy := mm.DefaultOrderColumn
	orderDirection := mm.DefaultOrderDirection

	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")
	orderByStr := r.URL.Query().Get("order")
	orderDirectionStr := r.URL.Query().Get("order-direction")

	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}

	if limitStr != "" {
		limit, _ = strconv.Atoi(limitStr)
	}

	if orderByStr != "" {
		orderBy = validateOrderBy(orderByStr, mm)
	}

	if orderDirectionStr != "" {
		orderDirection = validateOrderDirection(orderDirectionStr, mm)
	}

	return PaginationParams{
		Page:           page,
		Limit:          limit,
		OrderBy:        orderBy,
		OrderDirection: orderDirection,
	}
}

func validateOrderBy(orderBy string, mm ModelMeta) string {
	for _, column := range mm.DatabaseColumns {
		if column == orderBy {
			return orderBy
		}
	}
	return mm.DefaultOrderColumn
}

func validateOrderDirection(orderDirection string, mm ModelMeta) string {
	if orderDirection == "asc" || orderDirection == "desc" {
		return orderDirection
	}
	return mm.DefaultOrderDirection
}
