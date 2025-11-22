package cmd

import (
	"fmt"
	"log"
	"math/rand/v2"
	"miniDevPod/internal/k8s"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
)

func randomFreePort() int {
	max := 45000
	min := 40000

	return rand.IntN(max-min) + min
}

func runRsync(localDir string, localPort int) error {
	cmd := exec.Command(
		"rsync",
		"-az",
		"--delete",
		filepath.Clean(localDir)+"/",
		fmt.Sprintf("rsync://localhost:%d/workspace", localPort),
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func watchLoop(w *fsnotify.Watcher, localDir string, localPort int) {
	for {
		timer := time.NewTimer(400 * time.Millisecond)
		select {
		// Read from Errors.
		case err, ok := <-w.Errors:
			if !ok { // Channel was closed (i.e. Watcher.Close() was called).
				return
			}
			log.Printf("ERROR: %s", err)
		// Read from Events.
		case e, ok := <-w.Events:
			if !ok { // Channel was closed (i.e. Watcher.Close() was called).
				log.Println(e)
				return
			}

			log.Printf("Detected changes")
			// event received, reset timer
			timer.Reset(400 * time.Millisecond)
		case <-timer.C:
			// when timer ends stop debouncing and call rsync
			runRsync(localDir, localPort)
			log.Printf("Sync Complete")
		}
	}
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

	// initial rsync
	runRsync(localDir, localPort)

	log.Printf("Initial sync completed\n")

	// start filesystem watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()
	err = watcher.Add(localDir)
	if err != nil {
		return err
	}

	// start go routine that keeps updating files on pod
	go watchLoop(watcher, localDir, localPort)

	// listen for ctrl + c before exiting
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan

	return nil
}
