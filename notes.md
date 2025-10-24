
Creating Infra for first time

kubectl create namespace metallb-system --dry-run=client -o yaml | kubectl apply -f -

helm dependency build

helm upgrade --install metallb ./metallb \
  -n metallb-system \
  -f ./metallb/values.yaml