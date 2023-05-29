package global

import (
	"golang/kubeimooc/config"

	"k8s.io/client-go/kubernetes"
)

var (
	CONF          config.Server
	KubeConfigSet *kubernetes.Clientset
)
