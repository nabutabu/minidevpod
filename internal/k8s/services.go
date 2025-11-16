package k8s

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
)

func CreateService(client kubernetes.Interface, podName string) (int, error) {
	service := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-service",
			Namespace: "default",
			Labels: map[string]string{
				"app": "devpod",
			},
		},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{
				"app":         "devpod",
				"devpod-name": podName,
			},
			Ports: []corev1.ServicePort{
				{
					Name:       "ssh",
					Protocol:   corev1.ProtocolTCP,
					Port:       22,
					TargetPort: intstr.FromInt(22), // target port on pod
				},
			},
			Type: corev1.ServiceTypeNodePort,
		},
	}

	svc, err := client.CoreV1().Services("default").Create(context.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		return 0, err
	}

	return int(svc.Spec.Ports[0].NodePort), nil
}
