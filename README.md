# Build and test Go from source

[![linux AMD64](https://github.com/iwdgo/gotip-build/actions/workflows/linux_test.yml/badge.svg)](https://github.com/iwdgo/gotip-build/actions/workflows/linux_test.yml)  
[![macos AMD64](https://github.com/iwdgo/gotip-build/actions/workflows/macos_test.yml/badge.svg)](https://github.com/iwdgo/gotip-build/actions/workflows/macos_test.yml)  
[![windows AMD64](https://github.com/iwdgo/gotip-build/actions/workflows/windows_test.yml/badge.svg)](https://github.com/iwdgo/gotip-build/actions/workflows/windows_test.yml?branch=master-windows)  

A relevant go version must be available.
Currently, tip builds with [Go 1.17.13](https://github.com/golang/go/issues/44505)

Patch files found in directory are applied on tip.
Expected file format is from command like [`git format-patch master`](https://git-scm.com/docs/git-format-patch)

`GOROOT_BOOTSTRAP` is set to `go env GOROOT` when not set by `go_variables`.

Usage with `bash`:

```

    - name: Build Go from source
      uses: iwdgo/gotip-build@v0.3.0
      id: gotip
      with:
        go_variables: GOROOT_FINAL=/ CGO_ENABLED=0
        test_build: true

```

On Windows, using `powershell` is identical except for the version tag which is the branch name `master-windows`.

```

    - name: Build Go from source on Windows
      uses: iwdgo/gotip-build@master-windows
      id: gotip
      with:
        go_variables: $GOROOT_FINAL = "/"; $CGO_ENABLED = 0
        test_build: false

```

Details on [wiki](https://github.com/iwdgo/gotip-build/wiki).
