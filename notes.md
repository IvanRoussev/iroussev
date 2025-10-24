
Creating Infra for first time



MetalLB

kubectl create namespace metallb-system --dry-run=client -o yaml | kubectl apply -f -

helm dependency build

helm upgrade --install metallb ./metallb \
  -n metallb-system \
  -f ./metallb/values.yaml


Ingress Nginx

kubectl create namespace ingress-nginx --dry-run=client -o yaml | kubectl apply -f -

helm upgrade --install nginx-ingress ./nginx-ingress \
  -n ingress-nginx \
  -f ./nginx-ingress/values.yaml