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

install: build
	@echo "installing......"
	@mkdir -p ~/.local/bin/
	@rm -f ~/.local/bin/$(BIN) ~/.local/bin/$(SH)
	@mv $(BIN) ~/.local/bin/$(BIN)
	@cp $(SH) ~/.local/bin/$(SH)

# clean binary file
clean:
	@rm -f ./$(BIN)