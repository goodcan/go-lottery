package services

import (
	"../dao"
	"../dataSource"
	"../models"
)

type GiftService interface {
	GetAll() []models.Gift
	CountAll() int64
	Get(id int) *models.Gift
	Delete(id int) error
	Update(data *models.Gift, columns []string) error
	Insert(data *models.Gift) error
}

type giftService struct {
	dao *dao.GiftDao
}

func NewGiftService() GiftService {
	return &giftService{
		dao: dao.NewGiftDao(dataSource.NewMysqlMaster()),
	}
}

func (this *giftService) GetAll() []models.Gift {
	return this.dao.GetAll()
}

func (this *giftService) CountAll() int64 {
	return this.dao.CountAll()
}

func (this *giftService) Get(id int) *models.Gift {
	return this.dao.Get(id)
}

func (this *giftService) Delete(id int) error {
	return this.dao.Delete(id)
}

func (this *giftService) Update(data *models.Gift, columns []string) error {
	return this.dao.Update(data, columns)
}

func (this *giftService) Insert(data *models.Gift) error {
	return this.dao.Insert(data)
}
