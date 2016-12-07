.PHONY: build help

build:
	go vet ./ ./cmd ./resources ./simulator
	go install ./

dockerize:
	mkdir -p bin
	CGO_ENABLED=0 GOOS=linux go build -o bin/synthetic -a -installsuffix cgo
	docker build -t brandoncole/synthetic .

help:
	# TODO Only works on macOS right now.  Add support for Windows et al.
	- killall godoc
	godoc -http=:31234 &
	open http://localhost:31234/pkg/github.com/brandoncole/synthetic/