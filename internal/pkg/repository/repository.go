package repository

import "Aragog/internal/pkg/entity"

// Interface exposed for CRUD ops related to WebPage entity
type Repository interface {
	Put(webPage *entity.WebPage) (id string, err error)
	Get(id string) (webPage *entity.WebPage, err error)
	BatchScan(exclusiveStartKey string) (scanResults []*entity.WebPage, lastEvaluatedKey string, err error)
}
