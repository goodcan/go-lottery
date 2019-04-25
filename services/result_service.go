package services

import (
	"go-lottery/dao"
	"go-lottery/dataSource"
	"go-lottery/models"
)

type ResultService interface {
	GetAll() []models.Result
	CountAll() int64
	CountByGift(giftId int) int64
	CountByUser(uid int) int64
	Get(id int) *models.Result
	Delete(id int) error
	Update(data *models.Result, columns []string) error
	Insert(data *models.Result) error
	SearchByGift(giftId, page, size int) []models.Result
	SearchByUser(uid, page, size int) []models.Result
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

func (this *resultService) CountByGift(giftId int) int64 {
	return this.dao.CountByGift(giftId)
}
func (this *resultService) CountByUser(uid int) int64 {
	return this.dao.CountByUser(uid)
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

func (this *resultService) SearchByGift(giftId, page, size int) []models.Result {
	return this.dao.SearchByGift(giftId, page, size)
}

func (this *resultService) SearchByUser(uid, page, size int) []models.Result {
	return this.dao.SearchByUser(uid, page, size)
}
