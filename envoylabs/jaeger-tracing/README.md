To learn about this sandbox and for instructions on how to run it please head over
to the [envoy docs](https://www.envoyproxy.io/docs/envoy/latest/start/sandboxes/jaeger_tracing)

1. Build Flask Service image first.
```bash
docker build -t flask_service:python-3.10-slim-bullseye ../shared/flash
```

2. Build tracing image
```bash
docker build -t envoyproxy:tracing ../shared/tracing
```

