build:
	go build .

build-docker:
	docker build -t collector -f Dockerfile .
	docker tag collector cr.yandex/crpgqd2rqkfgjq6c275u/collector

push-docker:
	docker push cr.yandex/crpgqd2rqkfgjq6c275u/collector

build-docker-grafana-agent:
	docker build -t collector-grafana-agent -f infra/Dockerfile .
	docker tag collector-grafana-agent cr.yandex/crpgqd2rqkfgjq6c275u/collector-grafana-agent

push-docker-grafana-agent:
	docker push cr.yandex/crpgqd2rqkfgjq6c275u/collector-grafana-agent
