package glb

/**
	数据库全局对象
 */

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"alm-api/config"
)

var (
	DB  *gorm.DB
	err error
)

func init() {
	if DB, err = gorm.Open(conf.DBDialect, conf.DBConnStr); err != nil {
		log.Fatal(err.Error())
	}
	if DB.DB().Ping() != nil {
		log.Fatal(err.Error())
	}
	DB.DB().SetMaxOpenConns(conf.DBMaxOpenConns)
	DB.DB().SetMaxIdleConns(conf.DBMaxIdleConns)
}
