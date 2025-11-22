package k8s

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func DeleteDevPod(client *kubernetes.Clientset, name string) error {
	deletePolicy := metav1.DeletePropagationBackground // Or Foreground, Orphan
	deleteOptions := metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}

	err := client.CoreV1().Pods("default").Delete(context.TODO(), name, deleteOptions)
	if err != nil {
		return err
	}

	serviceName := name + "-service" // Replace with the actual service name

	// Delete the Service
	err = client.CoreV1().Services("default").Delete(context.TODO(), serviceName, metav1.DeleteOptions{})
	if err != nil {
		return err
	}

	return nil
}
