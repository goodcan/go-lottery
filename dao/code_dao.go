package dao

import (
	"log"

	"github.com/go-xorm/xorm"

	"../models"
)

type CodeDao struct {
	engine *xorm.Engine
}

func NewCodeDao(engine *xorm.Engine) *CodeDao {
	return &CodeDao{engine: engine}
}

func (this *CodeDao) Get(id int) *models.Code {
	data := &models.Code{Id: id}

	ok, err := this.engine.Get(data)

	if ok && err == nil {
		return data
	} else {
		return nil
	}

}

func (this *CodeDao) GetAll() []models.Code {
	dataList := make([]models.Code, 0)

	err := this.engine.
		Desc("id").
		Find(&dataList)

	if err != nil {
		log.Println("code_dao.GetAll error=", err)
		return dataList
	} else {
		return dataList
	}
}

func (this *CodeDao) CountAll() int64 {
	num, err := this.engine.Count(&models.Code{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

// 软删除
func (this *CodeDao) Delete(id int) error {
	data := &models.Code{Id: id, SysStatus: 1}
	_, err := this.engine.Id(data.Id).Update(data)
	return err
}

func (this *CodeDao) Update(data *models.Code, columns []string) error {
	_, err := this.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (this *CodeDao) Insert(data *models.Code) error {
	_, err := this.engine.Insert(data)
	return err
}
