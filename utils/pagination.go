package utils

type Pagination struct {
	TotalPages int `json:"totalPages"`
	Limit      int `json:"limit"`
	Page       int `json:"page"`
	TotalRows  int `json:"totalRows"`
}

func NewPagination(totalPages int, limit int, page int, totalRows int) Pagination {
	return Pagination{TotalPages: totalPages, Limit: limit, Page: page, TotalRows: totalRows}
}
