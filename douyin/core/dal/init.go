package dal

import "github.com/Git-YLLong/simple-douyin/douyin/core/dal/db"

// Init init dal
func Init() {
	db.InitUser()     // user table init
	db.InitUserInfo() // userInfo table init
	db.InitVideo()
}
