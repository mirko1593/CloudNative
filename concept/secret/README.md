### Access API Server

** Create Secret, requesting a token for the default ServiceAccount **
```bash
kubectl apply -f secret/secret.yaml
```

```bash
APISERVER=$(kubectl config view --minify -o jsonpath='{.clusters[0].cluster.server}')
TOKEN=$(kubectl get secret default-token -o jsonpath='{.data.token}' | base64 --decode)

curl $APISERVER/api --header "Authorization: Bearer $TOKEN" --insecure
```
