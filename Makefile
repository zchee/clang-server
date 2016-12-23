GIT_REVISION := $(shell git rev-parse --short HEAD)

LLVM_LIBDIR := $(shell llvm-config --libdir)
LLVM_LIBS := \
	libLLVMAnalysis \
	libLLVMAsmParser \
	libLLVMAsmPrinter \
	libLLVMBitReader \
	libLLVMBitWriter \
	libLLVMCodeGen \
	libLLVMCore \
	libLLVMCoroutines \
	libLLVMCoverage \
	libLLVMDebugInfoCodeView \
	libLLVMDebugInfoDWARF \
	libLLVMDebugInfoMSF \
	libLLVMDebugInfoPDB \
	libLLVMDemangle \
	libLLVMExecutionEngine \
	libLLVMGlobalISel \
	libLLVMIRReader \
	libLLVMInstCombine \
	libLLVMInstrumentation \
	libLLVMInterpreter \
	libLLVMLTO \
	libLLVMLibDriver \
	libLLVMLineEditor \
	libLLVMLinker \
	libLLVMMC \
	libLLVMMCDisassembler \
	libLLVMMCJIT \
	libLLVMMCParser \
	libLLVMMIRParser \
	libLLVMObjCARCOpts \
	libLLVMObject \
	libLLVMObjectYAML \
	libLLVMOption \
	libLLVMOrcJIT \
	libLLVMPasses \
	libLLVMProfileData \
	libLLVMRuntimeDyld \
	libLLVMScalarOpts \
	libLLVMSelectionDAG \
	libLLVMSupport \
	libLLVMSymbolize \
	libLLVMTableGen \
	libLLVMTarget \
	libLLVMTransformUtils \
	libLLVMVectorize \
	libLLVMX86AsmParser \
	libLLVMX86AsmPrinter \
	libLLVMX86CodeGen \
	libLLVMX86Desc \
	libLLVMX86Disassembler \
	libLLVMX86Info \
	libLLVMX86Utils \
	libLLVMipo \
	libclang \
	libclangAST \
	libclangAnalysis \
	libclangBasic \
	libclangDriver \
	libclangEdit \
	libclangFrontend \
	libclangIndex \
	libclangLex \
	libclangParse \
	libclangRewrite \
	libclangSema \
	libclangSerialization

GO_GCFLAGS ?= 
GO_LDFLAGS := -X "main.Revision=$(GIT_REVISION)"
CGO_CFLAGS ?=
CGO_LDFLAGS ?= -L$(LLVM_LIBDIR)

GO_BUILD_FLAGS ?=
GO_TEST_FLAGS := 

PACKAGES := $(shell glide novendor)


ifneq ($(CLANG_SERVER_DEBUG),)
GO_GCFLAGS += -N -l
CGO_CFLAGS += -g -O0
GO_BUILD_FLAGS += -v -x
GO_TEST_FLAGS += -v -race
else
GO_LDFLAGS += -w -s
CGO_CFLAGS += -O3
endif

ifneq ($(STATIC),)
GO_LDFLAGS += -extldflags "-static"
CGO_LDFLAGS ?= -L/usr/lib -lc++
CGO_LDFLAGS += $(foreach lib,$(LLVM_LIBS),$(LLVM_LIBDIR)/$(lib).a)
# CGO_LDFLAGS += $(shell find ~/src/llvm.org/build/lib -type f -name '*.a' | grep -v -e unwind -e lldb -e libc++)
CGO_LDFLAGS += /usr/local/opt/zlib/lib/libz.a /usr/local/opt/ncurses/lib/libncursesw.a
GO_BUILD_FLAGS += unsafe static
endif

ifneq ($(shell command -v ccache-clang 2> /dev/null),)
CC := ccache-clang
endif
ifneq ($(shell command -v ccache-clang++ 2> /dev/null),)
CXX := ccache-clang++
endif


default: build

build:
	go build -gcflags '$(GO_GCFLAGS)' -ldflags '$(GO_LDFLAGS)' -tags netgo -installsuffix netgo $(GO_BUILD_FLAGS) $(PACKAGES)

install:
	go install -gcflags '$(GO_GCFLAGS)' -ldflags '$(GO_LDFLAGS)'

run:
	go run ./main.go

test:
	go test $(GO_TEST_FLAGS) $(PACKAGES)

lint:
	@for pkg in $(shell go list ./... | grep -v -e vendor -e symbol/internal | sed 's/github.com\/zchee\/clang-server/\./g'); do golint $$pkg; done

vet:
	go vet -v -race $(PACKAGES)

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

fbs:
	${RM} -r ./symbol/internal/symbol
	flatc --go --grpc $(shell find ./symbol -type f -name '*.fbs')

clang-format:
	clang-format -i -sort-includes $(shell find testdata -type f -name '*.c' -or -name '*.cpp')

serve: clean build
	./clang-server

clean:
	${RM} clang-server

.PHONY: build install run test lint vet glide vendor/restore vendor/install vendor/update vendor/clean fbs clang-format serve clean
