# go-clang/v3.9 [![GoDoc](https://godoc.org/github.com/go-clang/v3.9?status.png)](https://godoc.org/github.com/go-clang/v3.9) [![Build Status](https://travis-ci.org/go-clang/v3.9.svg?branch=master)](https://travis-ci.org/go-clang/v3.9)

Native Go bindings for Clang's C API.

## Install/Update

```bash
CGO_LDFLAGS="-L`llvm-config --libdir`" \
  go get -u github.com/go-clang/v3.9/...
```

## Usage

An example on how to use the AST visitor of the Clang API can be found in [/cmd/go-clang-dump/main.go](/cmd/go-clang-dump/main.go)

## I need bindings for a different Clang version

The Go bindings are placed in their own repositories to provide the correct bindings for the corresponding Clang version. A list of supported versions can be found in [go-clang/gen's README](https://github.com/go-clang/gen#where-are-the-bindings).

## I found a bug/missing a feature in go-clang

We are using the issue tracker of the `go-clang/gen` repository. Please go through the [open issues](https://github.com/go-clang/gen/issues) in the tracker first. If you cannot find your request just open up a [new issue](https://github.com/go-clang/gen/issues/new).

## How is this binding generated?

The [go-clang/gen](https://github.com/go-clang/gen) repository is used to automatically generate this binding.

# License

This project, like all go-clang projects, is licensed under a BSD-3 license which can be found in the [LICENSE file](https://github.com/go-clang/license/blob/master/LICENSE) in [go-clang's license repository](https://github.com/go-clang/license)
