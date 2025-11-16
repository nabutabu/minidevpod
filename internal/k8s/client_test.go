package k8s

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestNewClientSet(t *testing.T) {
	clientSet, err := NewClientSet()
	if err != nil {
		t.Fatal(err) // Stops the test and marks it as failed
	}
	assert.NotNil(t, clientSet, "ClientSet is nil")

	podsClient := clientSet.CoreV1().Pods("default")

	list, err := podsClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		t.Fatal("could not list pods", err)
	}

	t.Logf("Listing pods...\n")
	for _, p := range list.Items {
		t.Logf(" * %s (Status: %s)\n", p.Name, p.Status.Phase)
	}
}
