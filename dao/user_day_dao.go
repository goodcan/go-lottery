package dao

import (
	"log"

	"github.com/go-xorm/xorm"

	"go-lottery/models"
)

type UserDayDao struct {
	engine *xorm.Engine
}

func NewUserDayDao(engine *xorm.Engine) *UserDayDao {
	return &UserDayDao{engine: engine}
}

func (this *UserDayDao) Get(id int) *models.UserDay {
	data := &models.UserDay{Id: id}

	ok, err := this.engine.Get(data)

	if ok && err == nil {
		return data
	} else {
		return nil
	}

}

func (this *UserDayDao) GetAll() []models.UserDay {
	dataList := make([]models.UserDay, 0)

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

func (this *UserDayDao) CountAll() int64 {
	num, err := this.engine.Count(&models.UserDay{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

// 软删除
func (this *UserDayDao) Delete(id int) error {
	data := &models.UserDay{Id: id, SysStatus: 1}
	_, err := this.engine.Id(data.Id).Update(data)
	return err
}

func (this *UserDayDao) Update(data *models.UserDay, columns []string) error {
	_, err := this.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (this *UserDayDao) Insert(data *models.UserDay) error {
	_, err := this.engine.Insert(data)
	return err
}

func (this *UserDayDao) GetByUid(uid int) *models.UserDay {
	dataList := make([]models.UserDay, 0)
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

func (this *UserDayDao) Search(uid int, day string) *models.UserDay {
	dataList := make([]models.UserDay, 0)
	err := this.engine.Where("uid=?", uid).
		Where("day=", day).
		Limit(1).
		Find(&dataList)
	if err != nil || len(dataList) == 1 {
		return nil
	} else {
		return &dataList[0]
	}
}
