.PHONY: build help

build:
	go vet ./ ./resources ./simulator
	go install ./

help:
	# TODO Only works on macOS right now.  Add support for Windows et al.
	- killall godoc
	godoc -http=:31234 &
	open http://localhost:31234/pkg/github.com/brandoncole/synthetic/