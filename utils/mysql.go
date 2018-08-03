package utils

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

func init() {
	fmt.Println("Mysql Init:", InitMysql())
}

var E *xorm.Engine

func InitMysql() (err error) {
	dbConf := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8",
		Conf.Mysql.UserName,
		Conf.Mysql.Password,
		Conf.Mysql.Host,
		Conf.Mysql.Port,
		Conf.Mysql.DBName,
	)
	E, err = xorm.NewEngine("mysql", dbConf)
	E.ShowSQL(true)
	E.SetMapper(core.GonicMapper{})
	E.Logger().SetLevel(core.LOG_DEBUG)
	E.SetMaxIdleConns(Conf.Mysql.IdleNu)
	E.SetMaxOpenConns(Conf.Mysql.MaxActive)
	return
}
