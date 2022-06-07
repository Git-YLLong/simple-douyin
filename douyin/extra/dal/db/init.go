package db

import (
	"github.com/Git-YLLong/simple-douyin/pkg/constants"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
)

var DB *gorm.DB

func ExtraConn() *gorm.DB {
	var err error
	DB, err = gorm.Open(mysql.Open(constants.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	return DB
}

func InitUserFavor() {
	var err error
	DB = ExtraConn()

	if err = DB.Use(gormopentracing.New()); err != nil {
		panic(err)
	}

	m := DB.Migrator()
	if m.HasTable(&UserFavor{}) {
		return
	}
	if err = m.CreateTable(&UserFavor{}); err != nil {
		panic(err)
	}
}
