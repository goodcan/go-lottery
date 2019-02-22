package models

import "time"

type BlackUser struct {
	Id         int       `xorm:"INT pk autoincr 'id'"`
	Uid        int       `xorm:"INT 'uid'"`
	Username   string    `xorm:"VARCHAR(50) 'username'"`
	BlackTime  time.Time `xorm:"DATETIME 'black_time'"`
	RealName   string    `xorm:"VARCHAR(50) 'real_name'"`
	Mobile     string    `xorm:"VARCHAR(50) 'mobile'"`
	Address    string    `xorm:"VARCHAR(255) 'address'"`
	SysCreated time.Time `xorm:"DATETIME DEFAULT CURRENT_TIMESTAMP created 'sys_created'"`
	SysUpdated time.Time `xorm:"DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP updated 'sys_updated'"`
	SysIP      string    `xorm:"VARCHAR(50) 'sys_ip'"`
	SysStatus  int       `xorm:"SMALLINT 'sys_status'"`
}

func (this *BlackUser) TableName() string {
	return "black_user"
}
