package initialize

import (
	"golang/kubeimooc/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	r := gin.Default()
	examplGroup := router.RouterGroupApp.ExampleRouterGroup
	k8sGroup := router.RouterGroupApp.K8SRouterGroup
	examplGroup.InitExample(r)
	k8sGroup.InitK8SRouter(r)
	return r
}
