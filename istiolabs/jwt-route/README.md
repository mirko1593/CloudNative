####Validating ingress routing based on JWT claims

1. Validate the ingress gateway returns the HTTP code 404 without JWT
```bash
curl -s -I "http://$INGRESS_HOST:$INGRESS_PORT/headers"
```

2. Validate the ingress gateway returns the HTTP code 401 with invalid JWT
```bash
curl -s -I "http://$INGRESS_HOST:$INGRESS_PORT/headers" -H "Authorization: Bearer some.invalid.token"
```

3. Validate the ingress gateway routes the request with a valid JWT token that includes the claim groups
```bash
TOKEN_GROUP=$(curl https://raw.githubusercontent.com/istio/istio/release-1.14/security/tools/jwt/samples/groups-scope.jwt -s) && echo "$TOKEN_GROUP" | cut -d '.' -f2 - | base64 --decode -

curl -s -I "http://$INGRESS_HOST:$INGRESS_PORT/headers" -H "Authorization: Bearer $TOKEN_GROUP"
```

4. Validate the ingress gateway returns the HTTP code 404 with a valid JWT but does not include the claim groups: group1:
```bash
TOKEN_NO_GROUP=$(curl https://raw.githubusercontent.com/istio/istio/release-1.14/security/tools/jwt/samples/demo.jwt -s) && echo "$TOKEN_NO_GROUP" | cut -d '.' -f2 - | base64 --decode -

curl -s -I "http://$INGRESS_HOST:$INGRESS_PORT/headers" -H "Authorization: Bearer $TOKEN_NO_GROUP"
```

####Cleanup
```bash
kubectl delete namespace foo

kubectl delete requestauthentication ingress-jwt -n istio-system
```


ps. Set the ingress IP and ports:
```bash
export INGRESS_HOST=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
export INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="http2")].port}')
export SECURE_INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="https")].port}')
export TCP_INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="tcp")].port}')
```
