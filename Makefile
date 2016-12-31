LLVM_LIBDIR := $(shell llvm-config --libdir)
GIT_REVISION := $(shell git rev-parse --short HEAD)

GO_GCFLAGS ?= 
GO_LDFLAGS := -X "main.Revision=$(GIT_REVISION)"
CGO_CFLAGS ?=
CGO_LDFLAGS ?= -L$(LLVM_LIBDIR)
GO_BUILD_FLAGS ?=
GO_TEST_FLAGS := 
GO_PACKAGES := $(shell glide novendor)


ifneq ($(CLANG_SERVER_DEBUG),)
	GO_GCFLAGS += -N -l
	CGO_CFLAGS += -g -O0
	GO_BUILD_FLAGS += -v -x
	GO_TEST_FLAGS += -v -race
else
	GO_LDFLAGS += -w -s
	CGO_CFLAGS += -O3
endif

ifneq ($(shell command -v ccache-clang 2> /dev/null),)
	CC := ccache-clang
endif
ifneq ($(shell command -v ccache-clang++ 2> /dev/null),)
	CXX := ccache-clang++
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

	LLVM_DEPS := \
		ncursesw \
		zlib \
		libffi

	# TODO(zchee): Support windows
	GO_OS := $(shell go env GOOS)
	ifeq ($(GO_OS),linux)
		# Basically, darwin ld linker does not support -static flag because that for only build the xnu kernel. And will not found -crt0.o object file.
		# If install the 'Csu' from opensource.apple.com, passes -crt0.o error but needs libpthread.a static library.
		GO_LDFLAGS += -extldflags "-static"
		CGO_LDFLAGS += -Wl,-Bstatic
	endif
	CGO_LDFLAGS += $(shell llvm-config --libfiles) $(foreach lib,$(LLVM_LIBS),$(LLVM_LIBDIR)/$(lib).a) $(shell pkg-config $(LLVM_DEPS) --libs --static) -lc++
	ifeq ($(GO_OS),linux)
		CGO_LDFLAGS += -Wl,-Bdynamic
	endif
endif


default: build

build:
	go build -gcflags '$(GO_GCFLAGS)' -ldflags '$(GO_LDFLAGS)' $(GO_BUILD_FLAGS) ./cmd/clang-server

build-race: GO_BUILD_FLAGS+=-race;build

install:
	go install -gcflags '$(GO_GCFLAGS)' -ldflags '$(GO_LDFLAGS)' $(GO_BUILD_FLAGS) ./cmd/clang-server

run: clean-db
	go run -gcflags '$(GO_GCFLAGS)' -ldflags '$(GO_LDFLAGS)' $(GO_BUILD_FLAGS) ./cmd/clang-server/main.go -path /Users/zchee/src/github.com/neovim/neovim

test:
	go test $(GO_TEST_FLAGS) $(GO_PACKAGES)

lint:
	@for pkg in $(shell go list ./... | grep -v -e vendor -e symbol/internal | sed 's/github.com\/zchee\/clang-server/\./g'); do golint $$pkg; done

vet:
	@go vet -v -race $(GO_PACKAGES)

glide:
ifeq ($(shell command -v glide 2> /dev/null),)
	go get -v github.com/Masterminds/glide
endif

vendor/restore: glide
	glide install

vendor/install: glide
	CGO_CFLAGS='$(CGO_CFLAGS)' CGO_LDFLAGS='$(CGO_LDFLAGS)' go install -v -x -tags '$(GO_BUILD_TAGS)' $(shell glide list 2> /dev/null  | awk 'NR > 1{print $$1}' | sed s'/^/.\/vendor\//g')
	CGO_CFLAGS='$(CGO_CFLAGS)' CGO_LDFLAGS='$(CGO_LDFLAGS)' go install -v -x -race -tags '$(GO_BUILD_TAGS)' $(shell glide list 2> /dev/null  | awk 'NR > 1{print $$1}' | sed s'/^/.\/vendor\//g')

vendor/update: glide
	glide cache-clear
	glide update

vendor/clean: glide
	@glide-vc --only-code --no-tests
	@cp -r $(GOPATH)/src/github.com/go-clang/v3.9/clang/clang-c ./vendor/github.com/go-clang/v3.9/clang

fbs:
	${RM} -r ./symbol/internal/symbol
	flatc --go --grpc $(shell find ./symbol -type f -name '*.fbs')

clang-format:
	clang-format -i -sort-includes $(shell find testdata -type f -name '*.c' -or -name '*.cpp')

serve: clean build
	./clang-server

clean:
	${RM} clang-server

clean/db:
	${RM} -r $(XDG_CACHE_HOME)/clang-server/

.PHONY: build install run test lint vet glide vendor/restore vendor/install vendor/update vendor/clean fbs clang-format serve clean clean/db
