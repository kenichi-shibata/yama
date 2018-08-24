OS=linux darwin windows
ARCH=386 amd64
EXT=

.PHONY: clean

define GOBUILD
	for GOOS in $(shell echo ${OS}); \
	do \
		for GOARCH in $(shell echo ${ARCH}); \
			do \
				export GOOS=$$GOOS; \
				export GOARCH=$$GOARCH \
				export CGO_ENABLED=0; \
				if [ $$GOOS == "windows" ]; then EXT=.exe; fi; \
				go build $(GOFLAGS) -o build/yama-$$GOOS-$$GOARCH$$EXT; \
			done \
	done
endef

clean: 
	rm -rf build/

build: GOFLAGS = -a -ldflags='-extldflags=-static'
build:
	mkdir -p build/
	$(call GOBUILD)
