package cmd

import (
	"fmt"
	"log"
	"math/rand/v2"
	"miniDevPod/internal/k8s"
	"os"
	"os/exec"
	"path/filepath"
)

func randomFreePort() int {
	max := 45000
	min := 40000

	return rand.IntN(max-min) + min
}

func Sync(name string, localDir string) error {
	localPort := randomFreePort()
	remotePort := 873

	clientSet, err := k8s.NewClientSet()
	if err != nil {
		return err
	}

	// k8s.ExecInPod(clientSet, name)
	config, err := k8s.GetConfig()
	if err != nil {
		return err
	}

	readyChan, err := k8s.PortForward(config, clientSet, name, localPort, remotePort)
	if err != nil {
		return err
	}
	<-readyChan

	log.Printf("Channel succesfully created on port: %d\n", localPort)

	cmd := exec.Command(
		"rsync",
		"-az",
		"--delete",
		filepath.Clean(localDir)+"/",
		fmt.Sprintf("rsync://localhost:%d/workspace", localPort),
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	log.Printf("rsync executed\n")

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
