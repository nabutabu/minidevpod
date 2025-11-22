package cmd

import (
	"log"
	"miniDevPod/internal/k8s"
)

func List() error {
	clientSet, err := k8s.NewClientSet()
	if err != nil {
		return err
	}

	podList, err := k8s.ListDevPods(clientSet)
	if err != nil {
		return err
	}

	for _, pod := range podList {
		log.Printf("\nName: %s\n\tStatus: %s\n\tCreation: %s\n", pod.Name, pod.Status.Phase, pod.CreationTimestamp)
	}
	return nil
}
