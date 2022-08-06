```shell
export INGRESS_HOST=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
export INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="http2")].port}')
export SECURE_INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="https")].port}')
```

1. Set GATEWAY_URL
```shell
export GATEWAY_URL=$INGRESS_HOST:$INGRESS_PORT
```

2. Generate Trace Data
```shell
for i in $(seq 1 100); do curl -s -o /dev/null "http://$GATEWAY_URL/productpage"; done
```

**TODO:**
1. use API Gateway to handle HTTPs
2. terminate SSL 
3. send plain http to Service
