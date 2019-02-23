package models

import "time"

type Code struct {
	Id         int       `xorm:"INT pk autoincr 'id'"`
	GiftId     int       `xorm:"INT  'gift_id'"`
	Code       string    `xorm:"VARCHAR(255) 'code'"`
	SysCreated time.Time `xorm:"DATETIME DEFAULT CURRENT_TIMESTAMP created 'sys_created'"`
	SysUpdated time.Time `xorm:"DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP updated 'sys_updated'"`
	SysStatus  int       `xorm:"SMALLINT 'sys_status'"`
}

func (this *Code) TableName() string {
	return "code"
}
