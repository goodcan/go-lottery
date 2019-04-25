package services

import (
	"strconv"
	"strings"

	"go-lottery/dao"
	"go-lottery/dataSource"
	"go-lottery/models"
)

type GiftService interface {
	GetAll() []models.Gift
	CountAll() int64
	Get(id int) *models.Gift
	Delete(id int) error
	Update(data *models.Gift, columns []string) error
	Insert(data *models.Gift) error
	GetAllUse() []models.GiftPrize
	DecrLeftNum(id, num int) (int64, error)
	IncrLeftNum(id, num int) (int64, error)
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

func (this *giftService) GetAllUse() []models.GiftPrize {
	dataList := make([]models.Gift, 0)
	dataList = this.dao.GetAllUse()

	if dataList != nil {
		gifts := make([]models.GiftPrize, 0)

		for _, gift := range dataList {
			codes := strings.Split(gift.PrizeCode, "-")
			if len(codes) == 2 {
				a, e1 := strconv.Atoi(codes[0])
				b, e2 := strconv.Atoi(codes[1])
				if e1 == nil && e2 == nil && b >= a && a >= 0 && b <= 10000 {

					data := models.GiftPrize{
						Id:           gift.Id,
						Title:        gift.Title,
						PrizeNum:     gift.PrizeNum,
						LeftNum:      gift.LeftNum,
						PrizeCodeA:   a,
						PrizeCodeB:   b,
						Img:          gift.Img,
						DisplayOrder: gift.DisplayOrder,
						Gtype:        gift.Gtype,
						Gdata:        gift.Gdata,
					}

					gifts = append(gifts, data)
				}
			}
		}

		return gifts

	} else {
		return []models.GiftPrize{}
	}
}

func (this *giftService) DecrLeftNum(id, num int) (int64, error) {
	return this.dao.DecrLeftNum(id, num)
}

func (this *giftService) IncrLeftNum(id, num int) (int64, error) {
	return this.dao.IncrLeftNum(id, num)
}
