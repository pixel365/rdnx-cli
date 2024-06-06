.PHONY: windows
windows:
		sh ./build.sh windows

.PHONY: darwin
darwin:
		sh ./build.sh darwin

.PHONY: linux
linux:
		sh ./build.sh linux

.PHONY: build
build:
		go build -o ./bin/rdnx-cli

release: windows darwin linux
.PHONY: release

.DEFAULT_GOAL := build