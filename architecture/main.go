package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

// distro contains the name of the image that docker must load after QEMU.
// When qemu architecture is empty, the GOARCH value can be used.
// When docker image name is empty, QEMU and docker must not be used.
var distro = []struct {
	o   string // GOOS
	a   string // GOARCH
	p   string // Processor variable, GOARM for now
	q   string // QEMU architecture name. Defaults to GOARCH.
	tag string // Docker default tag
	d   string // Full docker image name overriding all values
	doc string // Documentation
}{
	{"windows", "amd64", "", "", "", "", "native on Github CI"},
	{"linux", "amd64", "", "", "", "", "native on Github CI"},
	{"macos", "amd64", "", "", "", "", "native on Github CI"},
	{"linux", "s390x", "", "", "", "", ""},
	{"linux", "ppc64le", "", "", "", "", ""},
	{"linux", "riscv64", "", "", "edge", "", ""},
	{"linux", "arm", "5", "", "", "", "arm v5"},
	{"linux", "arm", "6", "", "", "", "arm v6"},
	{"linux", "arm", "7", "", "", "", "arm v7"},
	{"linux", "arm64", "8", "", "", "", "arm v8"},
	{"linux", "386", "", "", "", "i386/golang", ""},
}

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

// Build docker arm name like arm[32|64]v[5|6|7|8]
func buildArm(s string) (v string) {
	switch s {
	case "8":
		return "arm64v8"
	default:
		return fmt.Sprintf("arm32v%s", s)
	}
}

// If distro does not contain any line, a Docker image named "GOARCH/golang" will be loaded.
// The name of the image can be overridden using flag "imagebase".
func main() {
	// TODO "xcompile" will become "xalpine" on master branch
	imagebase := flag.String("imagebase", "golang", "name of image base")
	flag.Parse()
	containername := "x" + *imagebase
	goos := setDefault("GOOS", runtime.GOOS)
	goarch := setDefault("GOARCH", runtime.GOARCH)
	if goos == runtime.GOOS && goarch == runtime.GOARCH {
		log.Println("No cross-compiling requested")
		return
	}
	// processor is loaded from the relevant processor variable.
	// When empty the default docker image is loaded
	processor := ""
	switch goarch {
	case "arm64":
		processor = "8"
	case "arm":
		processor = os.Getenv("GOARM")
	}
	qemuarch, imagename, imagetag := "", "", ""
	processorfound := false
	for _, d := range distro {
		if d.o == goos && d.a == goarch {
			qemuarch = d.q
			imagename = d.d
			imagetag = d.tag
			if d.p == processor {
				processorfound = true
				break
			}
		}
	}
	if !processorfound {
		log.Fatalf("No known docker image for %s/%s with processor %s\n", goos, goarch, processor)
		return
	}
	if imagename == "" {
		s := *imagebase
		dockerarch := goarch
		switch dockerarch {
		case "arm", "arm64":
			dockerarch = buildArm(processor)
		}
		imagename = fmt.Sprintf("%s/%s", dockerarch, s)
		if imagetag != "" {
			imagename = fmt.Sprintf("%s:%s", imagename, imagetag)
		}
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
		"-e", setParam("GOTOOLCHAIN", "auto"),
		"--name", containername,
		imagename)
	log.Printf("%v", image)
	out, err = image.Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", out)
}
