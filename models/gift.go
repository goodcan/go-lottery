package models

type Gift struct {
	Id        int    `xorm:"int pk autoincr 'id'"`
	Title     string `xorm:"varchar(255) notnull default '' 'title'"`
	PrizeNum  int    `xorm:"int notnull default 0 'prize_num'"`
	LeftNum   int    `xorm:"int notnull default 0 'left_num'"`
	PrizeCode string `xorm:"varchar(50) notnull 'prize_code'"`
}
