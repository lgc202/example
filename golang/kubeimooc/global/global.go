package global

import (
	"kubeimooc/config"

	"k8s.io/client-go/kubernetes"
)

var (
	CONF          config.Server
	KubeConfigSet *kubernetes.Clientset
)
