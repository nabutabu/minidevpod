package k8s

import (
	"context"
	"log"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func CreatePod(client kubernetes.Interface, name string, repo string, branch string, cpuRequest string, memRequest string) error {
	log.Println("pods/CreatePod")
	podDefintion := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"app":         "devpod",
				"devpod-name": name,
				"created-by":  "mini-devpod",
			},
			Namespace: "default",
			Name:      name,
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:            "main",
					Image:           "devpod-base",
					ImagePullPolicy: "Never",
					VolumeMounts: []corev1.VolumeMount{
						{
							Name:      "workspace",
							MountPath: "/workspace",
						},
					},
					WorkingDir: "/workspace",
					Resources: corev1.ResourceRequirements{
						Requests: corev1.ResourceList{
							corev1.ResourceCPU:    resource.MustParse(cpuRequest),
							corev1.ResourceMemory: resource.MustParse(memRequest),
						},
						Limits: corev1.ResourceList{
							corev1.ResourceCPU:    resource.MustParse(cpuRequest),
							corev1.ResourceMemory: resource.MustParse(memRequest),
						},
					},
				},
			},
			Volumes: []corev1.Volume{
				{
					Name: "workspace",
					VolumeSource: corev1.VolumeSource{
						EmptyDir: &corev1.EmptyDirVolumeSource{},
					},
				},
			},
		},
	}

	if repo != "" {
		podDefintion.Spec.InitContainers = []corev1.Container{
			{
				Name:    "clone-repo",
				Image:   "alpine/git",
				Command: []string{"sh", "-c", "git clone --branch $BRANCH $REPO /workspace"},
				Env: []corev1.EnvVar{
					{Name: "REPO", Value: repo},
					{Name: "BRANCH", Value: branch},
				},
				VolumeMounts: []corev1.VolumeMount{
					{
						Name:      "workspace",
						MountPath: "/workspace",
					},
				},
			},
		}
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
