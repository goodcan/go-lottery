package dao

import (
	"log"

	"github.com/go-xorm/xorm"

	"go-lottery/models"
)

type ResultDao struct {
	engine *xorm.Engine
}

func NewResultDao(engine *xorm.Engine) *ResultDao {
	return &ResultDao{engine: engine}
}

func (this *ResultDao) Get(id int) *models.Result {
	data := &models.Result{Id: id}

	ok, err := this.engine.Get(data)

	if ok && err == nil {
		return data
	} else {
		return nil
	}

}

func (this *ResultDao) GetAll() []models.Result {
	dataList := make([]models.Result, 0)

	err := this.engine.
		Desc("id").
		Find(&dataList)

	if err != nil {
		log.Println("Result_dao.GetAll error=", err)
		return dataList
	} else {
		return dataList
	}
}

func (this *ResultDao) CountAll() int64 {
	num, err := this.engine.Count(&models.Result{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

func (this *ResultDao) CountByGift(giftId int) int64 {
	num, err := this.engine.
		Where("gift_di=?", giftId).
		Count(&models.Result{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

func (this *ResultDao) CountByUser(uid int) int64 {
	num, err := this.engine.
		Where("uid=?", uid).
		Count(&models.Result{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

// 软删除
func (this *ResultDao) Delete(id int) error {
	data := &models.Result{Id: id, SysStatus: 1}
	_, err := this.engine.Id(data.Id).Update(data)
	return err
}

func (this *ResultDao) Update(data *models.Result, columns []string) error {
	_, err := this.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (this *ResultDao) Insert(data *models.Result) error {
	_, err := this.engine.Insert(data)
	return err
}

func (this *ResultDao) SearchByGift(giftId, page, size int) []models.Result {
	dataList := make([]models.Result, 0)

	err := this.engine.
		Where("gift_id=?", giftId).
		Limit(page, (page-1)*size).
		Desc("id").
		Find(&dataList)

	if err != nil {
		log.Println("Result_dao.GetAll error=", err)
		return dataList
	} else {
		return dataList
	}
}

func (this *ResultDao) SearchByUser(uid, page, size int) []models.Result {
	dataList := make([]models.Result, 0)

	err := this.engine.
		Where("uid=?", uid).
		Limit(page, (page-1)*size).
		Desc("id").
		Find(&dataList)

	if err != nil {
		log.Println("Result_dao.GetAll error=", err)
		return dataList
	} else {
		return dataList
	}
}
