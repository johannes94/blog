kubectl config set-credentials moon --client-key=moon.key --client-certificate=moon.crt --embed-certs=true
kubectl config set-context moon --cluster=minikube --user=moon
kubectl config use-context moon