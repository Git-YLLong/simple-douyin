package main

import (
	"net/http"

	"github.com/Git-YLLong/simple-douyin/controller"
	"github.com/Git-YLLong/simple-douyin/douyin/api/handlers"
	"github.com/Git-YLLong/simple-douyin/douyin/api/rpc"
	"github.com/Git-YLLong/simple-douyin/pkg/constants"
	"github.com/Git-YLLong/simple-douyin/pkg/jwt"
	"github.com/Git-YLLong/simple-douyin/pkg/tracer"
	"github.com/cloudwego/kitex/pkg/klog"

	// "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Init() {
	tracer.InitJaeger(constants.ApiServiceName)
	rpc.InitRPC()
}

// 校验token中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取token
		tokeString := c.GetHeader("token")
		// fmt.Println(tokeString, "当前token")
		if tokeString == "" {
			c.JSON(http.StatusOK, gin.H{
				"code":    1005,
				"message": "必须传递token",
			})
			c.Abort()
			return
		}
		claims, err := jwt.ParseJwtToken(tokeString)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    1005,
				"message": "token解析错误",
			})
			c.Abort()
			return
		}
		// 从token中解析出来的数据挂载到上下文上,方便后面的控制器使用
		c.Set("userId", claims.UID)
		// c.Set("userName", claims.Username)
		c.Next()
	}
}

func initRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", controller.Feed)
	apiRouter.POST("/user/register/", handlers.Register)
	apiRouter.POST("/user/login/", controller.Login)
	apiRouter.GET("/user/", controller.UserInfo)

	// apiRouter.Use(AuthMiddleware()) // 验证token令牌
	// {

	apiRouter.POST("/publish/action/", controller.Publish)
	apiRouter.GET("/publish/list/", controller.PublishList)

	// extra apis - I
	apiRouter.POST("/favorite/action/", controller.FavoriteAction)
	apiRouter.GET("/favorite/list/", controller.FavoriteList)
	apiRouter.POST("/comment/action/", controller.CommentAction)
	apiRouter.GET("/comment/list/", controller.CommentList)

	// extra apis - II
	apiRouter.POST("/relation/action/", controller.RelationAction)
	apiRouter.GET("/relation/follow/list/", controller.FollowList)
	apiRouter.GET("/relation/follower/list/", controller.FollowerList)
	// }
}

func main() {
	Init()
	r := gin.Default()

	initRouter(r)

	if err := http.ListenAndServe(":8080", r); err != nil {
		klog.Fatal(err)
	}
}
