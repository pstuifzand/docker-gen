package main

import (
	"flag"
	"os"
	"testing"
)

func TestDefaultEndpoint(t *testing.T) {
	endpoint, err := getEndpoint()
	if err != nil {
		t.Fatal("%s", err)
	}
	if endpoint != "unix:///var/run/docker.sock" {
		t.Fatal("Expected unix:///var/run/docker.sock")
	}
}

func TestDockerHostEndpoint(t *testing.T) {
	err := os.Setenv("DOCKER_HOST", "tcp://127.0.0.1:4243")
	if err != nil {
		t.Fatalf("Unable to set DOCKER_HOST: %s", err)
	}

	endpoint, err := getEndpoint()
	if err != nil {
		t.Fatal("%s", err)
	}

	if endpoint != "tcp://127.0.0.1:4243" {
		t.Fatal("Expected tcp://127.0.0.1:4243")
	}
}

func TestDockerFlagEndpoint(t *testing.T) {

	initFlags()
	err := os.Setenv("DOCKER_HOST", "tcp://127.0.0.1:4243")
	if err != nil {
		t.Fatalf("Unable to set DOCKER_HOST: %s", err)
	}

	// flag value should override DOCKER_HOST and default value
	err = flag.Set("endpoint", "tcp://127.0.0.1:5555")
	if err != nil {
		t.Fatalf("Unable to set endpoint flag: %s", err)
	}

	endpoint, err := getEndpoint()
	if err != nil {
		t.Fatal("%s", err)
	}
	if endpoint != "tcp://127.0.0.1:5555" {
		t.Fatal("Expected tcp://127.0.0.1:5555")
	}
}

func TestUnixNotExists(t *testing.T) {

	endpoint = ""
	err := os.Setenv("DOCKER_HOST", "unix:///does/not/exist")
	if err != nil {
		t.Fatalf("Unable to set DOCKER_HOST: %s", err)
	}

	_, err = getEndpoint()
	if err == nil {
		t.Fatal("endpoint should have failed")
	}
}

func TestUnixBadFormat(t *testing.T) {
	endpoint = "unix:/var/run/docker.sock"
	_, err := getEndpoint()
	if err == nil {
		t.Fatal("endpoint should have failed")
	}
}
