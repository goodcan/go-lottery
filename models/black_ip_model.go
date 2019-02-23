package models

import "time"

type BlackIp struct {
	Id         int       `xorm:"INT pk autoincr 'id'"`
	Ip         string    `xorm:"VARCHAR(5) 'ip'"`
	BlackTime  time.Time `xorm:"DATETIME 'black_time'"`
	SysCreated time.Time `xorm:"DATETIME DEFAULT CURRENT_TIMESTAMP created 'sys_created'"`
	SysUpdated time.Time `xorm:"DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP updated 'sys_updated'"`
	SysStatus  int       `xorm:"SMALLINT 'sys_status'"`
}

func (this *BlackIp) TableName() string {
	return "black_ip"
}
