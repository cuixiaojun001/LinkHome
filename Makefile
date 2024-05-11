default: all
PKG_LIST := $(shell find . -type f -name "*.go" | grep -v vendor | sed 's/\.\//LinkHome\//g' | sed 's/\/[^/]*.go//g' | sort | uniq | grep -v "examples/")
LINT_LIST := $(shell echo ${PKG_LIST} | sed 's/LinkHome\///g')
GO_FILES := $(shell find . -name '*.go' | grep -v "_test.go")
REPORT_DIR := "test/reports"

init:
	@echo GOROOT=$(GOROOT)
	@echo GOPATH=$(GOPATH)
	@echo LD_LIBRARY_PATH=$(LD_LIBRARY_PATH)
	@echo PKG_CONFIG_PATH=$(PKG_CONFIG_PATH)
	@echo CGO_CFLAGS=$(CGO_CFLAGS)
	@echo CGO_LDFLAGS=$(CGO_LDFLAGS)

dep:
	yum install -y glibc-static
	yum install -y zlib-static

TARGETLIST=$(shell ls cmd | grep -v bootstrap)

all: ${TARGETLIST}

${TARGETLIST}: %:
	go build -ldflags '-extldflags "-static -logg"' -o bin/http.exe github.com/cuixiaojun001/LinkHome/cmd/http

clean:
	go clean

test: init ## Run unittests
	go test -v ${PKG_LIST}
lint: init ## Lint the files
	golangci-lint run ${LINT_LIST}
race: init ## Run data race detector
	go test -race -short ${PKG_LIST}
