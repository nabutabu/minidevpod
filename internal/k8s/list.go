package k8s

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ListDevPods(client *kubernetes.Clientset) ([]corev1.Pod, error) {
	// Create a pod interface for the given namespace
	podInterface := client.CoreV1().Pods("default")

	// List the pods in the given namespace, with the label selector filter
	podList, err := podInterface.List(context.TODO(), metav1.ListOptions{
		LabelSelector: "app=devpod",
	})
	if err != nil {
		return nil, err
	}

	return podList.Items, nil
}
