.PHONY: all install install-dependencies install-tools test test-full test-verbose

export ROOT_DIR := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

export CC := clang
export CXX := clang++

ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
$(eval $(ARGS):;@:) # turn arguments into do-nothing targets
export ARGS

all: install test

install:
	CGO_LDFLAGS="-L`llvm-config --libdir`" go install ./...
install-dependencies:
	go get -u github.com/stretchr/testify/...
install-tools:

test:
	CGO_LDFLAGS="-L`llvm-config --libdir`" go test -timeout 60s ./...
test-full:
	$(ROOT_DIR)/scripts/test-full.sh
test-verbose:
	CGO_LDFLAGS="-L`llvm-config --libdir`" go test -timeout 60s -v ./...
