build:
	go build .

build-docker:
	docker build -t collector-image -f Dockerfile .
