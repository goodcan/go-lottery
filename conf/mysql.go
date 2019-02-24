package conf

const MysqlDriverName = "mysql"

type MysqlConfig struct {
	Host      string
	Port      int
	User      string
	Pwd       string
	Database  string
	IsRunning bool
}

var MysqlMasterList = []MysqlConfig{
	{
		Host:      "127.0.0.1",
		Port:      3306,
		User:      "root",
		Pwd:       "peipeiyun071",
		Database:  "go-lottery",
		IsRunning: true,
	},
}

var MysqlMaster = MysqlMasterList[0]
