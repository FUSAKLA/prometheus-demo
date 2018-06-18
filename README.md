# Prometheus demo
This repository contains sources for Prometheus demo originaly presented on
the [Brno Metrics and Monitoring meetup](https://www.meetup.com/Brno-Metrics-and-Monitoring/).

The setup containes 4 components:

### Error server
Simple server written in `Go` which has 2 endpoints:

- `/demo` renders randomly `200` and `500` statuses with image
- `/demo/metrics` exposing Prometheus metrics

### Prometheus
Has simple configuration to scrape the `error-server` metrics.

Also loads alert rule firing when traffic increases on the `error-server`'s `/demo` endpoint.
This alert is sent to the Alertmanager.

### Grafana
Uses dashboard and data-source provisioning which automatically loads
datasource for the Prometheus and dashboard displaying siple traffic graphs for the `error-server`.

### Alertmanager
Alertmanager recieving alerts from Prometheus having simple configuration to
send e-mail with template containing link to presentation and some notes.


## How to run it
You need `docker` and `docker-compose`.

Than it should be enough to run:
```bash
docker-compose up -d
```

Than check:
- Error server: http://localhost:8080/demo and http://localhost:8080/demo/metrics
- Prometheus: http://localhost:9090
- Grafana: http://localhost:3000
- Alertmanager: - Prometheus: http://localhost:9093
