package repository

import (
	"github.com/ASP1234/Aragog/pkg/entity"
	customError "github.com/ASP1234/Aragog/pkg/error"
	"sync"
)

// Repository for in-memory CRUD ops related to WebPage entity.
type localRepository struct {
	repo map[string]*entity.WebPage
	mu   sync.RWMutex
}

var (
	once sync.Once
	rep  *localRepository
)

// Constructs/ Returns a new/ existing LocalRepository object.
func LocalRepository() *localRepository {

	once.Do(func() {
		rep = &localRepository{
			repo: make(map[string]*entity.WebPage),
		}
	})

	return rep
}

// Puts the webPage into the store.
func (rep *localRepository) Put(webPage *entity.WebPage) (id string, err error) {

	rep.mu.Lock()
	defer rep.mu.Unlock()

	id = webPage.GetAddress().String()
	rep.repo[id] = webPage

	return id, nil
}

// Retrieves the webPage if present within the store.
func (rep *localRepository) Get(id string) (webPage *entity.WebPage, err error) {

	rep.mu.RLock()
	defer rep.mu.RUnlock()

	webPage, ok := rep.repo[id]

	if !ok {
		return webPage, customError.NewDataNotFoundError("WebPage not found for id: " + id)
	}

	return webPage, nil
}

// Retrieves all the webPages from the store.
func (rep *localRepository) BatchScan(exclusiveStartKey string) (
	scanResults []*entity.WebPage, lastEvaluatedKey string, err error) {

	rep.mu.RLock()
	defer rep.mu.RUnlock()

	scanResults = make([]*entity.WebPage, 0)

	for _, webPage := range rep.repo {
		scanResults = append(scanResults, webPage)
	}

	return scanResults, lastEvaluatedKey, err
}
