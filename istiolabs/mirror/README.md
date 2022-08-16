CleanUp

1. Remove the rules
```bash
kubectl delete virtualservice httpbin
kubectl delete destinationrule httpbin
```

2. Shutdown httpbin service and client
```bash
kubectl delete deploy httpbin-v1 httpbin-v2 sleep
kubectl delete svc httpbin
```
