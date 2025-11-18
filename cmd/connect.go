package cmd

import (
	"context"
	"errors"
	"miniDevPod/internal/k8s"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Connect(name string) error {
	clientSet, err := k8s.NewClientSet()
	if err != nil {
		return err
	}

	// check if pod exists and Status == Running
	pod, _ := clientSet.CoreV1().Pods("default").Get(context.TODO(), name, metav1.GetOptions{})
	if pod == nil || pod.Status.Phase != v1.PodRunning {
		return errors.New("Pod does not exist or is not running")
	}

	// k8s.ExecInPod(clientSet, name)
	config, err := k8s.GetConfig()
	if err != nil {
		return err
	}

	err = k8s.ExecInPod(clientSet, config, name)
	if err != nil {
		return err
	}

	return nil
}
