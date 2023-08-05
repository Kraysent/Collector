build:
	go build .

build-docker:
	docker build -t collector -f Dockerfile .
	docker tag collector cr.yandex/crpgqd2rqkfgjq6c275u/collector

push-docker:
	docker push cr.yandex/crpgqd2rqkfgjq6c275u/collector
