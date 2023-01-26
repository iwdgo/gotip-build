# Build Go from source

A relevant go version must be available.
Currently, tip builds with [Go 1.17.13](https://github.com/golang/go/issues/44505)

Patch files found are committed on tip.
File format is expected to come from command like [`git format-patch master`](https://git-scm.com/docs/git-format-patch)

`GOROOT_BOOTSTRAP` is set to `go env GOROOT` when not set by `go_variables`.

Usage with `bash`:

```

    - name: Build Go from source
      uses: iwdgo/gotip-build@v0.2.1
      id: gotip
      with:
        go_variables: GOROOT_FINAL=/ CGO_ENABLED=0
        test_build: true

```

On Windows, using `powershell`:

```

    - name: Build Go from source on Windows
      uses: iwdgo/gotip-build@master-windows
      id: gotip
      with:
        test_build: false

```
