# Build and test Go from source

Go must be available to build the language from source.
In go1.21 cycle, tip built with [Go 1.17.13](https://github.com/golang/go/issues/44505).
[Forward compatibility](https://github.com/golang/go/issues/57001) justifies to use Go 1.19 with the introduction of `GOTOOLCHAIN`.

Besides the [install](https://go.dev/doc/install/source) documentation, more details on [wiki](https://github.com/iwdgo/gotip-build/wiki).

Patch files found in directory are applied on tip.
Expected file format is from command like [`git format-patch master`](https://git-scm.com/docs/git-format-patch)

`GOROOT_BOOTSTRAP` is set to `go env GOROOT` when not set by `go_variables`.

Usage with `bash`:

```

    - name: Build Go from source
      uses: iwdgo/gotip-build@v0.4.0
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
