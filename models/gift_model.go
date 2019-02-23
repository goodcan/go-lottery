package models

import (
	"time"
)

type Gift struct {
	Id           int       `xorm:"INT pk autoincr 'id'"`
	Title        string    `xorm:"VARCHAR(255) 'title'"`
	PrizeNum     int       `xorm:"INT 'prize_num'"`
	LeftNum      int       `xorm:"INT 'left_num'"`
	PrizeCode    string    `xorm:"VARCHAR(50) 'prize_code'"`
	PrizeTime    time.Time `xorm:"DATETIME 'prize_time'"`
	Img          string    `xorm:"VARCHAR(255) 'img'"`
	DisplayOrder int       `xorm:"INT 'display_order'"`
	Gtype        int       `xorm:"INT 'gtype'"`
	Gdata        string    `xorm:"VARCHAR(255) 'gdata'"`
	TimeBegin    time.Time `xorm:"DATETIME 'time_begin'"`
	TimeEnd      time.Time `xorm:"DATETIME 'time_end'"`
	PrizeData    string    `xorm:"MEDIUMTEXT 'prize_data'"`
	PrizeBegin   time.Time `xorm:"DATETIME 'prize_begin'"`
	PrizeEnd     time.Time `xorm:"DATETIME 'prize_end'"`
	SysCreated   time.Time `xorm:"DATETIME DEFAULT CURRENT_TIMESTAMP created 'sys_created'"`
	SysUpdated   time.Time `xorm:"DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP updated 'sys_updated'"`
	SysStatus    int       `xorm:"SMALLINT 'sys_status'"`
	SysIP        string    `xorm:"VARCHAR(50) 'sys_ip'"`
}

func (this *Gift) TableName() string {
	return "gift"
}
