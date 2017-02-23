# -----------------------------------------------------------------------------
# Go project environment

# for gitCommit version
GIT_REVISION = $(shell git rev-parse --short HEAD)
GO_PACKAGES = $(shell go list ./... | grep -v -e 'vendor' -e 'builtinheader' -e 'symbol/internal')

GO_BUILD_FLAGS := -v
GO_TEST_FLAGS := -v

GO_GCFLAGS ?= 
# insert gitCommit version
GO_LDFLAGS := -X "main.Revision=$(GIT_REVISION)"

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
	GO_LDFLAGS += -w -s
	CGO_CFLAGS += -O3
endif

# -----------------------------------------------------------------------------
# cgo compile flags

CC = $(shell llvm-config --bindir)/clang
CXX = $(shell llvm-config --bindir)/clang++

UNAME := $(shell uname)
LLVM_LIBDIR = $(shell llvm-config --libdir)

# static or dynamic link flags
ifneq ($(STATIC),)
	# add 'static' Go build tags for go-clang
	GO_BUILD_TAGS += static

	# add clang supported latest c++ std version and libc++ stdlib
	CGO_CXXFLAGS += -std=c++1z -stdlib=libc++

	ifeq ($(UNAME),Linux)
		# darwin ld(64) linker doesn't support "-extldflags='-static'" flag, because that is basicallly for building the xnu kernel.
		# Also will occur 'not found -crt0.o object file' error.
		# If install the 'Csu' from the opensource.apple.com, passes -crt0.o error, but needs libpthread.a static library.
		GO_LDFLAGS += -extldflags=-static
		# -Bstatic: Do not link against shared libraries.
		# for statically link the libclang libraries.
		CGO_LDFLAGS += -Wl,-Bstatic
	endif

	# add LLVM dependencies static libraries
	CGO_LDFLAGS += $(shell llvm-config --libfiles --link-static)

	LIBCLANG_STATIC_LIBS := \
		libclang \
		libclangAnalysis \
		libclangApplyReplacements \
		libclangARCMigrate \
		libclangAST \
		libclangASTMatchers \
		libclangBasic \
		libclangChangeNamespace \
		libclangCodeGen \
		libclangDriver \
		libclangDynamicASTMatchers \
		libclangEdit \
		libclangFormat \
		libclangFrontend \
		libclangFrontendTool \
		libclangIncludeFixer \
		libclangIncludeFixerPlugin \
		libclangIndex \
		libclangLex \
		libclangMove \
		libclangParse \
		libclangQuery \
		libclangRename \
		libclangReorderFields \
		libclangRewrite \
		libclangRewriteFrontend \
		libclangSema \
		libclangSerialization \
		libclangStaticAnalyzerCheckers \
		libclangStaticAnalyzerCore \
		libclangStaticAnalyzerFrontend \
		libclangTidy \
		libclangTidyBoostModule \
		libclangTidyCERTModule \
		libclangTidyCppCoreGuidelinesModule \
		libclangTidyGoogleModule \
		libclangTidyLLVMModule \
		libclangTidyMiscModule \
		libclangTidyModernizeModule \
		libclangTidyMPIModule \
		libclangTidyPerformanceModule \
		libclangTidyPlugin \
		libclangTidyReadabilityModule \
		libclangTidySafetyModule \
		libclangTidyUtils \
		libclangTooling \
		libclangToolingCore \
		libfindAllSymbols

	# add libclang static libraries
	CGO_LDFLAGS += $(foreach lib,$(LIBCLANG_STATIC_LIBS),$(LLVM_LIBDIR)/$(lib).a)

	ifeq ($(UNAME),Linux)
		# -Bdynamic: Link against dynamic libraries.
		# End of libclang static libraries list.
		CGO_LDFLAGS += -Wl,-Bdynamic
	endif

	# add LLVM dependency system library. such as libm, ncurses and zlib.
	CGO_LDFLAGS += $(shell llvm-config --system-libs --link-static)

	# add '-c++' for only Darwin.
	# avoid 'Undefined symbols for architecture x86_64 "std::__1::__shared_count::__add_shared()"' or etc.
	CGO_LDFLAGS += $(if $(findstring Darwin,$(UNAME)),-lc++,)
else
	# dynamic link build
	CGO_LDFLAGS += -L$(LLVM_LIBDIR)

	# link against LLVM's libc++ for only Darwin.
	# avoid link the macOS system libc++ library.
	ifeq ($(UNAME),Darwin)
		CGO_CFLAGS += -Wl,-rpath,$(LLVM_LIBDIR)
		CGO_CXXFLAGS += -Wl,-rpath,$(LLVM_LIBDIR)
	endif
endif

CGO_FLAGS = CC="$(CC)" CXX="$(CXX)" CGO_CFLAGS="$(CGO_CFLAGS)" CGO_CXXFLAGS="$(CGO_CXXFLAGS)" CGO_LDFLAGS="$(CGO_LDFLAGS)"

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

build:
	$(CGO_FLAGS) go build $(GO_BUILD_FLAGS) -tags '$(GO_BUILD_TAGS)' -gcflags '$(GO_GCFLAGS)' -ldflags '$(GO_LDFLAGS)' ./cmd/clang-server

build-race: GO_BUILD_FLAGS+=-race
build-race: build

install:
	$(CGO_FLAGS) go install $(GO_BUILD_FLAGS) -tags '$(GO_BUILD_TAGS)' -gcflags '$(GO_GCFLAGS)' -ldflags '$(GO_LDFLAGS)' ./cmd/clang-server

run: build clean/cache
	./clang-server -path /Users/zchee/src/github.com/ccache/ccache

run-race: GO_BUILD_FLAGS+=-race
run-race: run

test:
	go test $(GO_TEST_FLAGS) $(GO_PACKAGES)

lint:
	golint -min_confidence 0.1 $(GO_PACKAGES)

vet:
	go vet -v -race $(GO_PACKAGES)

vendor/install:
	$(CGO_FLAGS) go install -v -x -tags '$(GO_BUILD_TAGS)' $(shell go list ./vendor/...)
	$(CGO_FLAGS) go install -v -x -race -tags '$(GO_BUILD_TAGS)'  $(shell go list ./vendor/...)

vendor/update:
	dep ensure -update -v

vendor/clean:
	@rm -rf $(UNUSED)
	@find vendor -type f -name '*_test.go' -print -exec rm -fr {} ";" || true
	@find vendor \( -name 'testdata' -o -name 'cmd' -o -name 'examples' -o -name 'testutil' -o -name 'manualtest' \) -print | xargs rm -rf || true
	@find vendor \( -name 'Makefile' -o -name 'Dockerfile' -o -name 'CHANGELOG*' -o -name '.travis.yml' -o -name 'circle.yml' -o -name '.appveyor.yml' -o -name 'appveyor.yml' -o -name '*.json' -o -name '*.proto' -o -name '*.sh' -o -name '*.pl' -o -name 'codereview.cfg' -o -name '.github' -o -name '.gitignore' -o -name '.gitattributes' \) -print | xargs rm -rf || true

fbs:
	@${RM} -r ./symbol/internal/symbol
	flatc --go --grpc $(shell find ./symbol -type f -name '*.fbs')
	@gofmt -w ./symbol/internal/symbol

prof/cpu:
	go tool pprof -top -cum clang-server cpu.pprof

prof/mem:
	go tool pprof -top -cum clang-server mem.pprof

prof/block:
	go tool pprof -top -cum clang-server block.pprof

prof/trace:
	go tool pprof -top -cum clang-server trace.pprof

clang-format:
	clang-format -i -sort-includes $(shell find testdata -type f -name '*.c' -or -name '*.cpp')

clean:
	${RM} clang-server *.pprof

clean/cachedir:
	${RM} -r $(XDG_CACHE_HOME)/clang-server


.PHONY: build install run test lint vet glide vendor/restore vendor/install vendor/update vendor/clean fbs clang-format clean clean/cachedir
