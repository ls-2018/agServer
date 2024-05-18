package main

import (
	"flag"
	"my.domain/guestbook/cmd/server"

	genericserver "k8s.io/apiserver/pkg/server"
	"k8s.io/component-base/logs"
	"k8s.io/klog/v2"
)

func main() {
	// 把cobra cmd做出来
	stopCh := genericserver.SetupSignalHandler()
	options := server.NewServerOptions()
	cmd := server.NewCommandStartServer(options, stopCh)
	cmd.Flags().AddGoFlagSet(flag.CommandLine)
	// 初始化以下log， 需要在parse flag之后
	logs.InitLogs()
	defer logs.FlushLogs()
	// 启动
	if err := cmd.Execute(); err != nil {
		klog.Fatal(err)
	}
}
