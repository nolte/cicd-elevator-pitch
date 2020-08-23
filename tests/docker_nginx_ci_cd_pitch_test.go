package test

import (
	"crypto/tls"
	"fmt"
	"log"
	"testing"
	"time"

	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"

	"github.com/gruntwork-io/terratest/modules/docker"
	"github.com/phayes/freeport"
)

func TestDockerElevatorPitch(t *testing.T) {
	tag := "nolte/cicdpitch"
	buildOptions := &docker.BuildOptions{
		Tags: []string{tag},
	}

	docker.Build(t, "../pitch", buildOptions)
	port, err := freeport.GetFreePort()
	if err != nil {
		log.Fatal(err)
	}
	opts := &docker.RunOptions{
		OtherOptions: []string{"-d", "-p", fmt.Sprintf("%v:%v", port, 80)},
	}
	optsStop := &docker.StopOptions{}

	containerID := docker.RunAndGetID(t, tag, opts)
	defer docker.Stop(t, []string{containerID}, optsStop)
	tlsConfig := tls.Config{}

	http_helper.HttpGetWithRetryWithCustomValidation(
		t,
		fmt.Sprintf("http://localhost:%v/", port),
		&tlsConfig,
		30,
		10*time.Second,
		func(statusCode int, body string) bool {
			return statusCode == 200
		},
	)
}
