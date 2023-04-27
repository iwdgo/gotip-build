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
	// {"linux", "arm", "arm/v6", "arm32v6/alpine", ""},
	{"linux", "arm", "arm/v7", "arm32v7/alpine", ""},
	{"linux", "386", "", "i386/alpine", ""},
}

const containername = "xcompile"

func setDefault(s string, d string) (v string) {
	c := os.Getenv(s)
	if c == "" {
		return d
	}
	return c
}

func setParam(s string, d string) (v string) {
	return fmt.Sprintf("%s=%s", s, setDefault(s, d))
}

func main() {
	goos := setDefault("GOOS", runtime.GOOS)
	goarch := setDefault("GOARCH", runtime.GOARCH)
	if goos == runtime.GOOS && goarch == runtime.GOARCH {
		fmt.Println("No cross-compiling requested")
		return
	}
	qemuarch, imagename := "", ""
	for _, d := range distro {
		if d.o == goos && d.a == goarch {
			qemuarch = d.q
			imagename = d.d
			break
		}
	}
	if imagename == "" {
		fmt.Printf("Unsupported %s/%s\n", goos, goarch)
		return
	}
	if qemuarch == "" {
		qemuarch = goarch
	}

	qemu := exec.Command("docker", "run", "--rm", "--privileged", "tonistiigi/binfmt:latest",
		"--install", qemuarch)
	out, err := qemu.Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", out)
	log.Printf("Starting %s as %s for %s", imagename, containername, goarch)

	image := exec.Command("docker", "run", "-d", "-t",
		"--platform", goarch,
		"-e", setParam("GO_TEST_TIMEOUT_SCALE", "4"),
		"-e", setParam("GOPROXY", "https://proxy.golang.org,direct"),
		"-e", setParam("GOSUMDB", "sum.golang.org"),
		"--name", containername,
		imagename)
	log.Printf("%v", image)
	out, err = image.Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", out)
}
