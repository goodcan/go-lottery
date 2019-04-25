package services

import (
	"go-lottery/dao"
	"go-lottery/dataSource"
	"go-lottery/models"
)

type CodeService interface {
	GetAll() []models.Code
	CountAll() int64
	Get(id int) *models.Code
	Delete(id int) error
	Update(data *models.Code, columns []string) error
	Insert(data *models.Code) error
	UpdateByCode(data *models.Code, columns []string) error
	NextUsingCode(giftId, codeId int) *models.Code
}

type codeService struct {
	dao *dao.CodeDao
}

func NewCodeService() CodeService {
	return &codeService{
		dao: dao.NewCodeDao(dataSource.NewMysqlMaster()),
	}
}

func (this *codeService) GetAll() []models.Code {
	return this.dao.GetAll()
}

func (this *codeService) CountAll() int64 {
	return this.dao.CountAll()
}

func (this *codeService) Get(id int) *models.Code {
	return this.dao.Get(id)
}

func (this *codeService) Delete(id int) error {
	return this.dao.Delete(id)
}

func (this *codeService) Update(data *models.Code, columns []string) error {
	return this.dao.Update(data, columns)
}

func (this *codeService) Insert(data *models.Code) error {
	return this.dao.Insert(data)
}

func (this *codeService) UpdateByCode(data *models.Code, columns []string) error {
	return this.dao.UpdateByCode(data, columns)
}

func (this *codeService) NextUsingCode(giftId, codeId int) *models.Code {
	return this.dao.NextUsingCode(giftId, codeId)
}
