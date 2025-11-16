package k8s

import (
	"context"
	"log"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func CreatePod(client kubernetes.Interface, name string) error {
	log.Println("pods/CreatePod")
	podDefintion := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"app":         "devpod",
				"devpod-name": name,
			},
			Namespace: "default",
			Name:      name,
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:    "main",
					Image:   "busybox:latest",
					Command: []string{"sleep", "3600"},
				},
			},
		},
	}

	client, err := NewClientSet()
	if err != nil {
		return err
	}

	podsClient := client.CoreV1().Pods("default")

	// create a new pod
	_, err = podsClient.Create(context.TODO(), podDefintion, metav1.CreateOptions{})
	if err != nil {
		return err
	}

	return nil
}
