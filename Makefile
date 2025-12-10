PWD := $(shell pwd)
GOPATH := $(shell go env GOPATH)
# golang CPU arch
GOARCH := $(shell go env GOARCH)
# golang OS
GOOS := $(shell go env GOOS)

all: build

# Builds iptv_sc, runs the verifiers then runs the tests.

# Builds iptv_sc locally.
build: 
	@echo "Building iptv_sc binary"
	@GO111MODULE=on CGO_ENABLED=0 go build -trimpath -tags kqueue -gcflags="all=-N -l" -o iptv_process cmd/main.go

# Builds MinIO and installs it to $GOPATH/bin.
install: build
	@echo "Installing binary to '$(GOPATH)/bin/iptv_process'"
	@mkdir -p $(GOPATH)/bin && cp -f $(PWD)/iptv_process $(GOPATH)/bin/iptv_process
	@echo "Installation successful. To learn more, try \"iptv_process --help\"."

clean:
	@echo "Cleaning up all the generated files"
	@find . -name '*.test' | xargs rm -fv
	@find . -name '*~' | xargs rm -fv
	@rm -rvf iptv_process
	@rm -rvf build
	@rm -rvf release
