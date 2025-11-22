package k8s

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/portforward"
	"k8s.io/client-go/transport/spdy"
)

func PortForward(restConfig *rest.Config, client *kubernetes.Clientset, podName string, localPort int, remotePort int) (chan struct{}, error) {
	req := client.CoreV1().RESTClient().Post().
		Resource("pods").
		Namespace("default").
		Name(podName).
		SubResource("portforward")

	transport, upgrader, _ := spdy.RoundTripperFor(restConfig)
	dialer := spdy.NewDialer(upgrader, &http.Client{Transport: transport}, "POST", req.URL())

	stopChan := make(chan struct{}, 1)
	readyChan := make(chan struct{})

	ports := []string{fmt.Sprintf("%d:%d", localPort, remotePort)}

	fw, err := portforward.New(
		dialer,
		ports,
		stopChan,
		readyChan,
		os.Stdout, // messages from forwarder
		os.Stderr,
	)
	if err != nil {
		return nil, err
	}

	go func() error {
		if err := fw.ForwardPorts(); err != nil {
			return err
		}

		return nil
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	go func() {
		<-sigChan
		close(stopChan)
	}()

	return readyChan, nil
}
