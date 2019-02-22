package models

import "time"

type Result struct {
	Id         int       `xorm:"INT pk autoincr 'id'"`
	GiftId     int       `xorm:"INT 'gift_id'"`
	GiftName   string    `xorm:"VARCHAR(255) 'gift_name'"`
	GiftType   int       `xorm:"INT 'gift_type'"`
	Uid        int       `xorm:"INT 'uid'"`
	Username   string    `xorm:"VARCHAR(50) 'username'"`
	PrizeCode  int       `xorm:"INT 'prize_code'"`
	GiftData   string    `xorm:"VARCHAR(50) 'gift_data'"`
	SysCreated time.Time `xorm:"DATETIME DEFAULT CURRENT_TIMESTAMP created 'sys_created'"`
	SysStatus  int       `xorm:"SMALLINT 'sys_status'"`
	SysIP      string    `xorm:"VARCHAR(50) 'sys_ip'"`
}

func (this *Result) TableName() string {
	return "result"
}
