.PHONY: build example
all: build

build:
	go build -o muun .

example: build
	(cd example && ../muun -build-dir build)
