package main

import (
	"fmt"
	"log"
	"os"
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
	fmt.Println(os.Getenv("qemuarch"))
	fmt.Println(os.Getenv("imagename"))
	// Output: linux/arm64: QEMU=arm64, Docker=arm64v8/alpine
	// arm64
	// arm64v8/alpine
}
