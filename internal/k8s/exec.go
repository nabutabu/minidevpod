package k8s

import (
	"context"
	"log"

	"os"
	"strings"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
)

func ExecInPod(client *kubernetes.Clientset, restConfig *rest.Config, podName string) error {
	log.Println("k8s/ExecInPod")

	scheme := runtime.NewScheme()
	if err := v1.AddToScheme(scheme); err != nil {
		return err
	}
	parameterCodec := runtime.NewParameterCodec(scheme)

	req := client.CoreV1().RESTClient().Post().
		Resource("pods").
		Name(podName).
		Namespace("default").
		SubResource("exec").
		VersionedParams(&corev1.PodExecOptions{
			Command: []string{"/bin/bash"},
			Stdin:   true,
			Stdout:  true,
			Stderr:  true,
			TTY:     true,
		}, parameterCodec)

	exec, err := remotecommand.NewSPDYExecutor(restConfig, "POST", req.URL())

	err = exec.StreamWithContext(context.TODO(), remotecommand.StreamOptions{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Tty:    true,
	})

	if err != nil {
		if strings.Contains(err.Error(), "command terminated with exit code") {
			return err
		} else {
			return err
		}
	}

	return nil
}
