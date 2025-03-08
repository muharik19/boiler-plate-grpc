package utils

import (
	"regexp"
	"strings"
)

type PaginationRequest struct {
	Limit int
	Page  int
	Field string
	Sort  string
}

func Pagination(limit int, page int, field string, sort string) PaginationRequest {
	isValidLetter := regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

	if !isValidLetter(field) {
		field = "created_at"
	}

	switch {
	case strings.ToLower(sort) == "asc":
		sort = "ASC"
	case strings.ToLower(sort) == "desc":
		sort = "DESC"
	default:
		sort = "ASC"
	}

	if limit == 0 || page == 0 {
		limit = 10
		page = 1
	}

	return PaginationRequest{
		Limit: limit,
		Page:  page,
		Field: field,
		Sort:  sort,
	}
}
