### Bookinfo Application

***Deploy BookInfo Application***
1: details
```bash
kuberctl apply -f service/details.yaml
kuberctl apply -f serviceaccount/details.yaml
kuberctl apply -f deployment/details-v1.yaml
```

2. ratings
```bash
kuberctl apply -f service/ratings.yaml
kuberctl apply -f serviceaccount/ratings.yaml
kuberctl apply -f deployment/ratings-v1.yaml
```

3. reviews
```bash
kuberctl apply -f service/reviews.yaml
kuberctl apply -f serviceaccount/reviews.yaml
kuberctl apply -f deployment/reviews-v1.yaml
kuberctl apply -f deployment/reviews-v2.yaml
kuberctl apply -f deployment/reviews-v3.yaml
```

4. productpage
```bash
kuberctl apply -f service/productpage.yaml
kuberctl apply -f serviceaccount/productpage.yaml
kuberctl apply -f deployment/productpage-v1.yaml
```


***SETUP ENV***
```bash
export INGRESS_HOST=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
export INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="http2")].port}')
export SECURE_INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="https")].port}')
```


Set GATEWAY_URL:
```bash
export GATEWAY_URL=$INGRESS_HOST:$INGRESS_PORT
echo "http://$GATEWAY_URL/productpage"
```

generate payload
```bash
for i in $(seq 1 100); do curl -s -o /dev/null "http://$GATEWAY_URL/productpage"; done
```



***CLEAN UP***
```bash
kuberctl delete -f addons
istioctl manifest generate --set profofile=demo | kuberctl delete --ignore-not-found=true -f -
istioctl tag remove bookinfo

kubectl lable namespace bookinfo istio-injection-
```
