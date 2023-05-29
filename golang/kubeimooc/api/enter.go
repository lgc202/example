package api

import (
	"golang/kubeimooc/api/example"
	"golang/kubeimooc/api/k8s"
)

type ApiGroup struct {
	ExampleApiGroup example.ApiGroup
	K8SApiGroup     k8s.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
