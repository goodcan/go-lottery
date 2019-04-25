package dao

import (
	"log"

	"github.com/go-xorm/xorm"

	"go-lottery/models"
)

type BlackUserDao struct {
	engine *xorm.Engine
}

func NewBlackUserDao(engine *xorm.Engine) *BlackUserDao {
	return &BlackUserDao{engine: engine}
}

func (this *BlackUserDao) Get(id int) *models.BlackUser {
	data := &models.BlackUser{Id: id}

	ok, err := this.engine.Get(data)

	if ok && err == nil {
		return data
	} else {
		return nil
	}

}

func (this *BlackUserDao) GetAll() []models.BlackUser {
	dataList := make([]models.BlackUser, 0)

	err := this.engine.
		Desc("id").
		Find(&dataList)

	if err != nil {
		log.Println("black_user_dao.GetAll error=", err)
		return dataList
	} else {
		return dataList
	}
}

func (this *BlackUserDao) CountAll() int64 {
	num, err := this.engine.Count(&models.BlackUser{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

// 软删除
func (this *BlackUserDao) Delete(id int) error {
	data := &models.BlackUser{Id: id, SysStatus: 1}
	_, err := this.engine.Id(data.Id).Update(data)
	return err
}

func (this *BlackUserDao) Update(data *models.BlackUser, columns []string) error {
	_, err := this.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (this *BlackUserDao) Insert(data *models.BlackUser) error {
	_, err := this.engine.Insert(data)
	return err
}

func (this *BlackUserDao) GetByUid(uid int) *models.BlackUser {
	dataList := make([]models.BlackUser, 0)
	err := this.engine.Where("uid=?", uid).
		Desc("id").
		Limit(1).
		Find(&dataList)

	if err != nil || len(dataList) <= 1 {
		return nil
	} else {
		return &dataList[0]
	}
}
