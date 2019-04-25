package services

import (
	"go-lottery/dao"
	"go-lottery/dataSource"
	"go-lottery/models"
)

type BlackUserService interface {
	GetAll() []models.BlackUser
	CountAll() int64
	Get(id int) *models.BlackUser
	Delete(id int) error
	Update(data *models.BlackUser, columns []string) error
	Insert(data *models.BlackUser) error
	GetByUid(uid int) *models.BlackUser
}

type blackUserService struct {
	dao *dao.BlackUserDao
}

func NewBlackUserService() BlackUserService {
	return &blackUserService{
		dao: dao.NewBlackUserDao(dataSource.NewMysqlMaster()),
	}
}

func (this *blackUserService) GetAll() []models.BlackUser {
	return this.dao.GetAll()
}

func (this *blackUserService) CountAll() int64 {
	return this.dao.CountAll()
}

func (this *blackUserService) Get(id int) *models.BlackUser {
	return this.dao.Get(id)
}

func (this *blackUserService) Delete(id int) error {
	return this.dao.Delete(id)
}

func (this *blackUserService) Update(data *models.BlackUser, columns []string) error {
	return this.dao.Update(data, columns)
}

func (this *blackUserService) Insert(data *models.BlackUser) error {
	return this.dao.Insert(data)
}

func (this *blackUserService) GetByUid(uid int) *models.BlackUser {
	return this.dao.GetByUid(uid)
}
