package cmd

import (
	"context"
	"log"
	"miniDevPod/internal/k8s"
	"time"

	"github.com/google/uuid"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func generateName() string {
	return uuid.NewString()
}

func CreatePod(name string) {
	if name == "" {
		name = generateName()
	}

	name = "devpod-" + name

	// at this point we should have a unique name or a custom name for a new pod to create
	// get a new clientset
	clientSet, err := k8s.NewClientSet()
	if err != nil {
		log.Fatal(err)
	}

	// call pods.create
	log.Printf("Creating pod: %s\n", name)
	err = k8s.CreatePod(clientSet, name)
	if err != nil {
		log.Fatal(err)
	}

	// poll until ready
	pod, _ := clientSet.CoreV1().Pods("default").Get(context.TODO(), name, metav1.GetOptions{})
	for pod.Status.Phase != v1.PodRunning {
		pod, _ = clientSet.CoreV1().Pods("default").Get(context.TODO(), name, metav1.GetOptions{})
		time.Sleep(500 * time.Millisecond)
	}

	log.Println("DevPod is running!")

	_, err = k8s.CreateService(clientSet, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("SSH into pod using:\ngo run main.go connect %s\n", name)
}
