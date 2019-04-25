package dataSource

import (
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"

	"go-lottery/conf"
)

var MysqlMasterInst *xorm.Engine
var mysqlLock sync.Mutex

func MysqlInstMaster() *xorm.Engine {
	if MysqlMasterInst != nil {
		return MysqlMasterInst
	}

	// 处理高并发时避免重复定义实例
	mysqlLock.Lock()
	defer mysqlLock.Unlock()

	if MysqlMasterInst != nil {
		return MysqlMasterInst
	}

	return NewMysqlMaster()
}

func NewMysqlMaster() *xorm.Engine {
	engine, err := xorm.NewEngine(
		conf.MysqlDriverName,
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
			conf.MysqlMaster.User,
			conf.MysqlMaster.Pwd,
			conf.MysqlMaster.Host,
			conf.MysqlMaster.Port,
			conf.MysqlMaster.Database,
		),
	)

	if err != nil {
		log.Fatal("db_helper.NewMysqlMaster NewEngine error ", err)
		return nil
	}

	// 本地调试打开 SQL 调试
	engine.ShowSQL(true)

	MysqlMasterInst = engine

	return MysqlMasterInst
}
