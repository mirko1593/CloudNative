# Deploy MySQL


## Create Google Persistent Disk

```
gcloud compute disks create mysql
```

```
NAME        ZONE        SIZE_GB  TYPE         STATUS
mysql       us-west1-b  500      pd-standard  READY
```

## Create mysql PersistentVolume

```
kubectl create -f pv/mysql.yaml
```
```
persistentvolume "mysql" created
```

## Create mysql PersistentVolumeClaim

```
kubectl create -f pvc/mysql.yaml 
```
```
persistentvolumeclaim "mysql" created
```

## Create mysql Secrets

```
kubectl create secret generic lobsters \
  --from-literal=root-password=lobsters \
  --from-literal=mysql-password=lobsters \
  --from-literal='database-url=mysql2://lobsters:lobsters@mysql:3306/lobsters'
```

```
secret "lobsters" created
```

## Create mysql Deployment

```
kubectl create -f deployments/mysql.yaml 
```
```
deployment "mysql" created
```

## Create mysql Service

```
kubectl create -f services/mysql.yaml
```

```
service "mysql" created
```
