package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"

	"github.com/sirupsen/logrus"
)

func main() {
	err := start()
	if err != nil {
		logrus.Fatalf("%v", err)
	}
}

func start() (err error) {
	initArgs := []string{"init"}
	if remoteFile := os.Getenv("INPUT_BACKEND_CONFIG"); remoteFile != "" {
		initArgs = append(initArgs, "-backend-config="+remoteFile)
	}
	initArgs = append(initArgs, os.Getenv("INPUT_ROOT_DIR"))
	err = run("terraform", initArgs...)
	if err != nil {
		return err
	}

	for _, action := range strings.Split(os.Getenv("INPUT_ACTION"), ",") {
		err := runAction(action, os.Getenv("INPUT_VAR_FILE"), os.Getenv("INPUT_ROOT_DIR"))
		if err != nil {
			return err
		}
	}

	return nil
}

func runAction(action, varFile, rootDir string) (err error) {
	actionArgs := []string{action}
	if varFile != "" {
		actionArgs = append(actionArgs, "-var-file="+varFile)
	}
	actionArgs = append(actionArgs, rootDir)

	return run("terraform", actionArgs...)
}

func run(command string, args ...string) (err error) {
	fmt.Println(command, args)
	sigs := make(chan os.Signal, 1)
	cmd := exec.Command(command, args...)
	cmd.Env = os.Environ()
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

	if err := cmd.Start(); err != nil {
		logrus.Fatalf("%v", err)
	}
	// wait for the command to finish
	waitCh := make(chan error, 1)
	go func() {
		waitCh <- cmd.Wait()
		close(waitCh)
	}()

	for {
		select {
		case sig := <-sigs:
			if err = cmd.Process.Signal(sig); err != nil {
				logrus.Errorf("%v", err)
				break
			}
		case err := <-waitCh:
			var waitStatus syscall.WaitStatus
			if exitError, ok := err.(*exec.ExitError); ok {
				waitStatus = exitError.Sys().(syscall.WaitStatus)
				os.Exit(waitStatus.ExitStatus())
			}
			if err != nil {
				logrus.Fatalf("%v", err)
			}
			return nil
		}
	}
}
