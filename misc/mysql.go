package misc

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"onemore-service-iris/config"
)

func init() {
	fmt.Println("Mysql Init:", InitMysql())
}

var E *xorm.Engine

func InitMysql() (err error) {
	dbConf := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8",
		config.Conf.Mysql.UserName,
		config.Conf.Mysql.Password,
		config.Conf.Mysql.Host,
		config.Conf.Mysql.Port,
		config.Conf.Mysql.DBName,
	)
	E, err = xorm.NewEngine("mysql", dbConf)
	E.ShowSQL(true)
	E.SetMapper(core.GonicMapper{})
	E.Logger().SetLevel(core.LOG_DEBUG)
	E.SetMaxIdleConns(config.Conf.Mysql.IdleNu)
	E.SetMaxOpenConns(config.Conf.Mysql.MaxActive)
	return
}
