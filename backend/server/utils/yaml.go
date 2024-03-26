package utils

import (
	"k8s.io/apimachinery/pkg/util/yaml"
)

type Kubeconfig struct {
	Clusters []KubeconfigCluster `json:"clusters"`
}

type KubeconfigCluster struct {
	Name string `json:"name"`
}

func UnmarshalKubeconfig(content []byte) (*Kubeconfig, error) {
	var kubeconfig *Kubeconfig

	err := yaml.Unmarshal(content, kubeconfig)
	if err != nil {
		return nil, err
	}
	return kubeconfig, nil
}
