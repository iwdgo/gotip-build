package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

// Empty values means that GOARCH can be used
var distro = []struct {
	o   string // GOOS
	a   string // GOARCH
	q   string // QEMU architecture name
	d   string // Docker image name
	doc string // Documentation
}{
	{"windows", "amd64", "", "", "native on Github CI"},
	{"linux", "amd64", "", "", "native on Github CI"},
	{"macos", "amd64", "", "", "native on Github CI"},
	{"linux", "s390x", "", "s390x/alpine", ""},
	{"linux", "ppc64le", "", "ppc64le/alpine", ""},
	{"linux", "riscv64", "", "riscv64/alpine:edge", ""},
	{"linux", "arm64", "", "arm64v8/alpine", ""},
	{"linux", "arm", "arm/v7", "arm32v7/alpine", ""}, // TODO Add v6
	{"linux", "386", "", "i386/alpine", ""},
}

const containername = "xcompile"

func main() {
	goos := os.Getenv("GOOS")
	if goos == "" {
		goos = runtime.GOOS
	}
	goarch := os.Getenv("GOARCH")
	if goarch == "" {
		goarch = runtime.GOARCH
	}
	if goos == runtime.GOOS && goarch == runtime.GOARCH {
		fmt.Println("No cross-compiling requested")
		return
	}
	qemuarch, imagename := "", ""
	for _, d := range distro {
		if d.o == goos && d.a == goarch {
			qemuarch = d.q
			imagename = d.d
		}
	}
	if imagename == "" {
		fmt.Printf("Unsupported %s/%s\n", goos, goarch)
		return
	}
	if qemuarch == "" {
		qemuarch = goarch
	}

	qemu := exec.Command("sh", "docker", "run", "--rm", "--privileged",
		"tonistiigi/binfmt:latest", "--install", qemuarch)
	out, err := qemu.Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", out)
	log.Printf("Starting %s as 'xcompile' for %s", imagename, goarch)

	timeout := os.Getenv("GO_TEST_TIMEOUT_SCALE")
	if timeout == "" {
		timeout = "4"
	}
	timeout = fmt.Sprintf("GO_TEST_TIMEOUT_SCALE=%s", timeout)

	image := exec.Command("sh", "docker", "run", "-d", "-t",
		"--platform", goarch,
		// TODO Use go_variables when set
		"-e", timeout,
		"-e", "GOPROXY=https://proxy.golang.org,direct",
		"-e", "GOSUMDB=sum.golang.org",
		"--name", "xcompile",
		imagename)
	log.Printf("%v", image)
	out, err = image.Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", out)
	ps := exec.Command("sh", "docker", "ps")
	out, err = ps.Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", out)

}
