package initialize

import (
	"fmt"
	"go-service-api/common"
	"go-service-api/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Mysql() {
	dbConfig := config.DB
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Name,
		dbConfig.Charset)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,   // string 类型字段的默认长度
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败:" + err.Error())
	}
	common.DB = db
}
