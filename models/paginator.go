package models

import (
	"net/url"
	"strconv"
)

const (
	paginatorDefaultPage  = 1
	paginatorDefaultLimit = 10
	paginatorMaxLimit     = 100
)

type Paginator struct {
	Page  int32
	Limit int32
}

func NewPaginatorFromQuery(v url.Values) Paginator {
	page, err := strconv.Atoi(v.Get("page"))
	if err != nil || page < paginatorDefaultPage {
		page = paginatorDefaultPage
	}

	limit, err := strconv.Atoi(v.Get("limit"))
	if err != nil || limit < paginatorDefaultLimit {
		limit = paginatorDefaultLimit
	}
	if limit > paginatorMaxLimit {
		limit = paginatorMaxLimit
	}

	return Paginator{
		Page:  int32(page),
		Limit: int32(limit),
	}
}

func (p *Paginator) Offset() int32 {
	return (p.Page - 1) * p.Limit
}
