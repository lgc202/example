package k8s

import (
	"kubeimooc/response"

	"github.com/gin-gonic/gin"
)

type PodApi struct {
}

func (*PodApi) GetPodListOrDetail(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Query("name")
	keyword := c.Query("keyword")
	if name != "" {
		// TODO
		// detail, err := podService.GetPodDetail(namespace, name)
		// if err != nil {
		// 	response.FailWithMessage(c, err.Error())
		// 	return
		// }
		// response.SuccessWithDetailed(c, "获取Pod详情成功", detail)
	} else {
		err, items := podService.GetPodList(namespace, keyword, c.Query("nodeName"))
		if err != nil {
			response.FailWithMessage(c, err.Error())
			return
		}
		response.SuccessWithDetailed(c, "获取Pod列表成功", items)
	}
}
