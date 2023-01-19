# Build Go using toolchain on tip

A relevant go version must be available.
Currently, tip builds with [Go 1.17.13](https://github.com/golang/go/issues/44505)

Usage with bash:

```

    - name: Build Go from source
      uses: iwdgo/gotip-build@v0.1.0
      id: gotip
      with:
        test_build: true

```

On Windows, using powershell:

```

    - name: Build Go from source on Windows
      uses: iwdgo/gotip-build@master-windows
      id: gotip

```
