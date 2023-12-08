package main

import (
	"flag"
	"os"

	"github.com/Azure/kubelogin/internal/cmd"
	"github.com/spf13/pflag"
	klog "k8s.io/klog/v2"
)

func main() {
	klog.InitFlags(nil)
	pflag.CommandLine.AddGoFlag(flag.CommandLine.Lookup("v"))
	pflag.CommandLine.AddGoFlag(flag.CommandLine.Lookup("logtostderr"))
	_ = pflag.CommandLine.Set("logtostderr", "true")
	root := cmd.NewRootCmd()
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
