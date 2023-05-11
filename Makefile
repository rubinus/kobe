GOCMD=go
GOBUILD=$(GOCMD) build
BASEPATH := $(shell pwd)
BUILDDIR=$(BASEPATH)/dist

KOBE_SRC=$(BASEPATH)/cmd
KOBE_SERVER_NAME=kobe-server
KOBE_INVENTORY_NAME=kobe-inventory
KOBE_CLIENT_NAME=kobe

BIN_DIR=usr/local/bin
CONFIG_DIR=etc/kobe
BASE_DIR=var/kobe
GOPROXY="https://goproxy.cn,direct"

build_server_linux:
	GOOS=linux  GOARCH=$(GOARCH) $(GOBUILD) -o $(BUILDDIR)/$(BIN_DIR)/$(KOBE_SERVER_NAME) $(KOBE_SRC)/server/*.go
	GOOS=linux  GOARCH=$(GOARCH) $(GOBUILD) -o $(BUILDDIR)/$(BIN_DIR)/$(KOBE_INVENTORY_NAME) $(KOBE_SRC)/inventory/*.go
	mkdir -p $(BUILDDIR)/$(CONFIG_DIR) && cp -r  $(BASEPATH)/conf/* $(BUILDDIR)/$(CONFIG_DIR)
	mkdir -p $(BUILDDIR)/$(BASE_DIR)/plugins/callback && cp  $(BASEPATH)/plugin/* $(BUILDDIR)/$(BASE_DIR)//plugins/callback
build_server_darwin:
	GOOS=darwin  GOARCH=$(GOARCH) $(GOBUILD) -o $(BUILDDIR)/$(BIN_DIR)/$(KOBE_SERVER_NAME) $(KOBE_SRC)/server/*.go
	GOOS=darwin  GOARCH=$(GOARCH) $(GOBUILD) -o $(BUILDDIR)/$(BIN_DIR)/$(KOBE_INVENTORY_NAME) $(KOBE_SRC)/inventory/*.go
	mkdir -p $(BUILDDIR)/$(CONFIG_DIR) && cp -r  $(BASEPATH)/conf/* $(BUILDDIR)/$(CONFIG_DIR)
	mkdir -p $(BUILDDIR)/$(BASE_DIR)/plugins/callback && cp  $(BASEPATH)/plugin/* $(BUILDDIR)/$(BASE_DIR)//plugins/callback

clean:
	rm -fr $(BUILDDIR)

docker:
	@echo "build docker images"
	docker build -t rubinus/kobe:v1.0 .

generate_grpc:
	protoc --go_out=plugins=grpc:./api ./api/kobe.proto
