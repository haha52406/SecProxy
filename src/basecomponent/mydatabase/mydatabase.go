package mydatabase

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type MysqlClient struct {
	*gorm.DB
}

var MysqlConf = &MysqlClient{}

func InitMysql() error {
	client, err := gorm.Open("mysql", "root:root@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return err
	}
	defer client.Close()
	MysqlConf.DB = client
	return nil
}
