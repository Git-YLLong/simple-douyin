package db

import (
	"github.com/Git-YLLong/simple-douyin/pkg/constants"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
)

var DB *gorm.DB

func Conn() *gorm.DB {
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

func InitUser() {
	var err error
	DB = Conn()

	if err = DB.Use(gormopentracing.New()); err != nil {
		panic(err)
	}

	m := DB.Migrator()
	if m.HasTable(&User{}) {
		return
	}
	if err = m.CreateTable(&User{}); err != nil {
		panic(err)
	}
}

func InitUserInfo() {
	var err error
	DB = Conn()

	if err = DB.Use(gormopentracing.New()); err != nil {
		panic(err)
	}

	m := DB.Migrator()
	if m.HasTable(&UserInfo{}) {
		return
	}
	if err = m.CreateTable(&UserInfo{}); err != nil {
		panic(err)
	}
}

func InitVideo() {
	var err error
	DB = Conn()

	if err = DB.Use(gormopentracing.New()); err != nil {
		panic(err)
	}

	m := DB.Migrator()
	if m.HasTable(&Video{}) {
		return
	}
	if err = m.CreateTable(&Video{}); err != nil {
		panic(err)
	}
}
