package k8s

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	// v1 "k8s.io/api/core/v1"
	// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func TestCreateService(t *testing.T) {
	clientSet := fake.NewSimpleClientset()
	podName := "test-pod"
	assert.Nil(t, CreatePod(clientSet, podName, "", "", "", ""))
	port, err := CreateService(clientSet, podName)
	assert.Nil(t, err)
	log.Println(port)
}
