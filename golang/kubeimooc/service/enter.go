package service

import "golang/kubeimooc/service/pod"

type ServiceGroup struct {
	PodServiceGroup pod.PodServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
