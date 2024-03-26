package service

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"kube-assistant/cache"
)

func load(kubeconfig string, cache cache.Cache[interface{}]) {
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientSet
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	cache.Set(kubeconfig, clientSet)
}
