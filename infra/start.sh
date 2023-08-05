export REGISTRY_ID=crpgqd2rqkfgjq6c275u
docker run \
  --detach \
  --env-file .env \
  --network host \
  cr.yandex/${REGISTRY_ID}/collector:latest
docker run \
  --detach \
  -e AGENT_MODE=flow \
  -v ${PWD}/config.river:/etc/agent/config.river \
  --env-file .env \
  --network host \
  grafana/agent:latest \
    run --server.http.listen-addr=0.0.0.0:12345 /etc/agent/config.river
