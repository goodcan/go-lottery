package dao

import (
	"log"

	"github.com/go-xorm/xorm"

	"../models"
)

type GiftDao struct {
	engine *xorm.Engine
}

func NewGiftDao(engine *xorm.Engine) *GiftDao {
	return &GiftDao{engine: engine}
}

func (this *GiftDao) Get(id int) *models.Gift {
	data := &models.Gift{Id: id}

	ok, err := this.engine.Get(data)

	if ok && err == nil {
		return data
	} else {
		return nil
	}

}

func (this *GiftDao) GetAll() []models.Gift {
	dataList := make([]models.Gift, 0)

	err := this.engine.
		Asc("sys_status").
		Asc("display_order").
		Find(&dataList)

	if err != nil {
		log.Println("gift_dao.GetAll error=", err)
		return dataList
	} else {
		return dataList
	}
}

func (this *GiftDao) CountAll() int64 {
	num, err := this.engine.Count(&models.Gift{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

// 软删除
func (this *GiftDao) Delete(id int) error {
	data := &models.Gift{Id: id, SysStatus: 1}
	_, err := this.engine.Id(data.Id).Update(data)
	return err
}

func (this *GiftDao) Update(data *models.Gift, columns []string) error {
	_, err := this.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (this *GiftDao) Insert(data *models.Gift) error {
	_, err := this.engine.Insert(data)
	return err
}
