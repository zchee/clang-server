LLVM_LIBDIR = $(shell llvm-config --libdir)
GIT_REVISION = $(shell git rev-parse --short HEAD)

GO_GCFLAGS ?= 
GO_LDFLAGS := -X "main.Revision=$(GIT_REVISION)"

CGO_CFLAGS ?=
CGO_LDFLAGS ?= -L$(LLVM_LIBDIR)
CGO_FLAGS=CGO_CFLAGS='$(CGO_CFLAGS)' CGO_CXXFLAGS='$(CGO_CXXFLAGS)' CGO_LDFLAGS='$(CGO_LDFLAGS)'

GO_BUILD_FLAGS ?=
GO_TEST_FLAGS := 
GO_PACKAGES = $(shell go list ./... | grep -v -e 'vendor' -e 'builtinheader' -e 'symbol/internal')


ifneq ($(CLANG_SERVER_DEBUG),)
	GO_GCFLAGS+= -N -l
	CGO_CFLAGS+= -g
	GO_BUILD_FLAGS += -v -x
	GO_TEST_FLAGS += -v -race
else
	GO_LDFLAGS += -w -s
	CGO_CFLAGS += -O3
endif

ifneq ($(STATIC),)
	GO_BUILD_TAGS += static

	LLVM_LIBS := \
		libclang \
		libclangARCMigrate \
		libclangAST \
		libclangASTMatchers \
		libclangAnalysis \
		libclangApplyReplacements \
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
		libclangTidyMPIModule \
		libclangTidyMiscModule \
		libclangTidyModernizeModule \
		libclangTidyPerformanceModule \
		libclangTidyPlugin \
		libclangTidyReadabilityModule \
		libclangTidyUtils \
		libclangTooling \
		libclangToolingCore \
		libfindAllSymbols

	LLVM_DEPS := $(shell llvm-config --system-libs)

	# TODO(zchee): Support windows
	GO_OS := $(shell go env GOOS)
	ifeq ($(GO_OS),linux)
		# Basically, darwin ld linker flag does not support "-extldflags='-static'" flag, because that for only build the xnu kernel. And will not found -crt0.o object file.
		# If install the 'Csu' from opensource.apple.com, passes -crt0.o error but needs libpthread.a static library.
		GO_LDFLAGS += -extldflags=-static
		CGO_LDFLAGS += -Wl,-Bstatic
	endif
	CGO_CXXFLAGS := -std=c++1y -stdlib=libc++
	CGO_LDFLAGS += $(shell llvm-config --libfiles --link-static) $(foreach lib,$(shell command ls /usr/lib/llvm-3.9/lib/libclang*.a),$(lib)) $(LLVM_DEPS) -L/usr/lib -lc++
	ifeq ($(GO_OS),linux)
		CGO_LDFLAGS += -Wl,-Bdynamic
	endif
endif

UNUSED := \
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
	vendor/google.golang.org/grpc/Documentation \
	vendor/google.golang.org/grpc/benchmark \
	vendor/google.golang.org/grpc/test \
	vendor/google.golang.org/grpc/testdata \
	vendor/google.golang.org/grpc/transport/testdata \
	vendor/google.golang.org/grpc/reflection/grpc_testing


default: build

build: clean/cache
	$(CGO_FLAGS) go build -gcflags '$(GO_GCFLAGS)' -ldflags '$(GO_LDFLAGS)' $(GO_BUILD_FLAGS) -tags '$(GO_BUILD_TAGS)' ./cmd/clang-server

build-race: GO_BUILD_FLAGS+=-race
build-race: build

install:
	$(CGO_FLAGS) go install -gcflags '$(GO_GCFLAGS)' -ldflags '$(GO_LDFLAGS)' $(GO_BUILD_FLAGS) -tags '$(GO_BUILD_TAGS)' $(shell go list ./... | grep -v -e 'cmd' -e 'vendor' -e 'builtinheader' -e 'symbol/internal')

run: clean/cache
	$(CGO_FLAGS) go run -gcflags '$(GO_GCFLAGS)' -ldflags '$(GO_LDFLAGS)' $(GO_BUILD_FLAGS) ./cmd/clang-server/main.go -path /Users/zchee/src/github.com/neovim/neovim

run-race: GO_BUILD_FLAGS+=-race
run-race: run

test:
	go test $(GO_TEST_FLAGS) $(GO_PACKAGES)

lint:
	golint -min_confidence 0.1 $(GO_PACKAGES)

vet:
	go vet -v -race $(GO_PACKAGES)

prof/cpu:
	go tool pprof -top -cum clang-server cpu.pprof

prof/mem:
	go tool pprof -top -cum clang-server mem.pprof

prof/block:
	go tool pprof -top -cum clang-server block.pprof

prof/trace:
	go tool pprof -top -cum clang-server trace.pprof

vendor/install:
	$(CGO_FLAGS) go install -v -x -tags '$(GO_BUILD_TAGS)' $(shell go list ./vendor/...)
	$(CGO_FLAGS) go install -v -x -race -tags '$(GO_BUILD_TAGS)'  $(shell go list ./vendor/...)

vendor/update:
	dep ensure -update -v

vendor/clean:
	@rm -rf $(UNUSED)
	@find vendor -type f -name '*_test.go' -print -exec rm -fr {} ";" || true
	@find vendor \( -name 'testdata' -o -name 'cmd' -o -name 'examples' -o -name 'testutil' -o -name 'manualtest' \) -print | xargs rm -rf || true
	@find vendor \( -name 'Makefile' -o -name 'Dockerfile' -o -name 'CHANGELOG*' -o -name '.travis.yml' -o -name 'appveyor.yml' -o -name '*.json' -o -name '*.proto' -o -name '*.sh' -o -name '*.pl' -o -name 'codereview.cfg' -o -name '.github' -o -name '.gitignore' -o -name '.gitattributes' \) -print | xargs rm -rf || true

fbs:
	@${RM} -r ./symbol/internal/symbol
	flatc --go --grpc $(shell find ./symbol -type f -name '*.fbs')
	@gofmt -w ./symbol/internal/symbol

clang-format:
	clang-format -i -sort-includes $(shell find testdata -type f -name '*.c' -or -name '*.cpp')

serve: clean build
	./clang-server

clean:
	${RM} clang-server *.pprof

clean/cache:
	${RM} -r $(XDG_CACHE_HOME)/clang-server/

.PHONY: build install run test lint vet glide vendor/restore vendor/install vendor/update vendor/clean fbs clang-format serve clean clean/db
