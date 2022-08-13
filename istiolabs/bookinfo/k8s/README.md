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
