package main

import (
	"fmt"
	genericapiserver "k8s.io/apiserver/pkg/server"
	"my.domain/guestbook/cmd/v2/options"
)

func main() {
	opts := options.NewOptions()
	fmt.Println(runCommand(opts, genericapiserver.SetupSignalHandler()))
}
func runCommand(o *options.Options, stopCh <-chan struct{}) error {

	config, err := o.ServerConfig()

	if err != nil {
		return err
	}

	s, err := config.Complete()

	if err != nil {
		return err
	}

	return s.RunUntil(stopCh)
}
