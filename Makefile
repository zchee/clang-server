GO_GCFLAGS ?= 
GO_LDFLAGS ?=
CGO_LDFLAGS ?= -L$(shell llvm-config --libdir)
# CGO_LDFLAGS += $(foreach lib,libclang.a libclangAST.a libclangAnalysis.a libclangBasic.a libclangDriver.a libclangEdit.a libclangFrontend.a libclangIndex.a libclangLex.a libclangParse.a libclangRewrite.a libclangSema.a libclangSerialization.a,$(shell llvm-config --libdir)/$(lib))
# CGO_LDFLAGS += $(shell find ~/src/llvm.org/build/lib -type f -name '*.a' | grep -v -e unwind -e lldb -e libc++)
# CGO_LDFLAGS += /usr/local/opt/zlib/lib/libz.a /usr/local/opt/ncurses/lib/libncursesw.a

# GO_BUILD_TAGS ?= -tags 'unsafe static'
GO_BUILD_FLAGS := $(GO_BUILD_TAGS)
GO_TEST_FLAGS := 

PACKAGES := $(shell glide novendor)


ifneq ($(CLANG_SERVER_DEBUG),)
GO_GCFLAGS += -N -l
GO_BUILD_FLAGS += -v -x -race
GO_TEST_FLAGS += -v -race
else
GO_LDFLAGS += -w -s
endif


default: build

build:
	go build -gcflags '$(GO_GCFLAGS)' -ldflags '$(GO_LDFLAGS)' $(GO_BUILD_FLAGS) $(PACKAGES)

run:
	go run ./cmd/clang-server/main.go

test:
	go test $(GO_TEST_FLAGS) $(PACKAGES)

lint:
	for pkg in $(shell go list ./... | grep -v vendor | sed 's/github.com\/zchee\/clang-server/\./g'); do golint $$pkg; done

vet:
	go vet -v -race $(PACKAGES)

fbs:
	flatc --go $(shell find ./ -type f -name '*.fbs')

clang-format:
	clang-format -i -sort-includes $(shell find testdata -type f -name '*.c' -or -name '*.cpp')

vendor/install:
	go install -v -x $(GO_BUILD_TAGS) $(shell glide list 2> /dev/null  | awk 'NR > 1{print $$1}' | sed s'/^/.\/vendor\//g')
	go install -v -x -race $(GO_BUILD_TAGS) $(shell glide list 2> /dev/null  | awk 'NR > 1{print $$1}' | sed s'/^/.\/vendor\//g')

vendor/clean:
	${RM} -r \
		./vendor/github.com/google/flatbuffers/.gitattributes \
		./vendor/github.com/google/flatbuffers/.gitignore \
		./vendor/github.com/google/flatbuffers/.travis.yml \
		./vendor/github.com/google/flatbuffers/android \
		./vendor/github.com/google/flatbuffers/appveyor.yml \
		./vendor/github.com/google/flatbuffers/biicode \
		./vendor/github.com/google/flatbuffers/biicode.conf \
		./vendor/github.com/google/flatbuffers/CMake \
		./vendor/github.com/google/flatbuffers/CMakeLists.txt \
		./vendor/github.com/google/flatbuffers/CMakeLists.txt \
		./vendor/github.com/google/flatbuffers/composer.json \
		./vendor/github.com/google/flatbuffers/docs \
		./vendor/github.com/google/flatbuffers/grpc \
		./vendor/github.com/google/flatbuffers/include \
		./vendor/github.com/google/flatbuffers/ISSUE_TEMPLATE.md \
		./vendor/github.com/google/flatbuffers/java \
		./vendor/github.com/google/flatbuffers/js \
		./vendor/github.com/google/flatbuffers/net \
		./vendor/github.com/google/flatbuffers/php \
		./vendor/github.com/google/flatbuffers/pom.xml \
		./vendor/github.com/google/flatbuffers/python \
		./vendor/github.com/google/flatbuffers/reflection \
		./vendor/github.com/google/flatbuffers/samples \
		./vendor/github.com/google/flatbuffers/src \
		./vendor/github.com/google/flatbuffers/tests

serve: clean build
	./clang-server

clean:
	${RM} clang-server

.PHONY: fbs
