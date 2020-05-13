GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BASEPATH := $(shell pwd)
BUILDDIR=$(BASEPATH)/build

KOBE_SRC=$(BASEPATH)/cmd
KOBE_SERVER_NAME=kobe-server
KOBE_INVENTORY_NAME=kobe-inventory
KOBE_CLIENT_NAME=kobe

BIN_DIR=usr/local/bin
CONFIG_DIR=etc/kobe
BASE_DIR=var/kobe
LIB_DIR=$(BASE_DIR)/lib


build:
	$(GOBUILD) -o $(BUILDDIR)/$(BIN_DIR)/$(KOBE_CLIENT_NAME) $(KOBE_SRC)/client/kobe.go
	$(GOBUILD) -o $(BUILDDIR)/$(BIN_DIR)/$(KOBE_SERVER_NAME) $(KOBE_SRC)/server/kobe-server.go
	$(GOBUILD) -o $(BUILDDIR)/$(BIN_DIR)/$(KOBE_INVENTORY_NAME) $(KOBE_SRC)/inventory/inventory.go
	mkdir -p $(BUILDDIR)/$(LIB_DIR) && cp -r     $(BASEPATH)/ansible $(BUILDDIR)/$(LIB_DIR)
	mkdir -p $(BUILDDIR)/$(CONFIG_DIR) && cp -r  $(BASEPATH)/conf/* $(BUILDDIR)/$(CONFIG_DIR)

clean:
	$(GOCLEAN)
	rm -fr $(BUILDDIR)

docker:
	docker build -t kobe-server

