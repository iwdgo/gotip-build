# Build Go using toolchain on tip

A relevant go version must be available.
Currently, tip builds with [Go 1.17.13](https://github.com/golang/go/issues/44505)

Usage:

```

    - name: Build Go from source
      uses: actions/gotip-build@v0
      id: gotip

```
