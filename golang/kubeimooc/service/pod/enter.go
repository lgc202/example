package pod

import "golang/kubeimooc/convert"

type PodServiceGroup struct {
	PodService
}

var podConvert = convert.ConvertGroupApp.PodConvert
