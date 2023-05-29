package k8s

import "golang/kubeimooc/service"

type ApiGroup struct {
	PodApi
}

var podService = service.ServiceGroupApp.PodServiceGroup.PodService
