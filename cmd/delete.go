package cmd

import (
	"miniDevPod/internal/k8s"
)

func Delete(name string) error {
	clientSet, err := k8s.NewClientSet()
	if err != nil {
		return err
	}

	err = k8s.DeleteDevPod(clientSet, name)
	if err != nil {
		return err
	}

	return nil
}
