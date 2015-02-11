.PHONY: build serve
all: build

build:
	rm -r build
	go run ../*.go -build-dir build
	cp -r css build/

serve:
	php -S 127.0.0.1:8000 -t build
