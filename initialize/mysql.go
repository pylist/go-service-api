package initialize

import (
	"fmt"
	"go-service-api/config"
	"go-service-api/global"
	"go-service-api/model"

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
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用数据库外键约束
	})
	if err != nil {
		panic("数据库连接失败:" + err.Error())
	}
	global.DB = db
	db.AutoMigrate(&model.User{}, &model.Menu{}, &model.Role{}, &model.Dept{})
}
