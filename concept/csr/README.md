1. **usage** has to be 'client auth'
2. **expirationSeconds** could be made longer
3. **request** is the base64 encoded value of CSR file content. could get using command:
```bash
cat mirkowang.csr | base64 | tr -d "\n"
```


Get the list of CSRs:
```bash
kubectl get csr
```

Approve the CSR
```bash
kubectl certificate approve mirkowang
```

Retrieve the certificate from the CSR:
```bash
kubectl get csr/mirkowang -o yaml
```

Export the issued certificate from the CertificateSigningRequest.
```bash
kubectl get csr mirkowang -o jsonpath='{.status.certificate}'| base64 -d > mirkowang.crt
```

This is a sample command to create a Role for this new user:
```bash
kubectl create role developer --verb=create --verb=get --verb=list --verb=update --verb=delete --resource=pods
```

This is a sample command to create a RoleBinding for this new user:
```bash
kubectl create rolebinding developer-binding-mirkowang --role=developer --user=mirkowang
```

First, you need to add new credentials(set-credentials to add new user to cluster):
```bash
kubectl config set-credentials mirkowang --client-key=mirkowang.key --client-certificate=mirkowang.crt --embed-certs=true
```

Then, you need to add the context:
```bash
kubectl config set-context mirkocontext --cluster=gke_gke01-356103_us-west1-a_cloudlabs --user=mirkowang
```
To test it, change the context to myuser:
```bash
kubectl config use-context mirkocontext
```










