package utils

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	myConfig "onemore-service-iris/config"
)

func init() {
	InitMysql()
}

var OE *xorm.Engine

func InitMysql() {
	dbConf := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8",
		myConfig.Conf.Mysql.UserName,
		myConfig.Conf.Mysql.Password,
		myConfig.Conf.Mysql.Host,
		myConfig.Conf.Mysql.Port,
		myConfig.Conf.Mysql.DBName,
	)
	OE, _ = xorm.NewEngine("mysql", dbConf)
	OE.SetMaxIdleConns(100)
	OE.SetMaxOpenConns(1000)
}
