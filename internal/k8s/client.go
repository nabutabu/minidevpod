package k8s

import (
	"errors"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	//"k8s.io/client-go/util/homedir"
)

const (
	ConfigLocation = "/home/nabutabu/.kube/config"
)

func NewClientSet() (*kubernetes.Clientset, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		config, err = clientcmd.BuildConfigFromFlags("", ConfigLocation)
		if err != nil {
			return nil, errors.New("Could not build config")
		}
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, errors.New("Could not get client set")
	}

	return clientSet, nil
}
