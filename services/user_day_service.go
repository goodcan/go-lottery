package services

import (
	"fmt"

	"go-lottery/comm"
	"go-lottery/dao"
	"go-lottery/dataSource"
	"go-lottery/models"
)

type UserDayService interface {
	GetAll() []models.UserDay
	CountAll() int64
	Get(id int) *models.UserDay
	Delete(id int) error
	Update(data *models.UserDay, columns []string) error
	Insert(data *models.UserDay) error
	GetByUid(uid int) *models.UserDay
	GetUserToday(uid int) *models.UserDay
}

type userDayService struct {
	dao *dao.UserDayDao
}

func NewUserDayService() UserDayService {
	return &userDayService{
		dao: dao.NewUserDayDao(dataSource.NewMysqlMaster()),
	}
}

func (this *userDayService) GetAll() []models.UserDay {
	return this.dao.GetAll()
}

func (this *userDayService) CountAll() int64 {
	return this.dao.CountAll()
}

func (this *userDayService) Get(id int) *models.UserDay {
	return this.dao.Get(id)
}

func (this *userDayService) Delete(id int) error {
	return this.dao.Delete(id)
}

func (this *userDayService) Update(data *models.UserDay, columns []string) error {
	return this.dao.Update(data, columns)
}

func (this *userDayService) Insert(data *models.UserDay) error {
	return this.dao.Insert(data)
}

func (this *userDayService) GetByUid(uid int) *models.UserDay {
	return this.dao.GetByUid(uid)
}

func (this *userDayService) GetUserToday(uid int) *models.UserDay {
	y, m, d := comm.NowTime().Date()
	strDay := fmt.Sprintf("%d%02d%02d", y, m, d)
	return this.dao.Search(uid, strDay)
}
