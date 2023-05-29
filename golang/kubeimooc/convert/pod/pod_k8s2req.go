package pod

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	pod_res "golang/kubeimooc/model/pod/response"
)

type K8s2ReqConvert struct {
}

func (*K8s2ReqConvert) PodK8s2ItemRes(pod corev1.Pod) pod_res.PodListItem {

	var totalC, readyC, restartC int32
	for _, containerStatus := range pod.Status.ContainerStatuses {
		if containerStatus.Ready {
			readyC++
		}
		restartC += containerStatus.RestartCount
		totalC++
	}
	var podStatus string
	if pod.Status.Phase != "Running" {
		podStatus = "Error"
	} else {
		podStatus = "Running"
	}
	return pod_res.PodListItem{
		Name:     pod.Name,
		Ready:    fmt.Sprintf("%d/%d", readyC, totalC),
		Status:   podStatus,
		Restarts: restartC,
		Age:      pod.CreationTimestamp.Unix(),
		IP:       pod.Status.PodIP,
		Node:     pod.Spec.NodeName,
	}
}
