# Build Go from source on Windows

A relevant go version must be available.
Currently, tip builds with [Go 1.17.13](https://github.com/golang/go/issues/44505)

First patch files found in directory is applied on tip.
Expected file format is from command like [`git format-patch master`](https://git-scm.com/docs/git-format-patch)

Usage:

```

    - name: Build Go from source
      uses: iwdgo/gotip-build@master-windows
      id: gotip
      with:
        test_build: false

```
