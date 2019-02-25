package services

import (
	"../dao"
	"../dataSource"
	"../models"
)

type ResultService interface {
	GetAll() []models.Result
	CountAll() int64
	Get(id int) *models.Result
	Delete(id int) error
	Update(data *models.Result, columns []string) error
	Insert(data *models.Result) error
}

type resultService struct {
	dao *dao.ResultDao
}

func NewResultService() ResultService {
	return &resultService{
		dao: dao.NewResultDao(dataSource.NewMysqlMaster()),
	}
}

func (this *resultService) GetAll() []models.Result {
	return this.dao.GetAll()
}

func (this *resultService) CountAll() int64 {
	return this.dao.CountAll()
}

func (this *resultService) Get(id int) *models.Result {
	return this.dao.Get(id)
}

func (this *resultService) Delete(id int) error {
	return this.dao.Delete(id)
}

func (this *resultService) Update(data *models.Result, columns []string) error {
	return this.dao.Update(data, columns)
}

func (this *resultService) Insert(data *models.Result) error {
	return this.dao.Insert(data)
}
