# Collector

Small program that goes to TickTick and removes `needtime` tag from completed tasks so I do not have to do this by hand.

It also sends requests to WakaTime and logs the data from there to PostgreSQL DB.

It features:
- small version of TickTick API client
- small version of WakaTime API client
- `prometheus` metrics and their sending to Grafana
- Storage client for PostgreSQL
- `docker-compose` for collector itself, Grafana Agent and local PostgreSQL DB.