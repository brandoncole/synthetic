build:
	go vet ./ ./resources ./simulator
	go install ./