#!/bin/bash

set -exuo pipefail

LLVM_VERSION=$(clang --version | grep --max-count=1 "clang version" | sed -r 's/^.*clang version ([0-9]+\.[0-9]+).+$/\1/')

# Test with the race detector
CGO_LDFLAGS="-L`llvm-config --libdir`" go test -timeout 60s -v -race ./...

# Test with the address sanitizer
# TODO there is maybe a problem within clang https://github.com/go-clang/gen/issues/123
# if [ $(echo "$LLVM_VERSION>=3.9" | bc -l) -ne 0 ] && [ $(find `llvm-config --libdir` | grep libclang_rt.san-x86_64.a | wc -l) -ne 0 ]; then CGO_LDFLAGS="-L`llvm-config --libdir` -fsanitize=memory" CGO_CPPFLAGS='-fsanitize=memory -fsanitize-memory-track-origins -fno-omit-frame-pointer' go test -timeout 60s -v -msan ./...; fi
