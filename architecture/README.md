### How to

```bash
          go install github.com/iwdgo/gotip-build/architecture@latest
          architecture --imagebase=alpine
```

```log
go: downloading github.com/iwdgo/gotip-build v0.4.0
go: downloading github.com/iwdgo/gotip-build/architecture v0.0.0-20230618071648-36eba343132f
2023/06/18 08:57:46 {
  "supported": [
    "linux/amd64",
    "linux/arm64",
    "linux/386"
  ],
  "emulators": [
    "cli",
    "llvm-12-runtime.binfmt",
    "llvm-13-runtime.binfmt",
    "llvm-14-runtime.binfmt",
    "python3.10",
    "qemu-aarch64"
  ]
}
2023/06/18 08:57:46 Starting arm64v8/alpine as xalpine for arm64
2023/06/18 08:57:46 /usr/bin/docker run -d -t --platform arm64 -e GO_TEST_TIMEOUT_SCALE=4 -e GOPROXY=https://proxy.golang.org,direct -e GOSUMDB=sum.golang.org -e GOTOOLCHAIN=auto --name xalpine arm64v8/alpine
2023/06/18 08:57:47 f8eba263139dc873b764c71b8994600108ac0bceb6cacb3f975ceeabf4d3a183
```
