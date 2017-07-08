# -----------------------------------------------------------------------------
# Go project environment

GOPATH ?= $(shell go env GOPATH)
GO_SRCS = $(shell find . -type f -name '*.go' -and -not -iwholename '*vendor*' -and -not -iwholename '*testdata*')
GO_PACKAGES = $(shell go list ./... | grep -v -e 'vendor' -e 'builtinheader' -e 'internal/symbol')
GO_VENDOR_PACKAGES = $(shell go list ./vendor/...)

GO_BUILD_FLAGS := -v
GO_BUILD_TAGS ?=
GO_TEST_FLAGS := -v

GO_GCFLAGS ?= 
GO_LDFLAGS = -X "main.Revision=$(shell git rev-parse --short HEAD)"

CGO_CFLAGS = -Wno-deprecated-declarations
CGO_LDFLAGS ?=
CGO_CXXFLAGS ?=

# for developer build
ifneq ($(CLANG_SERVER_DEBUG),)
	GO_BUILD_FLAGS += -x
endif

# for debugg with debugger
ifneq ($(CLANG_SERVER_DEBUG_DWARF),)
	GO_GCFLAGS += -N -l
	ifeq ($(UNAME),Darwin)
		# need macOS cgo build debugging
		GO_LDFLAGS += -linkmode=internal
	endif
	CGO_CFLAGS += -g
else
	# GO_LDFLAGS += -w -s
	# CGO_CFLAGS += -O3
endif

# -----------------------------------------------------------------------------
# cgo compile flags

CC = $(shell llvm-config --bindir)/clang
CXX = $(shell llvm-config --bindir)/clang++

UNAME := $(shell uname)
LLVM_LIBDIR = $(shell llvm-config --libdir)

# static or dynamic link flags
ifneq ($(STATIC),)
	include ./mk/static.mk
else
	# dynamic link build
	CGO_LDFLAGS += -L$(LLVM_LIBDIR)

	# link against LLVM libraries.
	CGO_CFLAGS += -Wl,-rpath,$(LLVM_LIBDIR)
	CGO_CXXFLAGS += -Wl,-rpath,$(LLVM_LIBDIR)
endif

CGO_FLAGS = CC="$(CC)" CXX="$(CXX)"
ifneq ($(CGO_CFLAGS),)
	CGO_FLAGS += CGO_CFLAGS="$(CGO_CFLAGS)"
endif
ifneq ($(CGO_LDFLAGS),)
	CGO_FLAGS += CGO_LDFLAGS="$(CGO_LDFLAGS)"
endif
ifneq ($(CGO_CXXFLAGS),)
	CGO_FLAGS += CGO_CXXFLAGS="$(CGO_CXXFLAGS)"
endif

ifneq ($(GO_BUILD_TAGS),)
	GO_BUILD_FLAGS += -tags '$(GO_BUILD_TAGS)'
endif
ifneq ($(GO_GCFLAGS),)
	GO_BUILD_FLAGS += -gcflags '$(GO_GCFLAGS)'
endif
ifneq ($(GO_LDFLAGS),)
	GO_BUILD_FLAGS +=  -ldflags '$(GO_LDFLAGS)'
endif

# -----------------------------------------------------------------------------
# vendor packages

UNUSED := \
	vendor/github.com/tmthrgd/asm \
	vendor/github.com/google/flatbuffers/CMake \
	vendor/github.com/google/flatbuffers/CMakeLists.txt \
	vendor/github.com/google/flatbuffers/CONTRIBUTING.md \
	vendor/github.com/google/flatbuffers/ISSUE_TEMPLATE.md \
	vendor/github.com/google/flatbuffers/android \
	vendor/github.com/google/flatbuffers/appveyor.yml \
	vendor/github.com/google/flatbuffers/biicode.conf \
	vendor/github.com/google/flatbuffers/biicode \
	vendor/github.com/google/flatbuffers/composer.json \
	vendor/github.com/google/flatbuffers/docs \
	vendor/github.com/google/flatbuffers/grpc \
	vendor/github.com/google/flatbuffers/include \
	vendor/github.com/google/flatbuffers/java \
	vendor/github.com/google/flatbuffers/js \
	vendor/github.com/google/flatbuffers/net \
	vendor/github.com/google/flatbuffers/php \
	vendor/github.com/google/flatbuffers/pom.xml \
	vendor/github.com/google/flatbuffers/python \
	vendor/github.com/google/flatbuffers/reflection \
	vendor/github.com/google/flatbuffers/samples \
	vendor/github.com/google/flatbuffers/src \
	vendor/github.com/google/flatbuffers/tests \
	vendor/github.com/golang/protobuf/protoc-gen-go \
	vendor/golang.org/x/net/http2/h2i \
	vendor/google.golang.org/grpc/benchmark \
	vendor/google.golang.org/grpc/Documentation \
	vendor/google.golang.org/grpc/interop \
	vendor/google.golang.org/grpc/reflection/grpc_testing \
	vendor/google.golang.org/grpc/stress \
	vendor/google.golang.org/grpc/test \
	vendor/google.golang.org/grpc/testdata \
	vendor/google.golang.org/grpc/transport/testdata

# -----------------------------------------------------------------------------
# target

default: build

build: bin/clang-server bin/clang-client

bin:
	@mkdir ./bin

bin/clang-server: ${GOPATH}/pkg/darwin_amd64/github.com/zchee/clang-server $(GO_SRCS)
	$(CGO_FLAGS) go build $(GO_BUILD_FLAGS) -o ./bin/clang-server ./cmd/clang-server

bin/clang-client: ${GOPATH}/pkg/darwin_amd64/github.com/zchee/clang-server $(GO_SRCS)
	$(CGO_FLAGS) go build $(GO_BUILD_FLAGS) -o ./bin/clang-client ./cmd/clang-client

build-race: GO_BUILD_FLAGS+=-race
build-race: ${GOPATH}/pkg/darwin_amd64_race/github.com/zchee/clang-server
build-race: build/std-race
build-race: build

build/std-race:
	go install -v -x -race std

install:
	$(CGO_FLAGS) go install $(GO_BUILD_FLAGS) -gcflags '$(GO_GCFLAGS)' -ldflags '$(GO_LDFLAGS)' ./cmd/clang-server ./cmd/clang-client

run: build
	# ./bin/clang-server -path /Users/zchee/src/github.com/neovim/neovim
	./bin/clang-server -path /Users/zchee/src/github.com/ccache/ccache

run-race: GO_BUILD_FLAGS+=-race
run-race: run


test:
	@go test $(GO_TEST_FLAGS) $(GO_PACKAGES)

lint:
	@golint -set_exit_status $(GO_PACKAGES)

vet:
	@go vet $(GO_PACKAGES)


vendor/install: ${GOPATH}/pkg/darwin_amd64/github.com/zchee/clang-server ${GOPATH}/pkg/darwin_amd64_race/github.com/zchee/clang-server

${GOPATH}/pkg/darwin_amd64/github.com/zchee/clang-server:
	$(CGO_FLAGS) go install $(GO_BUILD_FLAGS) ./vendor/...

${GOPATH}/pkg/darwin_amd64_race/github.com/zchee/clang-server: GO_BUILD_FLAGS+=-race
${GOPATH}/pkg/darwin_amd64_race/github.com/zchee/clang-server: ${GOPATH}/pkg/darwin_amd64/github.com/zchee/clang-server

vendor/update:
	dep ensure -update -v

vendor/distclean:
	${RM} -r ${GOPATH}/pkg/darwin_amd64/github.com/zchee/clang-server
	${RM} -r ${GOPATH}/pkg/darwin_amd64-race/github.com/zchee/clang-server

fbs:
	@${RM} -r ./internal/symbol
	flatc --go --grpc $(shell find ./symbol -type f -name '*.fbs')
	@gofmt -w ./internal/symbol

clang-format:
	clang-format -i -sort-includes $(shell find testdata -type f -name '*.c' -or -name '*.cpp')


prof/cpu:
	@go tool pprof -top -cum clang-server cpu.pprof

prof/mem:
	@go tool pprof -top -cum clang-server mem.pprof

prof/block:
	@go tool pprof -top -cum clang-server block.pprof

prof/trace:
	@go tool pprof -top -cum clang-server trace.pprof


clean:
	${RM} -r ./bin *.pprof *.out

clean/cachedir:
	${RM} -r $(XDG_CACHE_HOME)/clang-server


.PHONY: build install run test lint vet glide vendor/restore vendor/install vendor/update vendor/clean fbs clang-format clean clean/cachedir
