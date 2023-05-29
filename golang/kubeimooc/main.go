package main

import (
	"golang/kubeimooc/global"
	"golang/kubeimooc/initialize"
)

func main() {
	r := initialize.Routers()
	initialize.Viper()
	initialize.K8S()
	panic(r.Run(global.CONF.System.Addr))
}
