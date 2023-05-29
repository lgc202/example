package initialize

import (
	"golang/kubeimooc/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	r := gin.Default()
	examplGroup := router.RouterGroupApp.ExampleRouterGroup
	examplGroup.InitExample(r)
	return r
}
