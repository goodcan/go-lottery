package services

import (
	"go-lottery/dao"
	"go-lottery/dataSource"
	"go-lottery/models"
)

type BlackIpService interface {
	GetAll() []models.BlackIp
	CountAll() int64
	Get(id int) *models.BlackIp
	Delete(id int) error
	Update(data *models.BlackIp, columns []string) error
	Insert(data *models.BlackIp) error
	GetByIp(ip string) *models.BlackIp
}

type blackIpService struct {
	dao *dao.BlackIpDao
}

func NewBlackIpService() BlackIpService {
	return &blackIpService{
		dao: dao.NewBlackIpDao(dataSource.NewMysqlMaster()),
	}
}

func (this *blackIpService) GetAll() []models.BlackIp {
	return this.dao.GetAll()
}

func (this *blackIpService) CountAll() int64 {
	return this.dao.CountAll()
}

func (this *blackIpService) Get(id int) *models.BlackIp {
	return this.dao.Get(id)
}

func (this *blackIpService) Delete(id int) error {
	return this.dao.Delete(id)
}

func (this *blackIpService) Update(data *models.BlackIp, columns []string) error {
	return this.dao.Update(data, columns)
}

func (this *blackIpService) Insert(data *models.BlackIp) error {
	return this.dao.Insert(data)
}

func (this *blackIpService) GetByIp(ip string) *models.BlackIp {
	return this.dao.GetByIp(ip)
}
