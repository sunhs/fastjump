# parameters
GO=go
GOGET=$(GO) get -u
GOTEST=$(GO) test -cover=true
GOBUILD=$(GO) build
GOINSTALL=$(GO) install
GOGENERATE=$(GO) generate
GOBIN=$(HOME)/go/bin

# config
BIN=fastjump
SH="fj.sh"

# ==============================================================================
# default, migrate & build
all: build install

# just build, with details
build:
	@echo "building......"
	@$(GOBUILD)\
		-o $(BIN) \

install:
	@echo "installing......"
	@mv $(BIN) $(GOBIN)

# clean binary file
clean:
	@rm -f ./$(BIN)