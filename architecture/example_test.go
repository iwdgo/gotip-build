package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"testing"
)

func TestNative(t *testing.T) {
	f, err := os.CreateTemp(t.TempDir(), "*.log")
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		if e := f.Close(); e != nil {
			t.Fatal(e)
		}
	}(f)
	log.SetOutput(f)

	main()

	sw := "No cross-compiling requested"
	if s, _ := os.ReadFile(f.Name()); !bytes.Contains(s, []byte(sw)) {
		t.Fatalf("expected %s is not in output %s", sw, s)
	}
}

// TODO log.Fatal is used to stop any processing
func TestUnsupported(t *testing.T) {
	t.Skip("Cannot recover from log.Fatal")
	err := os.Setenv("GOOS", "windows")
	if err != nil {
		log.Fatalf("%v", err)
	}
	err = os.Setenv("GOARCH", "arm64")
	if err != nil {
		log.Fatalf("%v", err)
	}
	main()
}

// TODO Requires docker to be available which seems useless
func TestSupported(t *testing.T) {
	t.Skip("Requires docker")
	err := os.Setenv("GOOS", "linux")
	if err != nil {
		log.Fatalf("%v", err)
	}
	err = os.Setenv("GOARCH", "arm64")
	if err != nil {
		log.Fatalf("%v", err)
	}
	main()
	ps := exec.Command("sh", "docker", "rm", "xalpine")
	out, err := ps.Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", out)
}
