package client

import (
	"k8s.io/client-go/kubernetes"
	"kube-assistant/cache"
)

var clientCache = cache.New[*kubernetes.Clientset]()

func GetCache() cache.Cache[*kubernetes.Clientset] {
	return clientCache
}
