package mydatabase

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type MysqlClient struct {
	db *gorm.DB
}

var MysqlConf = &MysqlClient{}

type Admin struct {
	Id       int
	Username string
	Password string
	Nickname string
	Mobile   string
	IsUsed   int
}

func InitMysql() error {
	client, err := gorm.Open("mysql", "root:yangyulong@/go_gin_api?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return err
	}
	defer client.Close()
	MysqlConf.db = client
	//使用方法
	// admin := &Admin{}
	// MysqlConf.db.Select("*").Where("id=?", 1).Table("admin").Find(admin)
	// fmt.Println("admin:", admin)
	return nil
}
