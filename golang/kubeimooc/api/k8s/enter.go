package k8s

import "kubeimooc/service"

type ApiGroup struct {
	PodApi
}

var podService = service.ServiceGroupApp.PodServiceGroup.PodService
