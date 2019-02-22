package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"

	"../models"
)

func main() {
	engire, _ := xorm.NewEngine(
		"mysql",
		"root:peipeiyun071@tcp(127.0.0.1:3306)/go-lottery",
	)

	tables := []interface{}{
		models.Gift{},
		models.BlackIp{},
		models.BlackUser{},
		models.Code{},
		models.Result{},
		models.UserDay{},
	}

	for _, table := range tables {
		if exist, _ := engire.IsTableExist(table); exist {
			_ = engire.DropTables(table)
		}
		err := engire.CreateTables()
		if err != nil {
			fmt.Print(err)
		}
	}

}
