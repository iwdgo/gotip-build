package main

import (
	"log"
	"os"
	"os/exec"
)

func Example_native() {
	main()
	// Output: No cross-compiling requested
}

func Example_unsupported() {
	err := os.Setenv("GOOS", "windows")
	if err != nil {
		log.Fatalf("%v", err)
	}
	err = os.Setenv("GOARCH", "arm64")
	if err != nil {
		log.Fatalf("%v", err)
	}
	main()
	// Output: Unsupported windows/arm64
}

func Example_supported() {
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
