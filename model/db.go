package model

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/tzr2020/gin_blog/utils"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db  *gorm.DB
	err error
)

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	)
	db, err = gorm.Open(utils.Db, dsn)
	err = db.DB().Ping()
	if err != nil {
		log.Println("连接数据库失败:", err)
		return
	}

	// 设置连接池中的最大闲置连接数
	db.DB().SetMaxIdleConns(10)
	// 设置数据库的最大连接数量
	db.DB().SetMaxOpenConns(100)
	// 设置连接的最大可复用时间
	db.DB().SetConnMaxLifetime(10 * time.Second)

	// 禁用默认表名的复数形式
	db.SingularTable(true)
	// GORM自动迁移模型
	db.AutoMigrate(&User{}, &Article{}, &Category{})
}
