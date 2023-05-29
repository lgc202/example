package convert

import "golang/kubeimooc/convert/pod"

type ConvertGroup struct {
	PodConvert pod.PodConvertGroup
}

var ConvertGroupApp = new(ConvertGroup)
