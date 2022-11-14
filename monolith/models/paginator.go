package models

import (
	"math"
	"net/url"
	"strconv"
)

const (
	paginatorDefaultPage  = 1
	paginatorDefaultLimit = 10
	paginatorMaxLimit     = 100
)

type Paginator struct {
	Page       int32 `json:"page"`
	Limit      int32 `json:"limit"`
	TotalPages int64 `json:"total_pages"`
}

type PaginationGameSlice struct {
	Games     GameSlice `json:"games"`
	Paginator Paginator `json:"paginator"`
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

func (p *Paginator) SetTotalPages(countGames int64) {
	p.TotalPages = int64(math.Ceil(float64(countGames) / float64(p.Limit)))
}

func NewPaginationGameSlice(g GameSlice, p Paginator) PaginationGameSlice {
	return PaginationGameSlice{
		Games:     g,
		Paginator: p,
	}
}
