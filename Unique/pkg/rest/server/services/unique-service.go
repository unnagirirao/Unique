package services

import (
	"github.com/unnagirirao/Unique/unique/pkg/rest/server/daos"
	"github.com/unnagirirao/Unique/unique/pkg/rest/server/models"
)

type UniqueService struct {
	uniqueDao *daos.UniqueDao
}

func NewUniqueService() (*UniqueService, error) {
	uniqueDao, err := daos.NewUniqueDao()
	if err != nil {
		return nil, err
	}
	return &UniqueService{
		uniqueDao: uniqueDao,
	}, nil
}

func (uniqueService *UniqueService) CreateUnique(unique *models.Unique) (*models.Unique, error) {
	return uniqueService.uniqueDao.CreateUnique(unique)
}

func (uniqueService *UniqueService) UpdateUnique(id int64, unique *models.Unique) (*models.Unique, error) {
	return uniqueService.uniqueDao.UpdateUnique(id, unique)
}

func (uniqueService *UniqueService) DeleteUnique(id int64) error {
	return uniqueService.uniqueDao.DeleteUnique(id)
}

func (uniqueService *UniqueService) ListUniques() ([]*models.Unique, error) {
	return uniqueService.uniqueDao.ListUniques()
}

func (uniqueService *UniqueService) GetUnique(id int64) (*models.Unique, error) {
	return uniqueService.uniqueDao.GetUnique(id)
}
