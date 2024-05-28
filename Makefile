DISTDIR=./dist
BIN=fj_cli
BUILD_CMD=go build -o $(DISTDIR)/$(BIN)-$@


all: macos-arm64 macos-amd64 linux-amd64

makedir:
	@mkdir $(DISTDIR)

macos-arm64: makedir
	@echo $@
	@GOOS=darwin GOARCH=arm64 $(BUILD_CMD)

macos-amd64: makedir
	@echo $@
	@GOOS=darwin GOARCH=amd64 $(BUILD_CMD)

linux-amd64: makedir
	@echo $@
	@GOOS=linux GOARCH=amd64 $(BUILD_CMD)
