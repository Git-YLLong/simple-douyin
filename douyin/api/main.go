package main

import (
	"net/http"

	"github.com/Git-YLLong/simple-douyin/douyin/api/handlers"
	"github.com/Git-YLLong/simple-douyin/douyin/api/rpc"
	"github.com/Git-YLLong/simple-douyin/pkg/constants"
	"github.com/Git-YLLong/simple-douyin/pkg/tracer"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
)

func Init() {
	tracer.InitJaeger(constants.ApiServiceName)
	rpc.InitRPC()
}

// 校验token中间件

func initRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public") // 不能动

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", handlers.Feed)
	apiRouter.POST("/user/register/", handlers.Register)
	apiRouter.POST("/user/login/", handlers.Login)
	apiRouter.GET("/user/", handlers.GetUser)
	apiRouter.POST("/publish/action/", handlers.PublishVideo)
	apiRouter.GET("/publish/list/", handlers.GetPublishedList)
	// extra apis - I
	apiRouter.POST("/favorite/action/", handlers.FavoriteAction)
	apiRouter.GET("/favorite/list/", handlers.GetFavoriteList)
	// apiRouter.POST("/comment/action/", controller.CommentAction)
	// apiRouter.GET("/comment/list/", controller.CommentList)

	// extra apis - II
	// apiRouter.POST("/relation/action/", controller.RelationAction)
	// apiRouter.GET("/relation/follow/list/", controller.FollowList)
	// apiRouter.GET("/relation/follower/list/", controller.FollowerList)
}

func main() {
	Init()
	r := gin.Default()

	initRouter(r)

	if err := http.ListenAndServe(":8080", r); err != nil {
		klog.Fatal(err)
	}
}
