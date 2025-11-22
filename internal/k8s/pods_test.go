package k8s

import (
	"testing"

	"github.com/stretchr/testify/assert"
	// v1 "k8s.io/api/core/v1"
	// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func TestCreateNewPod(t *testing.T) {
	clientSet := fake.NewSimpleClientset()
	assert.Nil(t, CreatePod(clientSet, "test-pod-2", "https://github.com/nabutabu/httpfromtcp", "using_net", "500m", "1Gi"))
}
