package response

import (
	"dev/models/domain"
	"time"
)

type ProductDetail struct {
	List       []domain.Product `json:"list"`
	TotalItems int              `json:"total_items"`
	TotalPages int              `json:"total_pages"`
	Page       int              `json:"page"`
	PageSize   int              `json:"page_size"`
	Start      time.Time        `json:"start"`
	Finish     time.Time        `json:"finish"`
	Duration   string           `json:"duration"`
}

type ProductListResponse struct {
	CorrelationID string    `json:"correlationid"`
	Success       bool      `json:"success"`
	Error         string    `json:"error"`
	Tin           time.Time `json:"tin"`
	Tout          time.Time `json:"tout"`
	Data          ProductDetail
}

type ProductErrorResponse struct {
	CorrelationID string    `json:"correlationid"`
	Success       bool      `json:"success"`
	Error         string    `json:"error"`
	Tin           time.Time `json:"tin"`
	Tout          time.Time `json:"tout"`
}
