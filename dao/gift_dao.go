package dao

import (
	"log"

	"github.com/go-xorm/xorm"

	"go-lottery/comm"
	"go-lottery/models"
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

// 获取到当前可以获取的奖品列表
// 有奖品限定，状态正常，时间期间内
// gtype 倒序，display_order 正序
func (this *GiftDao) GetAllUse() []models.Gift {
	now := comm.NowTime()
	dataList := make([]models.Gift, 0)

	err := this.engine.
		Cols("id",
			"title",
			"prize_num",
			"left_num",
			"prize_code",
			"prize_time",
			"img",
			"display_order",
			"gtype",
		).
		Desc("gtype").
		Asc("display_order").
		Where("prize_num>=?", 0).
		Where("sys_status=?", 0).
		Where("time_begin<=?", now).
		Where("time_end>=?", now).
		Find(&dataList)

	if err != nil {
		log.Println("gift_dao.GetAllUse error=", err)
	}

	return dataList
}

func (this *GiftDao) DecrLeftNum(id, num int) (int64, error) {
	r, err := this.engine.
		Id(id).
		Decr("left_num", num).
		Where("left_num>=?", num).
		Update(&models.Gift{Id: id})

	return r, err
}

func (this *GiftDao) IncrLeftNum(id, num int) (int64, error) {
	r, err := this.engine.
		Id(id).
		Incr("left_num", num).
		Update(&models.Gift{Id: id})

	return r, err
}
