package k8s

import (
	"golang/kubeimooc/api"

	"github.com/gin-gonic/gin"
)

type K8sRouter struct {
}

func (*K8sRouter) InitK8SRouter(r *gin.Engine) {
	group := r.Group("/k8s")
	apiGroup := api.ApiGroupApp.K8SApiGroup
	group.GET("/pod/:namespace", apiGroup.GetPodListOrDetail)
}
