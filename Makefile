
# Paths to packages
GO=$(shell which go)

# Paths to locations, etc
BUILD_DIR="build"
BUILD_MODULE="github.com/thevfxcoop/go-deadline-api"
BUILD_LD_FLAGS += -X $(BUILD_MODULE)/pkg/config.GitTag=$(shell git describe --tags)
BUILD_LD_FLAGS += -X $(BUILD_MODULE)/pkg/config.GitBranch=$(shell git name-rev HEAD --name-only --always)
BUILD_LD_FLAGS += -X $(BUILD_MODULE)/pkg/config.GitHash=$(shell git rev-parse HEAD)
BUILD_LD_FLAGS += -X $(BUILD_MODULE)/pkg/config.GoBuildTime=$(shell date -u '+%Y-%m-%dT%H:%M:%SZ')
BUILD_FLAGS = -ldflags "-s -w $(BUILD_LD_FLAGS)" 

.PHONY: all production development dependencies mkdir clean

all: production

dependencies:
ifeq (,${GO})
        $(error "Missing go binary")
endif

mkdir:
	@install -d ${BUILD_DIR}
	${GO} mod tidy

production: dependencies mkdir
	@echo Build for production: ./build/deadline
	${GO} build -o ${BUILD_DIR}/deadline ${BUILD_FLAGS} ./cmd/deadline

development: dependencies mkdir
	@echo Build for development: ./build/deadline
	${GO} build -tags debug -o ${BUILD_DIR}/deadline  ${BUILD_FLAGS} ./cmd/deadline

clean:
	rm -fr $(BUILD_DIR)
	${GO} clean
	${GO} mod tidy