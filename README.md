# Build Go from source on Windows

A relevant go version must be available.
Currently, tip builds with [Go 1.17.13](https://github.com/golang/go/issues/44505)

Patch files found in directory are applied on tip.
Expected file format is from command like [`git format-patch master`](https://git-scm.com/docs/git-format-patch)

`GOROOT_BOOTSTRAP` is set to `go env GOROOT` when not set by `go_variables`.

Usage:

```

    - name: Build Go from source
      uses: iwdgo/gotip-build@master-windows
      id: gotip
      with:
        go_variables: $GOROOT_FINAL = "/"; $CGO_ENABLED = 0
        test_build: false

```
