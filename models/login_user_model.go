package models

import "time"

type LoginUser struct {
	Uid      int
	Username string
	Now      time.Time
	Ip       string
	Sign     string
}
