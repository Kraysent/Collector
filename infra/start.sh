export REGISTRY_ID=crpgqd2rqkfgjq6c275u
docker run \
  -p 9100:9100 \
  --detach \
  --env-file .env \
  cr.yandex/${REGISTRY_ID}/collector:latest
docker run \
  --detach \
  -e AGENT_MODE=flow \
  -v ${PWD}/config.river:/etc/agent/config.river \
  -p 12345:12345 \
  --env-file .env \
  grafana/agent:latest \
    run --server.http.listen-addr=0.0.0.0:12345 /etc/agent/config.river
