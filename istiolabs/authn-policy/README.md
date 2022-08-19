### Authn-Policy

1. check sleep.bar to httpbin.foo reachability
```bash
kubectl exec "$(kubectl get pod -l app=sleep -n bar -o jsonpath={.items..metadata.name})" -c sleep -n bar -- curl http://httpbin.foo:8000/ip -s -o /dev/null -w "%{http_code}\n"
```

2. This one-liner command conveniently iterates through all reachability combinations
```bash
for from in "foo" "bar" "legacy"; do for to in "foo" "bar" "legacy"; do kubectl exec "$(kubectl get pod -l app=sleep -n ${from} -o jsonpath={.items..metadata.name})" -c sleep -n ${from} -- curl -s "http://httpbin.${to}:8000/ip" -s -o /dev/null -w "sleep.${from} to httpbin.${to}: %{http_code}\n"; done; done
```

3. Verify there is no peer authentication policy in the system with the following command:
```bash
kubectl get peerauthentication --all-namespaces
```

4. Verify that there are no destination rules that apply on the example services
```bash
kubectl get destinationrules.networking.istio.io --all-namespaces -o yaml | grep "host:"
```

5. Verify envoy proxy inject X-Forwrded-Client-Cert when enable mTLS
```bash
kubectl exec "$(kubectl get pod -l app=sleep -n foo -o jsonpath={.items..metadata.name})" -c sleep -n foo -- curl -s http://httpbin.foo:8000/headers -s | grep X-Forwarded-Client-Cert | sed 's/Hash=[a-z0-9]*;/Hash=<redacted>;/'
```

Plain HTTP
```bash
kubectl exec "$(kubectl get pod -l app=sleep -n foo -o jsonpath={.items..metadata.name})" -c sleep -n foo -- curl http://httpbin.legacy:8000/headers -s | grep X-Forwarded-Client-Cert
```
